package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/modules/farm/domain/field"
)

type fieldRepository struct {
	tx *sql.Tx
}

func NewFieldRepository(tx *sql.Tx) field.Repository {
	return &fieldRepository{tx: tx}
}

func (r *fieldRepository) Save(ctx context.Context, f *field.Field) error {
	query := `
	  INSERT INTO land_structure (
	      id, root_id, unit_type, name, geom,  status, created_at, updated_at
	  ) VALUES ($1, $2, $3, $4, ST_SetSRID(ST_GeomFromGeoJSON($5),4326), $6, $7, $8)
	  ON CONFLICT (id) DO UPDATE SET
	      name = EXCLUDED.name,
	      root_id = EXCLUDED.root_id,
	      unit_type = EXCLUDED.unit_type,
	      geom = EXCLUDED.geom,
	      status = EXCLUDED.status,
	      updated_at = EXCLUDED.updated_at
	`

	geomData, err := json.Marshal(f.Geom)
	if err != nil {
		return err
	}
	_, err = r.tx.ExecContext(ctx, query,
		f.Id.String(),
		f.Id.String(),
		types.FieldType,
		f.Name,
		geomData,
		f.Additions.Status,
		f.CreatedAt,
		f.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to save field: %w", err)
	}

	return nil
}

func (r *fieldRepository) FindByID(ctx context.Context, id types.FieldId) (*field.Field, error) {
	//query := `
	//    SELECT id, name, location, total_area, soil_type, status, blocks, created_at, updated_at
	//    FROM farm_fields
	//    WHERE id = $1
	//`
	//
	//row := r.tx.QueryRowContext(ctx, query, id.String())
	//
	//var f field.Field
	//var locationJSON []byte
	//var blocksJSON []byte
	//var soilType string
	//var status string
	//var totalArea float64
	//
	//err := row.Scan(
	//	&f.ID,
	//	&f.Name,
	//	&locationJSON,
	//	&totalArea,
	//	&soilType,
	//	&status,
	//	&blocksJSON,
	//	&f.CreatedAt,
	//	&f.UpdatedAt,
	//)
	//
	//if err == sql.ErrNoRows {
	//	return nil, field.ErrFieldNotFound
	//}
	//if err != nil {
	//	return nil, fmt.Errorf("failed to scan field: %w", err)
	//}
	//
	//// Десериализуем location
	//var location field.Location
	//if err := json.Unmarshal(locationJSON, &location); err != nil {
	//	return nil, fmt.Errorf("failed to unmarshal location: %w", err)
	//}
	//f.Location = &location
	//
	//// Десериализуем blocks
	//var blocks []field.BlockReference
	//if err := json.Unmarshal(blocksJSON, &blocks); err != nil {
	//	return nil, fmt.Errorf("failed to unmarshal blocks: %w", err)
	//}
	//f.Blocks = blocks
	//
	//f.TotalArea = types.Area(totalArea)
	//f.SoilType = field.SoilType(soilType)
	//f.Status = field.FieldStatus(status)
	//
	//return &f, nil
	return nil, nil
}

func (r *fieldRepository) FindAll(ctx context.Context) ([]*field.Field, error) {
	//query := `
	//    SELECT id, name, location, total_area, soil_type, status, blocks, created_at, updated_at
	//    FROM farm_fields
	//    ORDER BY name
	//`
	//
	//rows, err := r.tx.QueryContext(ctx, query)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to query fields: %w", err)
	//}
	//defer rows.Close()
	//
	//var fields []*field.Field
	//for rows.Next() {
	//	var f field.Field
	//	var locationJSON []byte
	//	var blocksJSON []byte
	//	var soilType string
	//	var status string
	//	var totalArea float64
	//
	//	err := rows.Scan(
	//		&f.ID,
	//		&f.Name,
	//		&locationJSON,
	//		&totalArea,
	//		&soilType,
	//		&status,
	//		&blocksJSON,
	//		&f.CreatedAt,
	//		&f.UpdatedAt,
	//	)
	//	if err != nil {
	//		return nil, fmt.Errorf("failed to scan field: %w", err)
	//	}
	//
	//	// Десериализуем (аналогично FindByID)
	//	// ...
	//
	//	f.TotalArea = types.Area(totalArea)
	//	f.SoilType = field.SoilType(soilType)
	//	f.Status = field.FieldStatus(status)
	//
	//	fields = append(fields, &f)
	//}

	//return fields, nil
	return nil, nil
}
