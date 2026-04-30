package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/auth/domain"
	"samurenkoroma/services/internal/modules/auth/infrastructure/jwt"
	"samurenkoroma/services/internal/modules/auth/infrastructure/persistence/postgres"
	"samurenkoroma/services/pkg/response"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
type LoginResult struct {
	TokenPair    *jwt.TokenPair `json:"tokenPair"`
	User         User           `json:"user"`
	CurrentOrgId string         `json:"currentOrgId,omitempty"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteValidationError(w, "invalid request body")
		return
	}

	if req.Email == "" {
		response.WriteValidationError(w, "email is required")
		return
	}
	if req.Password == "" || len(req.Password) < 6 {
		response.WriteValidationError(w, "password must be at least 6 characters")
		return
	}

	uow, err := h.uowFactory.Begin(r.Context())
	if err != nil {
		response.WriteInternalError(w, "failed to start transaction")
		return
	}
	ctx := r.Context()
	_, err = uow.Execute(ctx, postgres.NewPostgresAuthProvider, func(provider repository.RepositoryProvider) (any, error) {
		// Приводим провайдер к нужному типу
		authProvider, ok := provider.(*postgres.PostgresAuthProvider)
		if !ok {
			response.WriteInternalError(w, err.Error())
			return nil, fmt.Errorf("expected FarmProvider, got %T", provider)
		}

		userRepo := authProvider.Users()
		membershipRepo := authProvider.Memberships()

		// Ищем пользователя
		user, err := userRepo.FindByEmail(ctx, req.Email)
		if err != nil {
			response.WriteNotFound(w, err.Error())
			return nil, domain.ErrInvalidCredentials
		}

		// Проверяем пароль
		if !user.CheckPassword(req.Password) {
			response.WriteValidationError(w, domain.ErrInvalidCredentials.Error())
			return nil, domain.ErrInvalidCredentials
		}

		// Проверяем статус
		if !user.IsActive() {
			return nil, domain.ErrUserInactive
		}

		currentOrgID := user.GetCurrentOrganizationID()
		var orgRole string
		if currentOrgID != "" {
			// Получаем все членства пользователя
			membership, err := membershipRepo.FindByUserAndOrganization(ctx, user.ID, currentOrgID)
			if err != nil {
				return nil, err
			}
			orgRole = membership.GetRoleName()
		}

		// Генерируем токены с текущей организацией

		tokenPair, err := h.jwtService.GenerateTokenPair(
			user.ID,
			user.Username,
			user.Email,
			string(user.Role),
			currentOrgID,
			orgRole,
		)
		if err != nil {
			return nil, err
		}

		// Обновляем время последнего входа
		user.UpdateLastLogin()
		userRepo.Update(ctx, user)

		uow.RegisterAggregate(user)
		response.Success(
			LoginResult{
				TokenPair: tokenPair,
				User: User{
					Id:    user.ID,
					Name:  user.Username,
					Email: user.Email,
					Role:  user.Role.String(),
				},
				CurrentOrgId: currentOrgID,
			}).WriteJSON(w, http.StatusOK)

		return nil, nil
	})
	return
}
