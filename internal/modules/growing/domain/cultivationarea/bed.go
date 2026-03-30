package cultivationarea

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/core/spatial"
	"time"
)

// Bed — грядка
type Bed struct {
	aggregate.Entity[string]

	farmRefID string
	parentID  string // ID теплицы
	name      string
	geometry  spatial.GeoJSON // Точка (центр грядки)
	area      float64

	// Атрибуты грядки (хранятся в JSONB)
	attributes BedAttributes

	seasons         map[string]SeasonConfig
	currentSeasonID string
}

// BedAttributes — атрибуты грядки
type BedAttributes struct {
	Width     float64 `json:"width"`      // ширина (м)
	Length    float64 `json:"length"`     // длина (м)
	PositionX float64 `json:"position_x"` // позиция в % (0-100)
	PositionY float64 `json:"position_y"` // позиция в % (0-100)
}

// NewBed создаёт новую грядку
func NewBed(parentID, name string, centerPoint spatial.GeoJSON) *Bed {
	return &Bed{
		Entity:   aggregate.NewEntity(types.NewUUID()),
		parentID: parentID,
		name:     name,
		geometry: centerPoint,
		seasons:  make(map[string]SeasonConfig),
	}
}

// SetAttributes устанавливает атрибуты грядки
func (b *Bed) SetAttributes(width, length, posX, posY float64) {
	b.attributes = BedAttributes{
		Width:     width,
		Length:    length,
		PositionX: posX,
		PositionY: posY,
	}
	b.Update()
}

// GetAttributes возвращает атрибуты грядки
func (b *Bed) GetAttributes() BedAttributes {
	return b.attributes
}

// GetWidth возвращает ширину
func (b *Bed) GetWidth() float64 {
	return b.attributes.Width
}

// GetLength возвращает длину
func (b *Bed) GetLength() float64 {
	return b.attributes.Length
}

// GetPositionX возвращает позицию X
func (b *Bed) GetPositionX() float64 {
	return b.attributes.PositionX
}

// GetPositionY возвращает позицию Y
func (b *Bed) GetPositionY() float64 {
	return b.attributes.PositionY
}

// ConfigureForSeason — реализация интерфейса CultivationArea
func (b *Bed) ConfigureForSeason(seasonID string, config AreaConfig) error {
	//if config.CropPlanID == nil {
	//	return ErrCropPlanRequiredForBed
	//}

	if _, exists := b.seasons[seasonID]; exists {
		return ErrSeasonAlreadyConfigured
	}

	// Если имя не указано, используем текущее
	name := config.Name
	if name == "" {
		name = b.name
	}

	// Если геометрия не указана, используем текущую
	geom := config.Geometry
	if geom.Type == "" {
		geom = b.geometry
	}

	seasonConfig := SeasonConfig{
		SeasonID:   seasonID,
		Name:       name,
		Geometry:   geom,
		Area:       b.attributes.Width * b.attributes.Length, // для точки площадь 0
		CropPlanID: config.CropPlanID,
		BlockIDs:   []string{},
		Metadata:   config.Metadata,
		ValidFrom:  time.Now(),
	}

	b.seasons[seasonID] = seasonConfig
	b.currentSeasonID = seasonID
	b.name = name
	b.geometry = geom
	b.area = seasonConfig.Area
	b.Update()

	//TODO вынести засев грядки в отдельное событие
	b.AddEvent(BedConfigured{
		BedID:    b.Id,
		SeasonID: seasonID,
		//CropPlanID: *config.CropPlanID,
		Area: seasonConfig.Area,
	})

	return nil
}

// Getters
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

func (b *Bed) GetSeasonConfig(seasonID string) (*SeasonConfig, error) {
	if config, exists := b.seasons[seasonID]; exists {
		return &config, nil
	}
	return nil, ErrSeasonConfigNotFound
}

func (b *Bed) GetCropPlanForSeason(seasonID string) (string, error) {
	config, err := b.GetSeasonConfig(seasonID)
	if err != nil {
		return "", err
	}
	if config.CropPlanID == nil {
		return "", ErrNoCropPlanConfigured
	}
	return *config.CropPlanID, nil
}

// Rehydrate восстанавливает грядку из БД
func (b *Bed) Rehydrate(id, farmRefID string, attrs BedAttributes, createdAt, updatedAt time.Time) {
	b.Id = id
	b.farmRefID = farmRefID
	b.attributes = attrs
	b.CreatedAt = createdAt
	b.UpdatedAt = updatedAt
}

// SetFarmRefID устанавливает ссылку на farm модуль
func (b *Bed) SetFarmRefID(farmRefID string) {
	b.farmRefID = farmRefID
	b.Update()
}

func (b *Bed) GetParentID() string {
	return b.parentID
}
