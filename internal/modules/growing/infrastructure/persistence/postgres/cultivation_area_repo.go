package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"samurenkoroma/services/internal/infrastructure/persistence"
	"time"

	"samurenkoroma/services/internal/modules/growing/domain/cultivationarea"
)

type cultivationAreaRepository struct {
	db persistence.DBTX
}

func NewCultivationAreaRepository(db persistence.DBTX) cultivationarea.Repository {
	return &cultivationAreaRepository{db: db}
}

// Save сохраняет место выращивания
func (r *cultivationAreaRepository) Save(ctx context.Context, area cultivationarea.CultivationArea) error {
	query := `
        INSERT INTO public.growing_cultivation_areas (
            id, farm_ref_id, type, name,  area, attributes, created_at, updated_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        ON CONFLICT (id) DO UPDATE SET
            name = EXCLUDED.name,
            area = EXCLUDED.area,
            attributes = EXCLUDED.attributes,
            updated_at = EXCLUDED.updated_at
    `

	var attributesJSON []byte
	var err error
	switch a := area.(type) {
	case *cultivationarea.Bed:
		// Сохраняем атрибуты грядки
		attrs := a.GetAttributes()
		attributesJSON, err = json.Marshal(attrs)
		if err != nil {
			return fmt.Errorf("failed to marshal bed attributes: %w", err)
		}
	}

	_, err = r.db.ExecContext(ctx, query,
		area.GetId(),
		area.GetFarmRefID(),
		string(area.GetType()),
		area.GetName(),
		area.GetArea(),
		attributesJSON,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return fmt.Errorf("failed to save cultivation area: %w", err)
	}

	return nil
}

// FindByID находит место по ID
func (r *cultivationAreaRepository) FindById(ctx context.Context, id string) (cultivationarea.CultivationArea, error) {
	query := `
        SELECT id, farm_ref_id, type, name, area,  created_at, updated_at
        FROM public.growing_cultivation_areas
        WHERE id = $1
    `

	var (
		areaID    string
		farmRefID string
		areaType  string
		name      string
		areaValue float64
		createdAt time.Time
		updatedAt time.Time
	)

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&areaID, &farmRefID, &areaType, &name, &areaValue,
		&createdAt, &updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, cultivationarea.ErrAreaNotFound
		}
		return nil, fmt.Errorf("failed to find cultivation area: %w", err)
	}

	return r.hydrateArea(areaID, farmRefID, cultivationarea.AreaType(areaType), name, areaValue, createdAt, updatedAt)
}

// hydrateArea восстанавливает объект CultivationArea из данных
func (r *cultivationAreaRepository) hydrateArea(
	id, farmRefID string,
	areaType cultivationarea.AreaType,
	name string,
	areaValue float64,
	createdAt, updatedAt time.Time,
) (cultivationarea.CultivationArea, error) {
	switch areaType {
	case cultivationarea.AreaTypeField:
		field := cultivationarea.NewFieldArea(farmRefID, name, areaValue)
		field.Rehydrate(id, createdAt, updatedAt)
		return field, nil

	case cultivationarea.AreaTypeBed:
		bed := cultivationarea.NewBed(id, farmRefID, name, areaValue)
		bed.Rehydrate(createdAt, updatedAt)
		return bed, nil

	default:
		return nil, fmt.Errorf("%w: %s", cultivationarea.ErrUnknownAreaType, areaType)
	}
}
