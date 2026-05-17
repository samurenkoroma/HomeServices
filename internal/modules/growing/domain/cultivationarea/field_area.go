package cultivationarea

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/domain/types"
	"time"
)

type FieldArea struct {
	aggregate.Entity[string]

	FarmRefID string
	Name      string
	Area      float64
}

func (f *FieldArea) SetArea(area float64) {
	f.Area = area
}

func (f *FieldArea) SetName(name string) {
	f.Name = name
}

// NewFieldArea создаёт новое поле как место выращивания
func NewFieldArea(farmRefID, name string, area float64) *FieldArea {
	return &FieldArea{
		Entity:    aggregate.NewEntity(types.NewUUID()),
		FarmRefID: farmRefID,
		Name:      name,
		Area:      area,
	}
}

// Rehydrate восстанавливает поле из БД
func (f *FieldArea) Rehydrate(id string, createdAt, updatedAt time.Time) {
	f.Entity = aggregate.Entity[string]{Id: id, CreatedAt: createdAt, UpdatedAt: updatedAt}
}

// SetFarmRefID устанавливает ссылку на farm модуль
func (f *FieldArea) SetFarmRefID(farmRefID string) {
	f.FarmRefID = farmRefID
	f.Update()
}

// Геттеры
func (f *FieldArea) GetId() string        { return f.Id }
func (f *FieldArea) GetFarmRefID() string { return f.FarmRefID }
func (f *FieldArea) GetType() AreaType    { return AreaTypeField }
func (f *FieldArea) GetName() string      { return f.Name }
func (f *FieldArea) GetArea() float64     { return f.Area }
