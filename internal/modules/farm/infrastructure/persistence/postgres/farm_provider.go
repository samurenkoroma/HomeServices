package postgres

import (
	"database/sql"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/farm/domain/bed"
	"samurenkoroma/services/internal/modules/farm/domain/field"
	"samurenkoroma/services/internal/modules/farm/domain/field_block"
	"samurenkoroma/services/internal/modules/farm/domain/greenhouse"
)

// FarmProvider — провайдер репозиториев для контекста фермы
type FarmProvider struct {
	tx *sql.Tx

	// Кеш репозиториев
	fieldRepo      field.Repository
	blockRepo      field_block.Repository
	greenhouseRepo greenhouse.Repository
	bedRepo        bed.Repository
}

func (p *FarmProvider) ProviderName() string {
	return "farm"
}

// Проверяем, что FarmProvider реализует интерфейс RepositoryProvider
var _ repository.RepositoryProvider = (*FarmProvider)(nil)

func NewFarmProvider(tx *sql.Tx) *FarmProvider {
	return &FarmProvider{
		tx: tx,
	}
}

// Fields возвращает репозиторий полей
func (p *FarmProvider) Fields() field.Repository {
	if p.fieldRepo == nil {
		p.fieldRepo = NewFieldRepository(p.tx)
	}
	return p.fieldRepo
}

// FieldBlocks возвращает репозиторий блоков
func (p *FarmProvider) FieldBlocks() field_block.Repository {
	if p.blockRepo == nil {
		p.blockRepo = NewFieldBlockRepository(p.tx)
	}
	return p.blockRepo
}

// Greenhouses возвращает репозиторий теплиц
func (p *FarmProvider) Greenhouses() greenhouse.Repository {
	if p.greenhouseRepo == nil {
		p.greenhouseRepo = NewGreenhouseRepository(p.tx)
	}
	return p.greenhouseRepo
}

// Beds возвращает репозиторий грядок
func (p *FarmProvider) Beds() bed.Repository {
	if p.bedRepo == nil {
		p.bedRepo = NewBedRepository(p.tx)
	}
	return p.bedRepo
}
