package cultivationarea

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/core/spatial"
	"time"
)

// Bed — грядка (в теплице или участке)
type Bed struct {
	aggregate.Entity[string]

	farmRefID string
	parentID  string // ID блока или теплицы (из growing)
	name      string
	geometry  spatial.GeoJSON
	area      float64

	// Конфигурации по сезонам
	seasons         map[string]SeasonConfig
	currentSeasonID string
}

// NewBed создаёт новую грядку
func NewBed(parentID, name string, geom spatial.GeoJSON) *Bed {
	return &Bed{
		Entity:   aggregate.NewEntity(types.NewUUID()),
		parentID: parentID,
		name:     name,
		geometry: geom,
		area:     0,
		seasons:  make(map[string]SeasonConfig),
	}
}

// GetCropPlanForSeason — получить план культуры для сезона
func (b *Bed) GetCropPlanForSeason(seasonID string) (string, error) {
	config, exists := b.seasons[seasonID]
	if !exists {
		return "", ErrSeasonConfigNotFound
	}
	if config.CropPlanID == nil {
		return "", ErrNoCropPlanConfigured
	}
	return *config.CropPlanID, nil
}

// GetSeasonConfig — получить конфигурацию на сезон
func (b *Bed) GetSeasonConfig(seasonID string) (*SeasonConfig, error) {
	if config, exists := b.seasons[seasonID]; exists {
		return &config, nil
	}
	return nil, ErrSeasonConfigNotFound
}

// ConfigureForSeason — реализация интерфейса
func (b *Bed) ConfigureForSeason(seasonID string, config AreaConfig) error {
	if config.CropPlanID == nil {
		return ErrCropPlanRequiredForBlock
	}

	if _, exists := b.seasons[seasonID]; exists {
		return ErrSeasonAlreadyConfigured
	}

	seasonConfig := SeasonConfig{
		SeasonID:   seasonID,
		Name:       config.Name,
		Geometry:   config.Geometry,
		Area:       0,
		CropPlanID: config.CropPlanID,
		BlockIDs:   []string{},
		Metadata:   config.Metadata,
		ValidFrom:  time.Now(),
	}

	b.seasons[seasonID] = seasonConfig
	b.currentSeasonID = seasonID
	b.name = config.Name
	b.geometry = config.Geometry
	b.area = seasonConfig.Area
	b.Update()

	b.AddEvent(BlockConfigured{
		BlockID:    b.Id,
		SeasonID:   seasonID,
		CropPlanID: *config.CropPlanID,
		Area:       seasonConfig.Area,
	})

	return nil
}

// Rehydrate восстанавливает грядку из БД
func (b *Bed) Rehydrate(id, farmRefID string, createdAt, updatedAt time.Time) {
	b.Entity = aggregate.Entity[string]{Id: id, CreatedAt: createdAt, UpdatedAt: updatedAt}
	b.farmRefID = farmRefID
}

// GetParentID возвращает ID родителя (блока или теплицы)
func (b *Bed) GetParentID() string {
	return b.parentID
}

// SetFarmRefID устанавливает ссылку на farm модуль
func (b *Bed) SetFarmRefID(farmRefID string) {
	b.farmRefID = farmRefID
	b.Update()
}

// Геттеры
func (b *Bed) GetID() string                        { return b.Id }
func (b *Bed) GetFarmRefID() string                 { return b.farmRefID }
func (b *Bed) GetType() AreaType                    { return AreaTypeBed }
func (b *Bed) GetName() string                      { return b.name }
func (b *Bed) GetGeometry() spatial.GeoJSON         { return b.geometry }
func (b *Bed) GetArea() float64                     { return b.area }
func (b *Bed) GetCurrentSeasonID() string           { return b.currentSeasonID }
func (b *Bed) IsConfiguredForSeason(id string) bool { _, ok := b.seasons[id]; return ok }
func (b *Bed) HasBlocks() bool                      { return false }
func (b *Bed) GetBlocks() []string                  { return []string{} }
func (b *Bed) GetHistory() []AreaSnapshot           { return []AreaSnapshot{} }
