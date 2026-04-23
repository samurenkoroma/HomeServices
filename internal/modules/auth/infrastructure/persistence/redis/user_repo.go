package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"samurenkoroma/services/internal/modules/auth/domain"

	"github.com/redis/go-redis/v9"
)

// UserRepository реализация репозитория для PostgreSQL
type userRepository struct {
	client *redis.Client
}

// NewUserRepository создает новый репозиторий в транзакции
func NewRedisUserRepository(client *redis.Client) domain.Repository {
	return &userRepository{client: client}
}

func (u userRepository) Save(ctx context.Context, user *domain.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal session: %w", err)
	}

	err = u.client.Set(ctx, user.ID, data, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to save session: %w", err)
	}
	return nil
}

func (u userRepository) Update(ctx context.Context, user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	return nil, errors.New("not found")
}

func (u userRepository) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	return nil, errors.New("not found")
}

func (u userRepository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) List(ctx context.Context, filter domain.UserFilter) ([]*domain.User, int, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) FindByRole(ctx context.Context, role domain.Role) ([]*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) UpdateLastLogin(ctx context.Context, userID string) error {
	//TODO implement me
	panic("implement me")
}
