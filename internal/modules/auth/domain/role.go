package domain

// Role роль пользователя
type Role string

const (
	RoleSuperAdmin Role = "super_admin" // полный доступ
	RoleAdmin      Role = "admin"       // полный доступ
	RoleSupport    Role = "support"     // агроном (работа с планами)
	RoleBilling    Role = "billing"     // рабочий (выполнение заданий)
	RoleClient     Role = "client"      // только просмотр
)

// Permissions права доступа
type Permission string

const (
	PermPlanCreate   Permission = "plan:create"
	PermPlanRead     Permission = "plan:read"
	PermPlanUpdate   Permission = "plan:update"
	PermPlanDelete   Permission = "plan:delete"
	PermPlanComplete Permission = "plan:complete"

	PermTaskCreate   Permission = "task:create"
	PermTaskRead     Permission = "task:read"
	PermTaskUpdate   Permission = "task:update"
	PermTaskComplete Permission = "task:complete"

	PermUserCreate Permission = "user:create"
	PermUserRead   Permission = "user:read"
	PermUserUpdate Permission = "user:update"
	PermUserDelete Permission = "user:delete"

	PermVarietyCreate Permission = "variety:create"
	PermVarietyRead   Permission = "variety:read"
	PermVarietyUpdate Permission = "variety:update"

	PermReportView   Permission = "report:view"
	PermReportExport Permission = "report:export"
)

// RolePermissions маппинг ролей на права
var RolePermissions = map[Role][]Permission{
	RoleAdmin: {
		PermPlanCreate, PermPlanRead, PermPlanUpdate, PermPlanDelete, PermPlanComplete,
		PermTaskCreate, PermTaskRead, PermTaskUpdate, PermTaskComplete,
		PermUserCreate, PermUserRead, PermUserUpdate, PermUserDelete,
		PermVarietyCreate, PermVarietyRead, PermVarietyUpdate,
		PermReportView, PermReportExport,
	},
	RoleSupport: {
		PermPlanCreate, PermPlanRead, PermPlanUpdate, PermPlanComplete,
		PermTaskCreate, PermTaskRead, PermTaskUpdate,
		PermVarietyRead,
		PermReportView, PermReportExport,
	},
	RoleBilling: {
		PermPlanRead,
		PermTaskRead, PermTaskComplete,
		PermReportView,
	},
	RoleClient: {
		PermPlanRead,
		PermTaskRead,
		PermReportView,
	},
}

// HasPermission проверяет наличие права
func (r Role) HasPermission(perm Permission) bool {
	perms, ok := RolePermissions[r]
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

// String возвращает строковое представление роли
func (r Role) String() string {
	switch r {
	case RoleSuperAdmin:
		return "Супер-Админ"
	case RoleAdmin:
		return "Администратор"
	case RoleSupport:
		return "Техническая поддержка"
	case RoleBilling:
		return "Биллинг-менеджер"
	case RoleClient:
		return "Пользователь"
	default:
		return "Неизвестно"
	}
}
