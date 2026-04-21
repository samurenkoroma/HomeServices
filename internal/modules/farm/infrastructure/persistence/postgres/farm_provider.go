package postgres

import (
	"database/sql"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/farm/domain/physicalobject"
)

// FarmProvider — провайдер репозиториев для контекста фермы
type FarmProvider struct {
	tx *sql.Tx

	// Кеш репозиториев
	physicalObjects physicalobject.Repository
}

func (p *FarmProvider) ProviderName() string {
	return "farm"
}

// Проверяем, что FarmProvider реализует интерфейс RepositoryProvider
var _ repository.RepositoryProvider = (*FarmProvider)(nil)

func NewPostgresFarmProvider(tx *sql.Tx) repository.RepositoryProvider {
	return &FarmProvider{
		tx: tx,
	}
}

// Objects возвращает репозиторий всех объектов
func (p *FarmProvider) Objects() physicalobject.Repository {
	if p.physicalObjects == nil {
		p.physicalObjects = NewPhysicalObjectRepository(p.tx)
	}
	return p.physicalObjects
}
