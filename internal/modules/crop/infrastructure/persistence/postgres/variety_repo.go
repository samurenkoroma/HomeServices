package postgres

import (
	"context"
	"database/sql"
	"fmt"
	sqlErrors "samurenkoroma/services/internal/core/errors"
	"samurenkoroma/services/internal/modules/crop/domain/croptype"
	"samurenkoroma/services/internal/modules/crop/domain/variety"
)

type varietyRepo struct {
	tx *sql.Tx
}

func NewVarietyRepository(tx *sql.Tx) variety.Repository {
	return &varietyRepo{tx: tx}
}
func (repo *varietyRepo) Save(ctx context.Context, obj *variety.Variety) error {
	query := `INSERT INTO varieties ( 
                     id,crop_type_id, name, description, attributes, is_active, created_at, updated_at 
                     ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
                       ON CONFLICT (id) DO UPDATE SET
						crop_type_id = EXCLUDED.crop_type_id,
						name = EXCLUDED.name,
						description = EXCLUDED.description,
						attributes = EXCLUDED.attributes,
						is_active = EXCLUDED.is_active,
						updated_at = EXCLUDED.updated_at`

	attrData, err := obj.Attributes.Marshal()
	if err != nil {
		return err
	}
	_, err = repo.tx.ExecContext(ctx, query,
		obj.Id,
		obj.CropTypeID,
		obj.Name,
		obj.Description,
		attrData,
		obj.IsActive,
		obj.CreatedAt,
		obj.UpdatedAt,
	)
	if err != nil {
		if sqlErrors.IsUniqueViolation(err) {
			return croptype.ErrVarietyAlreadyExists
		}
		return fmt.Errorf("failed to save variety: %w", err)
	}
	return nil
}

func (repo *varietyRepo) GetByID(ctx context.Context, id variety.VarietyID) (*variety.Variety, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *varietyRepo) List(ctx context.Context) ([]*variety.Variety, error) {
	//TODO implement me
	panic("implement me")
}
