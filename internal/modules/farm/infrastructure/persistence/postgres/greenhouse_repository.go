package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/modules/farm/domain/greenhouse"
	"samurenkoroma/services/internal/modules/farm/domain/valueobject"
)

type greenhouseRepository struct {
	tx *sql.Tx
}

func (gr greenhouseRepository) FindByID(ctx context.Context, id types.GreenhouseId) (*greenhouse.Greenhouse, error) {
	//TODO implement me
	panic("implement me")
}

func (gr greenhouseRepository) Save(ctx context.Context, g *greenhouse.Greenhouse) error {
	query := `
	  INSERT INTO land_structure (
	      id, root_id, unit_type, name, geom,  status, properties, created_at, updated_at
	  ) VALUES ($1, $2, $3, $4, ST_SetSRID(ST_GeomFromGeoJSON($5),4326), $6, $7, $8, $9)
	  ON CONFLICT (id) DO UPDATE SET
	      name = EXCLUDED.name,
	      root_id = EXCLUDED.root_id,
	      unit_type = EXCLUDED.unit_type,
	      geom = EXCLUDED.geom,
	      status = EXCLUDED.status,
	      updated_at = EXCLUDED.updated_at
	`

	geomData, err := json.Marshal(g.Geom)
	if err != nil {
		return err
	}

	properties, err := g.Dimension.Marshall()
	if err != nil {
		return err
	}

	_, err = gr.tx.ExecContext(ctx, query,
		g.Id.String(),
		g.Id.String(),
		types.GreenhouseType,
		g.Name,
		geomData,
		valueobject.AreaStatusActive,
		properties,
		g.CreatedAt,
		g.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to save field: %w", err)
	}

	return nil
}

func (gr greenhouseRepository) FindAll(ctx context.Context) ([]*greenhouse.Greenhouse, error) {
	//TODO implement me
	panic("implement me")
}

func NewGreenhouseRepository(tx *sql.Tx) greenhouse.Repository {
	return greenhouseRepository{tx: tx}
}
