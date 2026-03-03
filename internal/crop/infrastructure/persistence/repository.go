package persistence

import (
	"database/sql"
	"samurenkoroma/services/internal/crop/application"
	"samurenkoroma/services/internal/crop/domain"
)

type CropPlanRepoImp struct {
}

func (c CropPlanRepoImp) Get(id domain.PlanID) (*domain.CropPlan, error) {
	//TODO implement me
	panic("implement me")
}

func (c CropPlanRepoImp) Save(plan *domain.CropPlan) error {
	//TODO implement me
	panic("implement me")
}

func NewCropRepo(tx *sql.Tx) application.CropPlanRepository {
	return &CropPlanRepoImp{}
}
