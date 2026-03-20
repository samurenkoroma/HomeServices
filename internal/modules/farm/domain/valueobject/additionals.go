package valueobject

import "time"

type Additions struct {
	IrrigationType IrrigationType // Тип полива
	Status         AreaStatus
}

func DefaultAdditions() Additions {
	return Additions{
		IrrigationType: IrrigationSprinkler,
		Status:         AreaStatusEmpty,
	}
}

type IrrigationType string

const (
	IrrigationDrip      IrrigationType = "drip"      // Капельный
	IrrigationSprinkler IrrigationType = "sprinkler" // Дождевание
	IrrigationFlood     IrrigationType = "flood"     // Затопление
	IrrigationNone      IrrigationType = "none"
)

type AreaStatus string

const (
	AreaStatusActive    AreaStatus = "active"
	AreaStatusPreparing AreaStatus = "preparing"
	AreaStatusEmpty     AreaStatus = "empty"
	AreaStatusPlanted   AreaStatus = "planted"
	AreaStatusGrowing   AreaStatus = "growing"
	AreaStatusHarvested AreaStatus = "harvested"
	AreaStatusFallow    AreaStatus = "fallow"
)

type CurrentCropInfo struct {
	CropCycleID         string // ID цикла выращивания
	CropID              string // ID культуры
	VarietyID           string // ID сорта
	PlantedAt           time.Time
	ExpectedHarvestDate time.Time
}
