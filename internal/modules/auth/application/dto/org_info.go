package dto

type UserOrganizationInfo struct {
	OrganizationID   string `json:"id"`
	OrganizationName string `json:"name"`
	Role             string `json:"role"`
	RoleName         string `json:"roleName"`
}
