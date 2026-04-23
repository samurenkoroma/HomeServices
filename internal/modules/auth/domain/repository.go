package domain

import (
	"context"
)

// UserRepository интерфейс репозитория пользователей
type UserRepository interface {
	// Базовые операции
	Save(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id string) (*User, error)
	Delete(ctx context.Context, id string) error

	// Поиск по уникальным полям
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)

	// Списки
	List(ctx context.Context, filter UserFilter) ([]*User, int, error)

	// Статус
	UpdateLastLogin(ctx context.Context, userID string) error
	UpdateCurrentOrganization(ctx context.Context, userID, organizationID string) error
}

// OrganizationRepository интерфейс репозитория организаций
type OrganizationRepository interface {
	// Базовые операции
	Save(ctx context.Context, org *Organization) error
	Update(ctx context.Context, org *Organization) error
	FindByID(ctx context.Context, id string) (*Organization, error)
	Delete(ctx context.Context, id string) error

	// Поиск
	FindByName(ctx context.Context, name string) (*Organization, error)
	FindByTaxID(ctx context.Context, taxID string) (*Organization, error)

	// Списки
	List(ctx context.Context, filter OrganizationFilter) ([]*Organization, int, error)
	ListByUser(ctx context.Context, userID string) ([]*Organization, error)

	// Статус
	Activate(ctx context.Context, id string) error
	Deactivate(ctx context.Context, id string) error
}

// MembershipRepository интерфейс репозитория членств
type MembershipRepository interface {
	// Базовые операции
	Save(ctx context.Context, membership *Membership) error
	Update(ctx context.Context, membership *Membership) error
	FindByID(ctx context.Context, id string) (*Membership, error)
	Delete(ctx context.Context, id string) error

	// Поиск по связям
	FindByUser(ctx context.Context, userID string) ([]*Membership, error)
	FindByOrganization(ctx context.Context, orgID string) ([]*Membership, error)
	FindByUserAndOrganization(ctx context.Context, userID, orgID string) (*Membership, error)

	// Проверка существования
	Exists(ctx context.Context, userID, orgID string) (bool, error)

	// Списки
	List(ctx context.Context, filter MembershipFilter) ([]*Membership, int, error)

	// Управление статусом
	Activate(ctx context.Context, id string) error
	Deactivate(ctx context.Context, id string) error
	ChangeRole(ctx context.Context, id string, newRole OrganizationRole) error
}

// UserFilter фильтр для пользователей
type UserFilter struct {
	Search string     `json:"search,omitempty"`
	Status UserStatus `json:"status,omitempty"`
	Role   Role       `json:"role,omitempty"`
	OrgID  string     `json:"org_id,omitempty"`
	Limit  int        `json:"limit,omitempty"`
	Offset int        `json:"offset,omitempty"`
}

// OrganizationFilter фильтр для организаций
type OrganizationFilter struct {
	Search   string `json:"search,omitempty"`
	IsActive *bool  `json:"is_active,omitempty"`
	Limit    int    `json:"limit,omitempty"`
	Offset   int    `json:"offset,omitempty"`
}

// MembershipFilter фильтр для членств
type MembershipFilter struct {
	UserID         string           `json:"user_id,omitempty"`
	OrganizationID string           `json:"organization_id,omitempty"`
	Role           OrganizationRole `json:"role,omitempty"`
	IsActive       *bool            `json:"is_active,omitempty"`
	Limit          int              `json:"limit,omitempty"`
	Offset         int              `json:"offset,omitempty"`
}
