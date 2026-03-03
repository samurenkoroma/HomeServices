package postgres

import (
	"database/sql"
	"samurenkoroma/services/internal/growing/application"
	"samurenkoroma/services/internal/growing/domain/cropplan/cropplan"
)

type CropPlanRepoImp struct {
}

func (c CropPlanRepoImp) Get(id cropplan.PlanID) (*cropplan.CropPlan, error) {
	//TODO implement me
	panic("implement me")
}

func (c CropPlanRepoImp) Save(plan *cropplan.CropPlan) error {
	//TODO implement me
	panic("implement me")
}

func NewCropRepo(tx *sql.Tx) application.CropPlanRepository {
	return &CropPlanRepoImp{}
}
