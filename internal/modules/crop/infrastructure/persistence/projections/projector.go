package projections

import (
	"database/sql"
)

type CropProjection struct {
	db *sql.DB
}

func NewCropProjection(db *sql.DB) *CropProjection {
	return &CropProjection{db: db}
}
