package cropcycle

import "time"

type OperationType string

const (
	OperationTypePlanting    OperationType = "planting"
	OperationTypeWatering    OperationType = "watering"
	OperationTypeFertilizing OperationType = "fertilizing"
	OperationTypePestControl OperationType = "pest_control"
	OperationTypeWeeding     OperationType = "weeding"
	OperationTypeHarvesting  OperationType = "harvesting"
)

// Operation - операция в цикле выращивания
type Operation struct {
	ID          string
	Type        OperationType
	Description string
	Amount      float64 // Количество (литры, кг)
	Unit        string  // Единица измерения
	PerformedBy string  // Кто выполнил
	Notes       string
	CreatedAt   time.Time
}
