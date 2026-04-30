package season

import (
	"context"
	"time"
)

type Filter struct {
	Limit   int    `json:"limit,omitempty"`
	Offset  int    `json:"offset,omitempty"`
	OwnerId string `json:"owner_id,omitempty"`
}

type ListItem struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedBy   string `json:"created_by"`
	CreatedAt   string `json:"created_at"`
}
type Detail struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Status      string    `json:"status"`
	Description string    `json:"description"`

	// Статистика
	AreasCount      int     `json:"areas_count"`
	ConfiguredAreas int     `json:"configured_areas"`
	ActiveCycles    int     `json:"active_cycles"`
	TotalArea       float64 `json:"total_area"`
}
type Projections interface {
	GetList(ctx context.Context, filter Filter) ([]*ListItem, error)
	GetByID(ctx context.Context, id string) (*Detail, error)
}
