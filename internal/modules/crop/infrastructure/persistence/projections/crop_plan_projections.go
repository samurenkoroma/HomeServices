package projections

import (
	"context"
	"encoding/json"
	"samurenkoroma/services/internal/modules/crop/domain/cropplan"
)

import (
	"database/sql"
)

type cropPlanProjections struct {
	db *sql.DB
}

func (p cropPlanProjections) loadStages(ctx context.Context, planID string) ([]cropplan.GrowthStage, error) {
	rows, err := p.db.QueryContext(ctx, `
        SELECT stage_order, name, duration, recommendations
        FROM crop_crop_plan_stages
        WHERE plan_id = $1
        ORDER BY stage_order
    `, planID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stages []cropplan.GrowthStage
	for rows.Next() {
		var stage cropplan.GrowthStage
		var attrJSON []byte
		if err := rows.Scan(&stage.Order, &stage.Name, &stage.Duration, &attrJSON); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(attrJSON, &stage.Recommendations); err != nil {
			return nil, err
		}
		stages = append(stages, stage)
	}

	return stages, nil
}

func (p cropPlanProjections) GetByID(ctx context.Context, id string) (*cropplan.Detail, error) {
	query := `
SELECT cp.id, cp.name, ct.id, ct.name, v.id, v.name, cp.status, cp.version
FROM crop_crop_plans cp
    LEFT JOIN crop_crop_types  ct ON ct.id = cp.crop_type_id
    LEFT JOIN crop_varieties  v ON v.id = cp.variety_id
WHERE cp.id = $1
    `
	row := p.db.QueryRowContext(ctx, query, id)
	var item cropplan.Detail
	if err := row.Scan(&item.Id, &item.Name, &item.CropTypeId, &item.CropTypeName, &item.VarietyId, &item.VarietyName, &item.Status, &item.Version); err != nil {
		return nil, err
	}
	stages, err := p.loadStages(ctx, item.Id)
	if err != nil {
		return nil, err
	}
	item.Stages = stages
	item.RotationRules = []cropplan.RotationRule{}
	return &item, nil
}

func (p cropPlanProjections) GetList(ctx context.Context, filter cropplan.Filter) ([]*cropplan.ListItem, error) {
	query := `
SELECT cp.id, cp.name, ct.id, ct.name, v.id, v.name
FROM crop_crop_plans cp
    LEFT JOIN crop_crop_types  ct ON ct.id = cp.crop_type_id
    LEFT JOIN crop_varieties  v ON v.id = cp.variety_id
    `
	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*cropplan.ListItem
	for rows.Next() {
		var item cropplan.ListItem
		if err := rows.Scan(&item.Id, &item.Name, &item.CropTypeId, &item.CropTypeName, &item.VarietyId, &item.VarietyName); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}
	return items, nil
}

func NewCropPlanProjections(db *sql.DB) cropplan.Projections {
	return &cropPlanProjections{db: db}
}
