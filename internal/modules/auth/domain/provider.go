package domain

type AuthProvider interface {
	Users() UserRepository
	Memberships() MembershipRepository
	Organizations() OrganizationRepository
	ProviderName() string
}
