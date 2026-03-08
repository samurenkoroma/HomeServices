package postgres

import (
	"context"
	"database/sql"
	"samurenkoroma/services/internal/modules/growing/domain"
	"samurenkoroma/services/internal/modules/growing/domain/facility"
	"time"

	"github.com/google/uuid"
)

type facilityReadRepository struct{ db *sql.DB }

func NewFacilityReadRepository(db *sql.DB) domain.FacilityReadRepository {
	return &facilityReadRepository{
		db: db,
	}
}

func (r *facilityReadRepository) GetOverview(ctx context.Context, id string) (*domain.FacilityOverviewDTO, error) {

	const query = `
	SELECT
		id,
		name,
		unit_type,
		length,
    	width
	FROM land_structure
	WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, id)

	dto := &domain.FacilityOverviewDTO{}

	if err := row.Scan(
		&dto.ID,
		&dto.Name,
		&dto.Type,
		&dto.Length,
		&dto.Width,
	); err != nil {
		return nil, err
	}

	return dto, nil
}

func (r *facilityReadRepository) GetList(ctx context.Context, params domain.FacilitiesListParams) (*domain.FacilitiesListDTO, error) {

	dto := &domain.FacilitiesListDTO{
		Items: []domain.FacilitiesListItemDTO{{
			Id:          uuid.New().String(),
			Name:        "Ферма Зеленая Долина",
			Type:        string(facility.FieldFacility),
			Area:        176.5,
			ActiveCrops: 5,
			Status:      "excellent",
			YieldTrend:  12.5,
			Location:    "Айова, США",
			Thumbnail:   nil,
			UpdatedAt:   time.Now(),
		},
			{
				Id:          uuid.New().String(),
				Name:        "Теплица Восход",
				Type:        string(facility.GreenhouseFacility),
				Area:        73.2,
				ActiveCrops: 8,
				Status:      "attention",
				YieldTrend:  5.1,
				Location:    "Калифорния, США",
				Thumbnail:   nil,
				UpdatedAt:   time.Now(),
			},
			{
				Id:          "facility-3",
				Name:        "Ферма Северные Равнины",
				Type:        "FIELD",
				Area:        130.0,
				ActiveCrops: 4,
				Status:      "good",
				YieldTrend:  8.3,
				Location:    "Небраска, США",
				Thumbnail:   nil,
				UpdatedAt:   time.Now(),
			}},
		Total: 45,
		Page:  1,

		Limit:      10,
		TotalPages: 5,
	}
	return dto, nil

}
