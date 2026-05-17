package cultivationarea

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"time"
)

// Bed — грядка
type Bed struct {
	aggregate.Entity[string]

	FarmRefID string
	Name      string
	Area      float64
	// Атрибуты грядки (хранятся в JSONB)
	attributes BedAttributes
}

func (b *Bed) SetArea(area float64) {
	b.Area = area
}

// BedAttributes — атрибуты грядки
type BedAttributes struct {
	Width     float64 `json:"width"`      // ширина (м)
	Length    float64 `json:"length"`     // длина (м)
	PositionX float64 `json:"position_x"` // позиция в % (0-100)
	PositionY float64 `json:"position_y"` // позиция в % (0-100)
}

// NewBed создаёт новую грядку
func NewBed(id, farmRefID, name string, area float64) *Bed {
	return &Bed{
		Entity:    aggregate.NewEntity(id),
		FarmRefID: farmRefID,
		Name:      name,
		Area:      area,
	}
}

func (b *Bed) SetName(name string) {
	b.Name = name
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

func (b *Bed) GetId() string        { return b.Id }
func (b *Bed) GetFarmRefID() string { return b.FarmRefID }
func (b *Bed) GetType() AreaType    { return AreaTypeBed }
func (b *Bed) GetName() string      { return b.Name }
func (b *Bed) GetArea() float64     { return b.Area }
func (b *Bed) HasBlocks() bool      { return false }
func (b *Bed) GetBlocks() []string  { return []string{} }

// Rehydrate восстанавливает грядку из БД
func (b *Bed) Rehydrate(createdAt, updatedAt time.Time) {
	b.CreatedAt = createdAt
	b.UpdatedAt = updatedAt
}

// SetFarmRefID устанавливает ссылку на farm модуль
func (b *Bed) SetFarmRefID(farmRefID string) {
	b.FarmRefID = farmRefID
	b.Update()
}
