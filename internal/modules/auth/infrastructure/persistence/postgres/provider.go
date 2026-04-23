package postgres

import (
	"database/sql"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/auth/domain"
)

type PostgresAuthProvider struct {
	tx *sql.Tx

	users       domain.UserRepository
	orgs        domain.OrganizationRepository
	memberships domain.MembershipRepository
}

func (p *PostgresAuthProvider) ProviderName() string {
	return "auth"
}

func NewPostgresAuthProvider(tx *sql.Tx) repository.RepositoryProvider {
	return &PostgresAuthProvider{
		tx: tx,
	}
}

func (p *PostgresAuthProvider) Memberships() domain.MembershipRepository {
	if p.memberships == nil {
		p.memberships = NewMembershipRepository(p.tx)
	}
	return p.memberships
}

func (p *PostgresAuthProvider) Organizations() domain.OrganizationRepository {
	if p.orgs == nil {
		p.orgs = NewOrganizationRepository(p.tx)
	}
	return p.orgs
}

func (p *PostgresAuthProvider) Users() domain.UserRepository {
	if p.users == nil {
		p.users = NewUserRepository(p.tx)
	}
	return p.users
}
