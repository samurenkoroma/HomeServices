package cultivationarea

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/core/spatial"
	"time"
)

// Block — участок внутри поля (существует только в рамках сезона)
type Block struct {
	aggregate.Entity[string]

	farmRefID     string
	parentFieldID string // ID поля из growing (FieldArea)
	name          string
	geometry      spatial.GeoJSON
	area          float64

	// Конфигурации по сезонам
	seasons         map[string]SeasonConfig
	currentSeasonID string
	childBeds       []string // Грядки внутри участка
}

// NewBlock создаёт новый участок
func NewBlock(parentFieldID, name string, geom spatial.GeoJSON) *Block {
	return &Block{
		Entity:        aggregate.NewEntity(types.NewUUID()),
		parentFieldID: parentFieldID,
		name:          name,
		geometry:      geom,
		area:          0,
		seasons:       make(map[string]SeasonConfig),
		childBeds:     []string{},
	}
}

// AddBed — добавить грядку в участок
func (b *Block) AddBed(seasonID string, bedID string) error {
	config, exists := b.seasons[seasonID]
	if !exists {
		return ErrSeasonConfigNotFound
	}

	config.BlockIDs = append(config.BlockIDs, bedID)
	b.seasons[seasonID] = config
	b.childBeds = append(b.childBeds, bedID)
	b.Update()

	return nil
}

// GetCropPlanForSeason — получить план культуры для сезона
func (b *Block) GetCropPlanForSeason(seasonID string) (string, error) {
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
func (b *Block) GetSeasonConfig(seasonID string) (*SeasonConfig, error) {
	if config, exists := b.seasons[seasonID]; exists {
		return &config, nil
	}
	return nil, ErrSeasonConfigNotFound
}

// ConfigureForSeason — реализация интерфейса CultivationArea
func (b *Block) ConfigureForSeason(seasonID string, config AreaConfig) error {
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

// Rehydrate восстанавливает блок из БД
func (b *Block) Rehydrate(id, farmRefID string, createdAt, updatedAt time.Time, seasons []SeasonConfig) {
	b.Entity = aggregate.Entity[string]{Id: id, CreatedAt: createdAt, UpdatedAt: updatedAt}
	for _, season := range seasons {
		b.seasons[season.SeasonID] = season
	}
	b.farmRefID = farmRefID
}

// GetParentFieldID возвращает ID родительского поля
func (b *Block) GetParentFieldID() string {
	return b.parentFieldID
}

// SetFarmRefID устанавливает ссылку на farm модуль
func (b *Block) SetFarmRefID(farmRefID string) {
	b.farmRefID = farmRefID
	b.Update()
}

// Геттеры
func (b *Block) GetId() string                        { return b.Id }
func (b *Block) GetFarmRefID() string                 { return b.farmRefID }
func (b *Block) GetType() AreaType                    { return AreaTypeBlock }
func (b *Block) GetName() string                      { return b.name }
func (b *Block) GetGeometry() spatial.GeoJSON         { return b.geometry }
func (b *Block) GetArea() float64                     { return b.area }
func (b *Block) GetCurrentSeasonID() string           { return b.currentSeasonID }
func (b *Block) IsConfiguredForSeason(id string) bool { _, ok := b.seasons[id]; return ok }
func (b *Block) HasBlocks() bool                      { return false }
func (b *Block) GetBlocks() []string                  { return []string{} }
func (b *Block) GetHistory() []AreaSnapshot           { return []AreaSnapshot{} }
