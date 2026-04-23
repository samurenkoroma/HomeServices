package dto

type UserOrganizationInfo struct {
	OrganizationID   string `json:"id"`
	OrganizationName string `json:"name"`
	Role             string `json:"role"`
	RoleName         string `json:"role_name"`
}

type SwitchOrganizationCmd struct {
	OrganizationID string `json:"organization_id"`
}
