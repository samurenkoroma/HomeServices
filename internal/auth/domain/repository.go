package domain

import "context"

// Repository интерфейс репозитория пользователей
type Repository interface {
	// Базовые операции
	Save(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	Delete(ctx context.Context, id string) error

	// Списки
	List(ctx context.Context, filter UserFilter) ([]*User, int, error)
	FindByRole(ctx context.Context, role Role) ([]*User, error)

	// Статус
	UpdateLastLogin(ctx context.Context, userID string) error
}

// UserFilter фильтр для списка пользователей
type UserFilter struct {
	Role   Role       `json:"role,omitempty"`
	Status UserStatus `json:"status,omitempty"`
	Search string     `json:"search,omitempty"`
	Limit  int        `json:"limit,omitempty"`
	Offset int        `json:"offset,omitempty"`
}
