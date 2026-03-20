package physicalobject

import (
	"encoding/json"
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/core/spatial"
	"samurenkoroma/services/internal/modules/farm/domain/valueobject"
	"time"
)

type ObjectType string

// GreenhouseType определяет тип теплицы
type GreenhouseType string

const (
	GreenhouseTypeFilm          GreenhouseType = "film"          // Пленочная
	GreenhouseTypeGlass         GreenhouseType = "glass"         // Стеклянная
	GreenhouseTypePolycarbonate GreenhouseType = "polycarbonate" // Поликарбонат
)
const (
	ObjectTypeField      ObjectType = "field"
	ObjectTypeGreenhouse ObjectType = "greenhouse"
	ObjectTypeBuilding   ObjectType = "building"
	ObjectTypeStorage    ObjectType = "storage"
)

// PhysicalObject - единый агрегат для всех физических объектов
type PhysicalObject struct {
	aggregate.Entity[types.PhysicalObjectID]

	Type        ObjectType
	Name        string
	Geometry    spatial.GeoJSON
	Status      valueobject.Status
	OwnerID     string
	Description string

	// Специфические атрибуты в зависимости от типа
	Attributes Attributes

	Area float64
}

// Attributes - типизированные атрибуты для разных объектов
type Attributes struct {
	// Для поля
	SoilType *string `json:"soil_type,omitempty"`

	// Для теплицы
	GreenhouseType *string  `json:"greenhouse_type,omitempty"` // film, glass, polycarbonate
	Height         *float64 `json:"height,omitempty"`
	Width          *float64 `json:"width,omitempty"`
	Length         *float64 `json:"length,omitempty"`
	HasHeating     *bool    `json:"has_heating,omitempty"`
	HasVentilation *bool    `json:"has_ventilation,omitempty"`
	HasLighting    *bool    `json:"has_lighting,omitempty"`

	// Общие
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

func (a *Attributes) Marshal() ([]byte, error) {
	return json.Marshal(a)
}

func (a *Attributes) Unmarshal(data []byte) error {
	return json.Unmarshal(data, a)
}

// Фабричные методы для разных типов

func NewField(
	name string,
	geom spatial.GeoJSON,
	soilType string,
	ownerID string,
) *PhysicalObject {

	obj := &PhysicalObject{
		Entity:   aggregate.NewEntity(types.PhysicalObjectID(types.NewUUID())),
		Type:     ObjectTypeField,
		Name:     name,
		Geometry: geom,
		Status:   valueobject.Active,
		OwnerID:  ownerID,
		Area:     0,
		Attributes: Attributes{
			SoilType: &soilType,
		},
	}
	obj.AddEvent(NewFieldCreated(string(obj.Id), obj.Name, obj.Area))
	return obj
}

func NewGreenhouse(
	name string,
	dim valueobject.Dimension,
	geom spatial.GeoJSON,
	ghType string,
	ownerID string,
) *PhysicalObject {

	obj := &PhysicalObject{
		Entity:   aggregate.NewEntity(types.PhysicalObjectID(types.NewUUID())),
		Type:     ObjectTypeGreenhouse,
		Name:     name,
		Geometry: geom,
		Status:   "active",
		Area:     dim.Area(),
		OwnerID:  ownerID,
		Attributes: Attributes{
			GreenhouseType: &ghType,
			Width:          dim.Width,
			Height:         dim.Height,
			Length:         dim.Length,

			HasHeating:     new(bool),
			HasVentilation: new(bool),
			HasLighting:    new(bool),
		},
	}

	obj.AddEvent(NewGreenhouseCreated(string(obj.Id), obj.Name, obj.Area))
	return obj
}

// Методы для работы с атрибутами

func (o *PhysicalObject) SetSoilType(soilType string) {
	if o.Type != ObjectTypeField {
		return // или ошибка
	}
	o.Attributes.SoilType = &soilType
	o.UpdatedAt = time.Now()
}

func (o *PhysicalObject) SetGreenhouseEquipment(heating, ventilation, lighting bool) {
	if o.Type != ObjectTypeGreenhouse {
		return
	}
	o.Attributes.HasHeating = &heating
	o.Attributes.HasVentilation = &ventilation
	o.Attributes.HasLighting = &lighting
	o.UpdatedAt = time.Now()
}

// Вспомогательные методы для получения данных

func (o *PhysicalObject) IsField() bool {
	return o.Type == ObjectTypeField
}

func (o *PhysicalObject) IsGreenhouse() bool {
	return o.Type == ObjectTypeGreenhouse
}

func RehydrateField(
	id string,
	name string,
	area float64,
) *PhysicalObject {

	return &PhysicalObject{
		Entity: aggregate.Entity[types.PhysicalObjectID]{
			Id: types.PhysicalObjectID(id),
		},
		Type:   ObjectTypeField,
		Name:   name,
		Status: valueobject.Active,
		Area:   area,
	}
}
