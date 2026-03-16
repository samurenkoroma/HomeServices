package postgres

import (
	"context"
	"database/sql"
	"samurenkoroma/services/internal/core/domain/types"
	"samurenkoroma/services/internal/modules/farm/domain/field_block"
)

type fieldBlockRepository struct {
	tx *sql.Tx
}

func NewFieldBlockRepository(tx *sql.Tx) field_block.Repository {
	return &fieldBlockRepository{tx: tx}
}

func (r *fieldBlockRepository) Save(ctx context.Context, b *field_block.FieldBlock) error {
	//query := `
	//    INSERT INTO farm_blocks (
	//        id, field_id, number, area, irrigation_type, current_crop, status, created_at, updated_at
	//    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	//    ON CONFLICT (id) DO UPDATE SET
	//        field_id = EXCLUDED.field_id,
	//        number = EXCLUDED.number,
	//        area = EXCLUDED.area,
	//        irrigation_type = EXCLUDED.irrigation_type,
	//        current_crop = EXCLUDED.current_crop,
	//        status = EXCLUDED.status,
	//        updated_at = EXCLUDED.updated_at
	//`
	//
	//// Сериализуем current_crop в JSON, если есть
	//var currentCropJSON []byte
	//var err error
	//if b.CurrentCrop != nil {
	//	currentCropJSON, err = json.Marshal(b.CurrentCrop)
	//	if err != nil {
	//		return fmt.Errorf("failed to marshal current crop: %w", err)
	//	}
	//}
	//
	//_, err = r.tx.ExecContext(ctx, query,
	//	b.ID.String(),
	//	b.FieldID,
	//	b.Number,
	//	b.Area.Float64(),
	//	string(b.IrrigationType),
	//	currentCropJSON,
	//	string(b.Status),
	//	b.CreatedAt,
	//	b.UpdatedAt,
	//)
	//
	//if err != nil {
	//	return fmt.Errorf("failed to save block: %w", err)
	//}
	//
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
