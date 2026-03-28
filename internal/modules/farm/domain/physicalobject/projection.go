package physicalobject

import (
	"context"
	"samurenkoroma/services/internal/core/spatial"
	"samurenkoroma/services/internal/modules/farm/domain/valueobject"
	"time"
)

type POFilter struct {
	Status  valueobject.Status
	Type    ObjectType
	OwnerId string
	Search  string
	Limit   int
	Offset  int
}

type POListItem struct {
	Id        string  `json:"id"`
	TypeObj   string  `json:"type"`
	Name      string  `json:"name"`
	Area      float64 `json:"area"`
	Status    string  `json:"status"`
	OwnerId   string  `json:"owner_id"`
	CreatedAt string  `json:"created_at"`
}

type PODetail struct {
	Id          string          `json:"id"`
	TypeObj     string          `json:"type"`
	Name        string          `json:"name"`
	Geometry    spatial.GeoJSON `json:"geometry"`
	Area        float64         `json:"area"`
	Status      string          `json:"status"`
	OwnerId     string          `json:"owner_id"`
	Description string          `json:"description"`
	Attributes  Attributes      `json:"attributes"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

type ObjectProjections interface {
	GetList(ctx context.Context, filter POFilter) ([]POListItem, error)
	GetByID(ctx context.Context, id string) (PODetail, error)
}
