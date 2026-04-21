package domain

// Role роль пользователя
type Role string

const (
	RoleAdmin   Role = "admin"   // полный доступ
	RoleAgronom Role = "agronom" // агроном (работа с планами)
	RoleWorker  Role = "worker"  // рабочий (выполнение заданий)
	RoleViewer  Role = "viewer"  // только просмотр
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
	RoleAgronom: {
		PermPlanCreate, PermPlanRead, PermPlanUpdate, PermPlanComplete,
		PermTaskCreate, PermTaskRead, PermTaskUpdate,
		PermVarietyRead,
		PermReportView, PermReportExport,
	},
	RoleWorker: {
		PermPlanRead,
		PermTaskRead, PermTaskComplete,
		PermReportView,
	},
	RoleViewer: {
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
	case RoleAdmin:
		return "Администратор"
	case RoleAgronom:
		return "Агроном"
	case RoleWorker:
		return "Рабочий"
	case RoleViewer:
		return "Наблюдатель"
	default:
		return "Неизвестно"
	}
}
