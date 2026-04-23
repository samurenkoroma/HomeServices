package memory

import (
	"database/sql"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/auth/domain"
)

type authProvider struct {
	tx *sql.Tx

	users domain.Repository
}

func (p *authProvider) ProviderName() string {
	return "auth"
}

func NewMemoryAuthProvider(tx *sql.Tx) repository.RepositoryProvider {
	return &authProvider{
		tx: tx,
	}
}

func (p *authProvider) Users() domain.Repository {
	if p.users == nil {
		p.users = NewUserRepository()
	}
	return p.users
}
