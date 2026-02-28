package postgres

import (
	"database/sql"
	"samurenkoroma/services/internal/field/application"
	"samurenkoroma/services/internal/field/domain/cropplan"
)

type CropPlanRepoImp struct {
}

func NewCropRepo(db *sql.DB) application.CropPlanRepository {
	return &CropPlanRepoImp{}
}

func (c CropPlanRepoImp) Get(id cropplan.CropPlanID) (*cropplan.CropPlan, error) {
	//TODO implement me
	panic("implement me")
}

func (c CropPlanRepoImp) Save(plan *cropplan.CropPlan) error {
	//TODO implement me
	panic("implement me")
}
