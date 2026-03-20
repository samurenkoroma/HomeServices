package postgres

import (
	"database/sql"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/crop/domain/cropplan"
	"samurenkoroma/services/internal/modules/crop/domain/croptype"
)

// CropProvider — провайдер репозиториев для контекста культур
type CropProvider struct {
	tx *sql.Tx

	// Кеш репозиториев
	cropPlans cropplan.Repository
	cropTypes croptype.Repository
}

func (p *CropProvider) ProviderName() string {
	return "crop"
}

// Проверяем, что FarmProvider реализует интерфейс RepositoryProvider
var _ repository.RepositoryProvider = (*CropProvider)(nil)

func NewCropProvider(tx *sql.Tx) *CropProvider {
	return &CropProvider{
		tx: tx,
	}
}

// CropPlans возвращает репозиторий всех объектов
func (p *CropProvider) CropPlans() cropplan.Repository {
	if p.cropPlans == nil {
		p.cropPlans = NewCropPlanRepository(p.tx)
	}
	return p.cropPlans
}

func (p *CropProvider) CropTypes() croptype.Repository {
	if p.cropTypes == nil {
		p.cropTypes = NewCropTypeRepository(p.tx)
	}
	return p.cropTypes
}
