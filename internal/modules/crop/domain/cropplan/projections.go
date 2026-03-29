package cropplan

import (
	"context"
)

type Filter struct {
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

type ListItem struct {
	Id            string         `json:"id"`
	CropTypeId    string         `json:"crop_type_id"`
	CropTypeName  string         `json:"crop_type_name"`
	VarietyId     *string        `json:"variety_id,omitempty"`
	VarietyName   *string        `json:"variety_name,omitempty"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	Duration      float64        `json:"duration"`
	Version       int            `json:"version"`
	Status        string         `json:"status"`
	Stages        []GrowthStage  `json:"stages"`
	RotationRules []RotationRule `json:"rotation_rules"`
	CreatedBy     string         `json:"created_by"`
	CreatedAt     string         `json:"created_at"`
	PublishedAt   string         `json:"published_at"`
}
type Detail struct {
	Id            string         `json:"id"`
	CropTypeId    string         `json:"crop_type_id"`
	CropTypeName  string         `json:"crop_type_name"`
	VarietyId     *string        `json:"variety_id,omitempty"`
	VarietyName   *string        `json:"variety_name,omitempty"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	Duration      float64        `json:"duration"`
	Version       int            `json:"version"`
	Status        string         `json:"status"`
	Stages        []GrowthStage  `json:"stages"`
	RotationRules []RotationRule `json:"rotation_rules"`
	CreatedBy     string         `json:"created_by"`
	CreatedAt     string         `json:"created_at"`
	PublishedAt   string         `json:"published_at"`
}

type Projections interface {
	GetList(ctx context.Context, filter Filter) ([]*ListItem, error)
	GetByID(ctx context.Context, id string) (*Detail, error)
}
