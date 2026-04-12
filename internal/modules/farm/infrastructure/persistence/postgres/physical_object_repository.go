package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/spatial"
	"samurenkoroma/services/internal/modules/farm/domain/physicalobject"
	"samurenkoroma/services/internal/modules/farm/domain/valueobject"
	"time"
)

type physicalObjectRepository struct {
	tx *sql.Tx
}

func NewPhysicalObjectRepository(tx *sql.Tx) physicalobject.Repository {
	return &physicalObjectRepository{tx: tx}
}
func (r *physicalObjectRepository) FindAll(ctx context.Context) ([]*physicalobject.PhysicalObject, error) {
	query := `
        SELECT id, name, area
        FROM public.farm_physical_objects
        WHERE status = 'active'
        ORDER BY name
    `

	rows, err := r.tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fields []*physicalobject.PhysicalObject
	for rows.Next() {
		var (
			id   string
			name string
			area float64
		)

		if err := rows.Scan(&id, &name, &area); err != nil {
			return nil, err
		}

		// Конвертация в field.Field
		field := physicalobject.RehydrateField(id, name, area)
		fields = append(fields, field)
	}

	return fields, nil
}

func (r *physicalObjectRepository) Delete(ctx context.Context, obj *physicalobject.PhysicalObject) error {
	query := `UPDATE farm_physical_objects  set delete_at = $2, status = $3 WHERE id = $1`
	_, err := r.tx.ExecContext(ctx, query, obj.Id.String(), obj.DeletedAt, obj.Status)

	return err

}
func (r *physicalObjectRepository) Save(ctx context.Context, obj *physicalobject.PhysicalObject) error {
	query := `
        INSERT INTO farm_physical_objects (
            id, type, name, geometry, status, owner_id, description, attributes, created_at, updated_at, area
        ) VALUES ($1, $2, $3, ST_SetSRID(ST_GeomFromGeoJSON($4),4326), $5, $6, $7, $8, $9, $10, $11)
        ON CONFLICT (id) DO UPDATE SET
            name = EXCLUDED.name,
            geometry = EXCLUDED.geometry,
            status = EXCLUDED.status,
            description = EXCLUDED.description,
            attributes = EXCLUDED.attributes,
            updated_at = EXCLUDED.updated_at
    `

	geomData, err := json.Marshal(obj.Geometry)
	if err != nil {
		return err
	}

	attrData, err := obj.Attributes.Marshal()
	if err != nil {
		return err
	}

	_, err = r.tx.ExecContext(ctx, query,
		obj.Id.String(),
		string(obj.Type),
		obj.Name,
		geomData,
		obj.Status,
		obj.OwnerID,
		obj.Description,
		attrData,
		obj.CreatedAt,
		obj.UpdatedAt,
		obj.Area,
	)

	return err
}

func (r *physicalObjectRepository) FindByID(ctx context.Context, id physicalobject.PhysicalObjectID) (*physicalobject.PhysicalObject, error) {
	query := `
        SELECT 
            id, type, name, ST_AsGeoJSON(geometry), 
            status, owner_id, description, attributes,
            created_at, updated_at
        FROM farm_physical_objects
        WHERE id = $1
    `

	var (
		objID       string
		objType     string
		name        string
		geomJSON    string
		status      string
		ownerID     string
		description sql.NullString
		attrJSON    []byte
		createdAt   time.Time
		updatedAt   time.Time
	)

	err := r.tx.QueryRowContext(ctx, query, id.String()).Scan(
		&objID, &objType, &name, &geomJSON,
		&status, &ownerID, &description, &attrJSON,
		&createdAt, &updatedAt,
	)

	if err != nil {
		return nil, err
	}

	// Парсим GeoJSON
	var geom spatial.GeoJSON
	if err := json.Unmarshal([]byte(geomJSON), &geom); err != nil {
		return nil, err
	}

	// Парсим атрибуты
	var attrs physicalobject.Attributes
	if err := json.Unmarshal(attrJSON, &attrs); err != nil {
		return nil, err
	}

	// Инициализируем Metadata если nil
	if attrs.Metadata == nil {
		attrs.Metadata = make(map[string]interface{})
	}

	obj := &physicalobject.PhysicalObject{
		Entity:      aggregate.NewEntity(physicalobject.PhysicalObjectID(objID)),
		Type:        physicalobject.ObjectType(objType),
		Name:        name,
		Geometry:    geom,
		Status:      valueobject.Status(status),
		OwnerID:     ownerID,
		Description: description.String,
		Attributes:  attrs,
	}

	return obj, nil
}

func (r *physicalObjectRepository) FindByType(ctx context.Context, objType physicalobject.ObjectType) ([]*physicalobject.PhysicalObject, error) {
	query := `
        SELECT id, name, ST_AsGeoJSON(geometry), status, attributes
        FROM farm_physical_objects
        WHERE type = $1 AND status = 'active'
        ORDER BY name
    `

	rows, err := r.tx.QueryContext(ctx, query, string(objType))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var objects []*physicalobject.PhysicalObject
	for rows.Next() {
		var (
			id       string
			name     string
			geomJSON string
			status   string
			attrJSON []byte
		)

		if err := rows.Scan(&id, &name, &geomJSON, &status, &attrJSON); err != nil {
			return nil, err
		}

		// Конвертация (упрощенно)
		obj := &physicalobject.PhysicalObject{
			Entity: aggregate.NewEntity(physicalobject.PhysicalObjectID(id)),
			Type:   objType,
			Name:   name,
			Status: valueobject.Status(status),
		}

		objects = append(objects, obj)
	}

	return objects, nil
}

// FindInBounds Геопространственные запросы теперь очень простые!
func (r *physicalObjectRepository) FindInBounds(ctx context.Context, bounds spatial.BoundingBox) ([]*physicalobject.PhysicalObject, error) {
	query := `
        SELECT id, type, name, ST_AsGeoJSON(geometry), status
        FROM farm_physical_objects
        WHERE geometry && ST_MakeEnvelope($1, $2, $3, $4, 4326)
          AND status = 'active'
        ORDER BY type, name
    `

	rows, err := r.tx.QueryContext(ctx, query,
		bounds.MinLon, bounds.MinLat, bounds.MaxLon, bounds.MaxLat,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var objects []*physicalobject.PhysicalObject
	for rows.Next() {
		var (
			id       string
			objType  string
			name     string
			geomJSON string
			status   string
		)

		if err := rows.Scan(&id, &objType, &name, &geomJSON, &status); err != nil {
			return nil, err
		}

		// Создаем объект (упрощенно)
		objects = append(objects, &physicalobject.PhysicalObject{
			Entity: aggregate.NewEntity(physicalobject.PhysicalObjectID(id)),
			Type:   physicalobject.ObjectType(objType),
			Name:   name,
			Status: valueobject.Status(status),
		})
	}

	return objects, nil
}
