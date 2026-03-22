package postgres

import (
	"database/sql"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/domain/cropcycle"
	"samurenkoroma/services/internal/modules/growing/domain/croptemplate"
	"samurenkoroma/services/internal/modules/growing/domain/cultivationarea"
	"samurenkoroma/services/internal/modules/growing/domain/season"
)

// GrowingProvider — провайдер репозиториев для контекста культур
type GrowingProvider struct {
	tx *sql.Tx

	// Кеш репозиториев
	seasonsRepo         season.Repository
	cultivationAreaRepo cultivationarea.Repository
	cropTemplateRepo    croptemplate.Repository
	cropCyclesRepo      cropcycle.Repository
}

func (p *GrowingProvider) ProviderName() string {
	return "growing"
}

// Проверяем, что GrowingProvider реализует интерфейс RepositoryProvider
var _ repository.RepositoryProvider = (*GrowingProvider)(nil)

func NewGrowingProvider(tx *sql.Tx) repository.RepositoryProvider {
	return &GrowingProvider{
		tx: tx,
	}
}

// Seasons возвращает репозиторий всех объектов
func (p *GrowingProvider) Seasons() season.Repository {
	if p.seasonsRepo == nil {
		p.seasonsRepo = NewSeasonRepository(p.tx)
	}
	return p.seasonsRepo
}

// CultivationAreas возвращает репозиторий всех объектов
func (p *GrowingProvider) CultivationAreas() cultivationarea.Repository {
	if p.cultivationAreaRepo == nil {
		p.cultivationAreaRepo = NewCultivationAreaRepository(p.tx)
	}
	return p.cultivationAreaRepo
}
func (p *GrowingProvider) CropTemplates() croptemplate.Repository {
	if p.cropTemplateRepo == nil {
		p.cropTemplateRepo = NewCropTemplateRepository(p.tx)
	}
	return p.cropTemplateRepo
}

func (p *GrowingProvider) CropCycles() cropcycle.Repository {
	if p.cropCyclesRepo == nil {
		p.cropCyclesRepo = NewCropCycleRepository(p.tx)
	}
	return p.cropCyclesRepo
}
