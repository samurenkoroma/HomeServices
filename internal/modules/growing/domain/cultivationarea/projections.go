package cultivationarea

import (
	"context"
)

type Filter struct {
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

type ListItem struct {
	Id              string  `json:"id"`
	FarmRefId       string  `json:"farmRefId"`
	Type            string  `json:"start_date"`
	Name            string  `json:"name"`
	Geometry        string  `json:"geometry"`
	Area            float64 `json:"area"`
	ParentId        string  `json:"parentId"`
	CurrentSeasonId string  `json:"currentSeasonId"`
	CreatedAt       string  `json:"created_at"`
	IsConfigured    bool    `json:"is_configured"`
}

//	type Detail struct {
//		ID          string    `json:"id"`
//		Name        string    `json:"name"`
//		StartDate   time.Time `json:"start_date"`
//		EndDate     time.Time `json:"end_date"`
//		Status      string    `json:"status"`
//		Description string    `json:"description"`
//
//		// Статистика
//		AreasCount      int     `json:"areas_count"`
//		ConfiguredAreas int     `json:"configured_areas"`
//		ActiveCycles    int     `json:"active_cycles"`
//		TotalArea       float64 `json:"total_area"`
//	}
type Projections interface {
	GetList(ctx context.Context, filter Filter) ([]*ListItem, error)
	//GetByID(ctx context.Context, id string) (*Detail, error)
}
