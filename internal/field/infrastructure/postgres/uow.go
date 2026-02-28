package postgres

import (
	"database/sql"
	"samurenkoroma/services/internal/field/application"
)

type PgUow struct {
	tx       *sql.Tx
	landRepo application.LandUnitRepository
	cropRepo application.CropPlanRepository
}

func NewUow(db *sql.DB) (*PgUow, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	return &PgUow{
		tx:       tx,
		landRepo: NewLandRepo(db),
		cropRepo: NewCropRepo(db),
	}, nil
}

func (u *PgUow) LandUnits() application.LandUnitRepository {
	return u.landRepo
}

func (u *PgUow) CropPlans() application.CropPlanRepository {
	return u.cropRepo
}

func (u *PgUow) Commit() error {
	return u.tx.Commit()
}

func (u *PgUow) Rollback() error {
	return u.tx.Rollback()
}
