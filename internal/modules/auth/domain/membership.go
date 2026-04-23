package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Membership членство пользователя в организации
type Membership struct {
	ID             string           `json:"id"`
	UserID         string           `json:"user_id"`
	OrganizationID string           `json:"organization_id"`
	Role           OrganizationRole `json:"role"`
	IsActive       bool             `json:"is_active"`
	JoinedAt       time.Time        `json:"joined_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
}

// NewMembership создает новое членство
func NewMembership(userID, organizationID string, role OrganizationRole) (*Membership, error) {
	if userID == "" {
		return nil, errors.New("user id is required")
	}
	if organizationID == "" {
		return nil, errors.New("organization id is required")
	}
	if role == "" {
		role = OrgRoleViewer
	}

	now := time.Now()
	return &Membership{
		ID:             generateMembershipID(),
		UserID:         userID,
		OrganizationID: organizationID,
		Role:           role,
		IsActive:       true,
		JoinedAt:       now,
		UpdatedAt:      now,
	}, nil
}

// ChangeRole изменяет роль пользователя в организации
func (m *Membership) ChangeRole(newRole OrganizationRole) {
	m.Role = newRole
	m.UpdatedAt = time.Now()
}

// Deactivate деактивирует членство
func (m *Membership) Deactivate() {
	m.IsActive = false
	m.UpdatedAt = time.Now()
}

func (m *Membership) GetRoleName() string {
	switch m.Role {
	case OrgRoleOwner:
		return "Владелец"
	case OrgRoleAdmin:
		return "Администратор"
	case OrgRoleAgronom:
		return "Агроном"
	case OrgRoleWorker:
		return "Рабочий"
	case OrgRoleViewer:
		return "Наблюдатель"
	default:
		return "Неизвестно"
	}
}

func generateMembershipID() string {
	return uuid.New().String()
}
