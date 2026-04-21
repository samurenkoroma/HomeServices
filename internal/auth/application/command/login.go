package command

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"samurenkoroma/services/internal/auth/domain"
	"samurenkoroma/services/internal/auth/infrastructure/jwt"
	"samurenkoroma/services/internal/common/application/uow"
)

// LoginHandler команда входа
type LoginHandler struct {
	UowFactory uow.Factory
	JWTService *jwt.Service
}

// LoginCmd структура команды
type LoginCmd struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse ответ с токенами
type LoginResponse struct {
	TokenPair *jwt.TokenPair `json:"token_pair"`
	UserID    string         `json:"user_id"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Role      string         `json:"role"`
	RoleName  string         `json:"role_name"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
}

// DecodeLogin декодирует JSON в команду
func DecodeLogin(data []byte) (any, error) {
	var cmd LoginCmd
	if err := json.Unmarshal(data, &cmd); err != nil {
		return nil, fmt.Errorf("failed to decode Login command: %w", err)
	}

	if cmd.Username == "" {
		return nil, errors.New("username is required")
	}
	if cmd.Password == "" {
		return nil, errors.New("password is required")
	}

	return cmd, nil
}

// Handle выполняет команду
func (h *LoginHandler) Handle(ctx context.Context, cmd any) (any, error) {
	c, ok := cmd.(LoginCmd)
	if !ok {
		return nil, errors.New("invalid command type")
	}

	uowObj, err := h.UowFactory.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin unit of work: %w", err)
	}
	defer uowObj.Rollback()

	userRepo := uowObj.Users()

	// Ищем пользователя по email или username
	user, err := userRepo.FindByEmail(ctx, c.Username)
	if err != nil && !errors.Is(err, domain.ErrUserNotFound) {
		user, err = userRepo.FindByUsername(ctx, c.Username)
	}
	if err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	// Проверяем пароль
	if !user.CheckPassword(c.Password) {
		return nil, domain.ErrInvalidCredentials
	}

	// Проверяем статус
	if !user.IsActive() {
		return nil, domain.ErrUserInactive
	}

	// Генерируем токены
	tokenPair, err := h.JWTService.GenerateTokenPair(
		user.ID,
		user.Username,
		user.Email,
		string(user.Role),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}

	// Обновляем время последнего входа
	user.UpdateLastLogin()
	if err := userRepo.Update(ctx, user); err != nil {
		// Не критично, продолжаем
	}

	uowObj.RegisterAggregate(user)
	uowObj.Commit()

	return &LoginResponse{
		TokenPair: tokenPair,
		UserID:    user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      string(user.Role),
		RoleName:  user.Role.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}
