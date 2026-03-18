package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/modules/farm/domain/field_block"
)

type fieldBlockRepository struct {
	tx *sql.Tx
}

func NewFieldBlockRepository(tx *sql.Tx) field_block.Repository {
	return &fieldBlockRepository{tx: tx}
}

func (r *fieldBlockRepository) Save(ctx context.Context, fb *field_block.FieldBlock) error {
	query := `
	  INSERT INTO land_structure (
	      id, root_id, parent_id, unit_type, name, geom,  status, created_at, updated_at
	  ) VALUES ($1, $2, $3, $4,$5, ST_SetSRID(ST_GeomFromGeoJSON($6),4326), $7, $8, $9)
	  ON CONFLICT (id) DO UPDATE SET
	      name = EXCLUDED.name,
	      root_id = EXCLUDED.root_id,
	      parent_id = EXCLUDED.parent_id,
	      unit_type = EXCLUDED.unit_type,
	      geom = EXCLUDED.geom,
	      status = EXCLUDED.status,
	      updated_at = EXCLUDED.updated_at
	`

	geomData, err := json.Marshal(fb.Geom)
	if err != nil {
		return err
	}
	_, err = r.tx.ExecContext(ctx, query,
		fb.Id.String(),
		fb.ParentId.String(),
		fb.ParentId.String(),
		types.BlockType,
		fb.Name,
		geomData,
		fb.Additions.Status,
		fb.CreatedAt,
		fb.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to save field: %w", err)
	}

	return nil
}

func (r *fieldBlockRepository) FindByID(ctx context.Context, id types.FieldBlockId) (*field_block.FieldBlock, error) {
	//query := `
	//    SELECT id, field_id, number, area, irrigation_type, current_crop, status, created_at, updated_at
	//    FROM farm_blocks
	//    WHERE id = $1
	//`
	//
	//row := r.tx.QueryRowContext(ctx, query, id.String())
	//
	//var b block.Block
	//var currentCropJSON []byte
	//var irrigationType string
	//var status string
	//var area float64
	//
	//err := row.Scan(
	//	&b.ID,
	//	&b.FieldID,
	//	&b.Number,
	//	&area,
	//	&irrigationType,
	//	&currentCropJSON,
	//	&status,
	//	&b.CreatedAt,
	//	&b.UpdatedAt,
	//)
	//
	//if err == sql.ErrNoRows {
	//	return nil, block.ErrBlockNotFound
	//}
	//if err != nil {
	//	return nil, fmt.Errorf("failed to scan block: %w", err)
	//}
	//
	//b.Area = types.Area(area)
	//b.IrrigationType = block.IrrigationType(irrigationType)
	//b.Status = block.BlockStatus(status)
	//
	//// Десериализуем current_crop, если есть
	//if len(currentCropJSON) > 0 {
	//	var cropInfo block.CurrentCropInfo
	//	if err := json.Unmarshal(currentCropJSON, &cropInfo); err != nil {
	//		return nil, fmt.Errorf("failed to unmarshal current crop: %w", err)
	//	}
	//	b.CurrentCrop = &cropInfo
	//}

	return nil, nil
}

func (r *fieldBlockRepository) FindByFieldID(ctx context.Context, fieldID string) ([]*field_block.FieldBlock, error) {
	//query := `
	//    SELECT id, field_id, number, area, irrigation_type, current_crop, status, created_at, updated_at
	//    FROM farm_blocks
	//    WHERE field_id = $1
	//    ORDER BY number
	//`
	//
	//rows, err := r.tx.QueryContext(ctx, query, fieldID)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to query blocks: %w", err)
	//}
	//defer rows.Close()
	//
	//var blocks []*block.Block
	//for rows.Next() {
	//	var b block.Block
	//	var currentCropJSON []byte
	//	var irrigationType string
	//	var status string
	//	var area float64
	//
	//	err := rows.Scan(
	//		&b.ID,
	//		&b.FieldID,
	//		&b.Number,
	//		&area,
	//		&irrigationType,
	//		&currentCropJSON,
	//		&status,
	//		&b.CreatedAt,
	//		&b.UpdatedAt,
	//	)
	//	if err != nil {
	//		return nil, fmt.Errorf("failed to scan block: %w", err)
	//	}
	//
	//	b.Area = types.Area(area)
	//	b.IrrigationType = block.IrrigationType(irrigationType)
	//	b.Status = block.BlockStatus(status)
	//
	//	if len(currentCropJSON) > 0 {
	//		var cropInfo block.CurrentCropInfo
	//		json.Unmarshal(currentCropJSON, &cropInfo)
	//		b.CurrentCrop = &cropInfo
	//	}
	//
	//	blocks = append(blocks, &b)
	//}

	return nil, nil
}
