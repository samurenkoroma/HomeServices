package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Organization организация (хозяйство, ферма, предприятие)
type Organization struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	TaxID     string    `json:"tax_id"` // ИНН
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewOrganization создает новую организацию
func NewOrganization(name, taxID, address, phone, email string) (*Organization, error) {
	if name == "" {
		return nil, errors.New("organization name is required")
	}

	now := time.Now()
	return &Organization{
		ID:        uuid.New().String(),
		Name:      name,
		TaxID:     taxID,
		Address:   address,
		Phone:     phone,
		Email:     email,
		IsActive:  true,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// OrganizationRole роль пользователя в организации
type OrganizationRole string

const (
	OrgRoleOwner   OrganizationRole = "owner"   // владелец (полный доступ)
	OrgRoleAdmin   OrganizationRole = "admin"   // администратор организации
	OrgRoleAgronom OrganizationRole = "agronom" // агроном
	OrgRoleWorker  OrganizationRole = "worker"  // рабочий
	OrgRoleViewer  OrganizationRole = "viewer"  // только просмотр
)

// OrganizationPermissions права в рамках организации
var OrganizationPermissions = map[OrganizationRole][]Permission{
	OrgRoleOwner: {
		PermPlanCreate, PermPlanRead, PermPlanUpdate, PermPlanDelete, PermPlanComplete,
		PermTaskCreate, PermTaskRead, PermTaskUpdate, PermTaskComplete,
		PermUserCreate, PermUserRead, PermUserUpdate, PermUserDelete,
		PermVarietyCreate, PermVarietyRead, PermVarietyUpdate,
		PermReportView, PermReportExport,
		"org:manage", "org:delete",
	},
	OrgRoleAdmin: {
		PermPlanCreate, PermPlanRead, PermPlanUpdate, PermPlanComplete,
		PermTaskCreate, PermTaskRead, PermTaskUpdate,
		PermUserCreate, PermUserRead, PermUserUpdate,
		PermVarietyRead,
		PermReportView, PermReportExport,
		"org:manage",
	},
	OrgRoleAgronom: {
		PermPlanCreate, PermPlanRead, PermPlanUpdate, PermPlanComplete,
		PermTaskCreate, PermTaskRead, PermTaskUpdate,
		PermVarietyRead,
		PermReportView, PermReportExport,
	},
	OrgRoleWorker: {
		PermPlanRead,
		PermTaskRead, PermTaskComplete,
		PermReportView,
	},
	OrgRoleViewer: {
		PermPlanRead,
		PermTaskRead,
		PermReportView,
	},
}

// HasPermission проверяет наличие права в организации
func (r OrganizationRole) HasPermission(perm Permission) bool {
	perms, ok := OrganizationPermissions[r]
	if !ok {
		return false
	}
	for _, p := range perms {
		if p == perm {
			return true
		}
	}
	return false
}
