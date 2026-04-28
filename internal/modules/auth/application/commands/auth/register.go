package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/auth/domain"
	"samurenkoroma/services/internal/modules/auth/infrastructure/persistence/postgres"
	"samurenkoroma/services/pkg/response"
)

// RegisterRequest запрос на регистрацию
type RegisterRequest struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
}

// RegisterResponse ответ на регистрацию
type RegisterResponse struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Message  string `json:"message"`
}

// POST /auth/register
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteValidationError(w, "invalid request body")
		return
	}

	// Валидация
	if req.Email == "" {
		response.WriteValidationError(w, "email is required")
		return
	}
	if req.Username == "" {
		response.WriteValidationError(w, "username is required")
		return
	}
	if req.Password == "" || len(req.Password) < 6 {
		response.WriteValidationError(w, "password must be at least 6 characters")
		return
	}

	ctx := r.Context()
	// Начинаем транзакцию
	uow, err := h.uowFactory.Begin(ctx)
	if err != nil {
		response.WriteInternalError(w, "failed to start transaction")
		return
	}
	_, err = uow.Execute(ctx, postgres.NewPostgresAuthProvider, func(provider repository.RepositoryProvider) (any, error) {
		// Приводим провайдер к нужному типу
		authProvider, ok := provider.(*postgres.PostgresAuthProvider)
		if !ok {
			return nil, fmt.Errorf("expected FarmProvider, got %T", provider)
		}
		userRepo := authProvider.Users()

		// Проверяем, не существует ли пользователь
		existing, _ := userRepo.FindByEmail(ctx, req.Email)
		if existing != nil {
			return nil, domain.ErrUserAlreadyExists
		}

		existing, _ = userRepo.FindByUsername(ctx, req.Username)
		if existing != nil {
			return nil, domain.ErrUserAlreadyExists
		}

		// Определяем роль
		role := domain.Role(req.Role)
		if role == "" {
			role = domain.RoleClient // роль по умолчанию
		}

		// Создаем пользователя
		user, err := domain.NewUser(
			req.Email, req.Username, req.Password,
			req.FirstName, req.LastName, req.Phone,
		)
		if err != nil {
			response.WriteValidationError(w, err.Error())
			return nil, err
		}

		// Сохраняем
		if err := userRepo.Save(ctx, user); err != nil {
			return nil, err
		}

		uow.RegisterAggregate(user)
		response.Success(RegisterResponse{
			UserID:   user.ID,
			Email:    user.Email,
			Username: user.Username,
			Role:     string(user.Role),
			Message:  "User registered successfully",
		}).WriteJSON(w, http.StatusCreated)
		return nil, nil
	})
	return

}
