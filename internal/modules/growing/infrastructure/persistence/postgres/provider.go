package postgres

import (
	"database/sql"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/domain/season"
)

// GrowingProvider — провайдер репозиториев для контекста культур
type GrowingProvider struct {
	tx *sql.Tx

	// Кеш репозиториев
	seasons season.Repository
}

func (p *GrowingProvider) ProviderName() string {
	return "growing"
}

// Проверяем, что FarmProvider реализует интерфейс RepositoryProvider
var _ repository.RepositoryProvider = (*GrowingProvider)(nil)

func NewGrowingProvider(tx *sql.Tx) repository.RepositoryProvider {
	return &GrowingProvider{
		tx: tx,
	}
}

// Seasons возвращает репозиторий всех объектов
func (p *GrowingProvider) Seasons() season.Repository {
	if p.seasons == nil {
		p.seasons = NewSeasonRepository(p.tx)
	}
	return p.seasons
}
