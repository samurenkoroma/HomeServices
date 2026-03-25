package projections

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"samurenkoroma/services/internal/modules/farm/domain/physicalobject"
	"samurenkoroma/services/internal/modules/growing/domain/season"
)

type poProjection struct {
	db *sql.DB
}

func (f poProjection) GetList(ctx context.Context, filter physicalobject.POFilter) ([]physicalobject.POListItem, error) {
	query := `
	SELECT
		po.id, po.type, po.name,  po.area, po.status, po.owner_id, po.created_at
	FROM farm_physical_objects po
	WHERE ($1 = '' OR po.status = $1)
	  AND ($2 = '' OR po.type = $2)
	  AND ($3 = '' OR po.type = $3)
	  AND ($4 = '' OR po.name ILIKE '%' || $4 || '%')
	GROUP BY po.id
	ORDER BY po.name 
	LIMIT $5 OFFSET $6
		`

	rows, err := f.db.QueryContext(ctx, query,
		filter.Status, filter.Type, filter.OwnerId, filter.Search, filter.Limit, filter.Offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query seasons: %w", err)
	}
	defer rows.Close()

	var items []physicalobject.POListItem
	for rows.Next() {
		var item physicalobject.POListItem
		if err := rows.Scan(&item.Id, &item.TypeObj, &item.Name, &item.Area, &item.Status, &item.OwnerId, &item.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan item: %w", err)
		}
		items = append(items, item)
	}

	return items, nil
}

func (f poProjection) GetByID(ctx context.Context, id string) (physicalobject.PODetail, error) {
	query := `
        SELECT 
            id, type, name, ST_AsGeoJSON(geometry), area, status, owner_id, description, attributes, created_at, updated_at
           
        FROM farm_physical_objects po
        WHERE po.id = $1
        GROUP BY po.id
    `
	var attrJSON []byte
	var geomJSON string
	var detail physicalobject.PODetail
	err := f.db.QueryRowContext(ctx, query, id).Scan(
		&detail.Id, &detail.TypeObj, &detail.Name, &geomJSON, &detail.Area, &detail.Status, &detail.OwnerId, &detail.Description,
		&attrJSON, &detail.CreatedAt, &detail.UpdatedAt)

	// Парсим GeoJSON
	//var geom spatial.GeoJSON
	if err := json.Unmarshal([]byte(geomJSON), &detail.Geometry); err != nil {
		return physicalobject.PODetail{}, err
	}
	//detail.Geometry = geom
	var attrs physicalobject.Attributes
	if err := json.Unmarshal(attrJSON, &attrs); err != nil {
		return physicalobject.PODetail{}, err
	}
	detail.Attributes = attrs
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return physicalobject.PODetail{}, season.ErrSeasonNotFound
		}
		return physicalobject.PODetail{}, fmt.Errorf("failed to get season detail: %w", err)
	}

	return detail, nil
}

func NewPoProjection(db *sql.DB) physicalobject.POProjection {
	return &poProjection{db: db}
}
