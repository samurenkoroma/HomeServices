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
	Type            string  `json:"type"`
	Name            string  `json:"name"`
	Geometry        string  `json:"geometry"`
	Area            float64 `json:"area"`
	ParentId        *string `json:"parentId"`
	CurrentSeasonId string  `json:"currentSeasonId"`
	CreatedAt       string  `json:"created_at"`
	IsConfigured    bool    `json:"is_configured"`
}

type Detail struct {
	Id              string  `json:"id"`
	FarmRefId       string  `json:"farmRefId"`
	Type            string  `json:"type"`
	Name            string  `json:"name"`
	Geometry        string  `json:"geometry"`
	Area            float64 `json:"area"`
	ParentId        *string `json:"parentId"`
	CurrentSeasonId string  `json:"currentSeasonId"`
	CreatedAt       string  `json:"created_at"`
	IsConfigured    bool    `json:"is_configured"`
}
type Projections interface {
	GetList(ctx context.Context, filter Filter) ([]*ListItem, error)
	GetByID(ctx context.Context, id string) (*Detail, error)
}
