package cropplan

import "samurenkoroma/services/internal/core/domain/event"

// CropPlanCreated - план создан
type CropPlanCreated struct {
	event.BaseEvent
	PlanID     string
	CropTypeID string
	Name       string
}

func (e CropPlanCreated) EventName() string { return "crop.plan.created" }

// CropPlanPublished - план опубликован
type CropPlanPublished struct {
	event.BaseEvent
	PlanID  string
	Version int
	Name    string
}

func (e CropPlanPublished) EventName() string { return "crop.plan.published" }

// CropPlanDeprecated - план деактивирован
type CropPlanDeprecated struct {
	event.BaseEvent
	PlanID string
	Reason string
}

func (e CropPlanDeprecated) EventName() string { return "crop.plan.deprecated" }

// CropPlanVersionCreated - создана новая версия
type CropPlanVersionCreated struct {
	event.BaseEvent
	OriginalPlanID string
	NewPlanID      string
	Version        int
}

func (e CropPlanVersionCreated) EventName() string { return "crop.plan.version_created" }

// GrowthStageAdded - добавлен этап роста
type GrowthStageAdded struct {
	event.BaseEvent
	PlanID string
	Order  int
	Name   string
	Stage  GrowthStage
}

func (e GrowthStageAdded) EventName() string { return "crop.plan.stage_added" }

// RotationRuleAdded - добавлено правило севооборота
type RotationRuleAdded struct {
	event.BaseEvent
	PlanID                string
	PredecessorCropTypeID string
	MinYears              int
}

func (e RotationRuleAdded) EventName() string { return "crop.plan.rotation_rule_added" }

// CropTypeCreated - тип культуры создан
type CropTypeCreated struct {
	event.BaseEvent
	CropTypeID string
	Name       string
	Category   string
}

func (e CropTypeCreated) EventName() string { return "crop.type.created" }

// VarietyCreated - сорт создан
type VarietyCreated struct {
	event.BaseEvent
	VarietyID  string
	CropTypeID string
	Name       string
}

func (e VarietyCreated) EventName() string { return "crop.variety.created" }
