package command

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"samurenkoroma/services/internal/core/domain/repository"

	"samurenkoroma/services/internal/auth/domain"
)

// RegisterUserHandler команда регистрации пользователя
type RegisterUserHandler struct {
	UowFactory repository.Factory
}

// RegisterUserCmd структура команды
type RegisterUserCmd struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
}

// DecodeRegisterUser декодирует JSON в команду
func DecodeRegisterUser(data []byte) (any, error) {
	var cmd RegisterUserCmd
	if err := json.Unmarshal(data, &cmd); err != nil {
		return nil, fmt.Errorf("failed to decode RegisterUser command: %w", err)
	}

	if cmd.Email == "" {
		return nil, errors.New("email is required")
	}
	if cmd.Username == "" {
		return nil, errors.New("username is required")
	}
	if cmd.Password == "" {
		return nil, errors.New("password is required")
	}
	if len(cmd.Password) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}

	return cmd, nil
}

// Handle выполняет команду
func (h *RegisterUserHandler) Handle(ctx context.Context, cmd any) error {
	c, ok := cmd.(RegisterUserCmd)
	if !ok {
		return errors.New("invalid command type")
	}

	uowObj, err := h.UowFactory.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin unit of work: %w", err)
	}
	defer uowObj.Rollback()

	userRepo := uowObj.Users()

	// Проверяем, не существует ли пользователь
	existing, _ := userRepo.FindByEmail(ctx, c.Email)
	if existing != nil {
		return domain.ErrUserAlreadyExists
	}

	existing, _ = userRepo.FindByUsername(ctx, c.Username)
	if existing != nil {
		return domain.ErrUserAlreadyExists
	}

	// Определяем роль
	role := domain.Role(c.Role)
	if role == "" {
		role = domain.RoleWorker // роль по умолчанию
	}

	// Создаем пользователя
	user, err := domain.NewUser(
		c.Email, c.Username, c.Password,
		c.FirstName, c.LastName, c.Phone,
		role,
	)
	if err != nil {
		return err
	}

	// Сохраняем
	if err := userRepo.Save(ctx, user); err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}

	uowObj.RegisterAggregate(user)

	if err := uowObj.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
