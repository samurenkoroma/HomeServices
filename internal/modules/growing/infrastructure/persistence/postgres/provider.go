package postgres

import (
	"database/sql"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/growing/domain/croptemplate"
	"samurenkoroma/services/internal/modules/growing/domain/cultivationarea"
	"samurenkoroma/services/internal/modules/growing/domain/season"
)

// GrowingProvider — провайдер репозиториев для контекста культур
type GrowingProvider struct {
	tx *sql.Tx

	// Кеш репозиториев
	seasons          season.Repository
	cultivationArea  cultivationarea.Repository
	cropTemplateRepo croptemplate.Repository
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
	if p.seasons == nil {
		p.seasons = NewSeasonRepository(p.tx)
	}
	return p.seasons
}

// CultivationAreas возвращает репозиторий всех объектов
func (p *GrowingProvider) CultivationAreas() cultivationarea.Repository {
	if p.cultivationArea == nil {
		p.cultivationArea = NewCultivationAreaRepository(p.tx)
	}
	return p.cultivationArea
}
func (p *GrowingProvider) CropTemplates() croptemplate.Repository {
	if p.cropTemplateRepo == nil {
		p.cropTemplateRepo = NewCropTemplateRepository(p.tx)
	}
	return p.cropTemplateRepo
}
