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
type PhysicalObjectID string

func (i PhysicalObjectID) String() any {
	return string(i)
}

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
	aggregate.Entity[PhysicalObjectID]

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
		Entity:   aggregate.NewEntity(PhysicalObjectID(types.NewUUID())),
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
	obj.AddEvent(FieldCreated{
		ID:       string(obj.Id),
		Name:     obj.Name,
		Geometry: obj.Geometry,
	})
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
		Entity:   aggregate.NewEntity(PhysicalObjectID(types.NewUUID())),
		Type:     ObjectTypeGreenhouse,
		Name:     name,
		Geometry: geom,
		Status:   valueobject.Active,
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

	obj.AddEvent(GreenhouseCreated{
		ID:       string(obj.Id),
		Name:     obj.Name,
		Geometry: obj.Geometry,
	})
	return obj
}

// Методы для работы с атрибутами

func (p *PhysicalObject) SetSoilType(soilType string) {
	if p.Type != ObjectTypeField {
		return // или ошибка
	}
	p.Attributes.SoilType = &soilType
	p.UpdatedAt = time.Now()
}

func (p *PhysicalObject) SetGreenhouseEquipment(heating, ventilation, lighting bool) {
	if p.Type != ObjectTypeGreenhouse {
		return
	}
	p.Attributes.HasHeating = &heating
	p.Attributes.HasVentilation = &ventilation
	p.Attributes.HasLighting = &lighting
	p.UpdatedAt = time.Now()
}

// SetName устанавливает новое имя
func (p *PhysicalObject) SetName(name string) {
	p.Name = name
	p.UpdatedAt = time.Now()
}

// SetDescription устанавливает описание
func (p *PhysicalObject) SetDescription(desc string) {
	p.Description = desc
	p.UpdatedAt = time.Now()
}

// SetGeometry устанавливает новую геометрию
func (p *PhysicalObject) SetGeometry(geom spatial.GeoJSON) error {
	// Валидация геометрии
	if err := validateGeometry(geom); err != nil {
		return err
	}

	p.Geometry = geom
	// Пересчитываем площадь
	//p.TotalArea = calculateArea(geom)
	p.Update()

	return nil
}

// SetAttributes устанавливает атрибуты
func (p *PhysicalObject) SetAttributes(attrs map[string]interface{}) {
	// Обновляем атрибуты, сохраняя существующие
	for k, v := range attrs {
		p.Attributes.Metadata[k] = v
	}
	p.UpdatedAt = time.Now()
}

// Activate активирует объект
func (p *PhysicalObject) Activate() error {
	if p.Status == "active" {
		return ErrAlreadyActive
	}

	p.Status = "active"
	p.UpdatedAt = time.Now()

	p.AddEvent(PhysicalObjectActivated{
		ObjectID: string(p.Id),
		Type:     string(p.Type),
	})

	return nil
}

// Deactivate деактивирует объект
func (p *PhysicalObject) Deactivate() error {
	if p.Status == "inactive" {
		return ErrAlreadyInactive
	}

	p.Status = "inactive"
	p.UpdatedAt = time.Now()

	p.AddEvent(PhysicalObjectDeactivated{
		ObjectID: string(p.Id),
		Type:     string(p.Type),
	})

	return nil
}

// validateGeometry проверяет корректность геометрии
func validateGeometry(geom spatial.GeoJSON) error {
	if geom.Type != spatial.Polygon && geom.Type != spatial.MultiPolygon {
		return ErrInvalidGeometry
	}
	if len(geom.Coordinates) == 0 {
		return ErrEmptyGeometry
	}
	return nil
}

// calculateArea вычисляет площадь по геометрии
func calculateArea(geom spatial.GeoJSON) float64 {
	// В реальности используем PostGIS ST_Area
	// Здесь заглушка
	return 0
}

// Вспомогательные методы для получения данных

func (p *PhysicalObject) IsField() bool {
	return p.Type == ObjectTypeField
}

func (p *PhysicalObject) IsGreenhouse() bool {
	return p.Type == ObjectTypeGreenhouse
}

func RehydrateField(
	id string,
	name string,
	area float64,
) *PhysicalObject {

	return &PhysicalObject{
		Entity: aggregate.Entity[PhysicalObjectID]{
			Id: PhysicalObjectID(id),
		},
		Type:   ObjectTypeField,
		Name:   name,
		Status: valueobject.Active,
		Area:   area,
	}
}
