package integration

//
//import (
//	"context"
//	"fmt"
//	"time"
//
//	"samurenkoroma/services/internal/core/port/messaging"
//	"samurenkoroma/services/internal/modules/crop/domain/cropplan"
//)
//
//// GrowingClient — клиент для взаимодействия с модулем growing
//// Использует событийную шину для асинхронной коммуникации
//type GrowingClient struct {
//	eventBus messaging.EventBus
//}
//
//// NewGrowingClient создает новый клиент
//func NewGrowingClient(eventBus messaging.EventBus) *GrowingClient {
//	return &GrowingClient{
//		eventBus: eventBus,
//	}
//}
//
//// NotifyCropPlanPublished уведомляет growing о публикации плана культуры
//// Growing должен создать на основе этого шаблон выращивания (CropTemplate)
//func (c *GrowingClient) NotifyCropPlanPublished(
//	ctx context.Context,
//	plan *cropplan.CropPlan,
//) error {
//	// Создаем событие для growing модуля
//	event := CropPlanPublishedEvent{
//		PlanID:        string(plan.GetID()),
//		CropTypeID:    plan.GetCropTypeID(),
//		VarietyID:     plan.GetVarietyID(),
//		Name:          plan.GetName(),
//		Description:   plan.Description,
//		Duration:      int(plan.GetDuration()),
//		Version:       plan.GetVersion(),
//		Stages:        convertStages(plan.GetStages()),
//		Environment:   convertEnvironment(plan.GetEnvironment()),
//		Nutrients:     convertNutrients(plan.GetNutrients()),
//		RotationRules: convertRotationRules(plan.GetRotationRules()),
//		PublishedAt:   *plan.PublishedAt,
//	}
//
//	// Публикуем событие в шину
//	// Growing модуль подписан на это событие
//	if err := c.eventBus.Publish(ctx, event); err != nil {
//		return fmt.Errorf("failed to publish crop plan published event: %w", err)
//	}
//
//	return nil
//}
//
//// NotifyCropPlanDeprecated уведомляет growing о деактивации плана
//func (c *GrowingClient) NotifyCropPlanDeprecated(
//	ctx context.Context,
//	planID string,
//	reason string,
//) error {
//	event := CropPlanDeprecatedEvent{
//		PlanID:       planID,
//		Reason:       reason,
//		DeprecatedAt: time.Now(),
//	}
//
//	if err := c.eventBus.Publish(ctx, event); err != nil {
//		return fmt.Errorf("failed to publish crop plan deprecated event: %w", err)
//	}
//
//	return nil
//}
//
//// NotifyCropPlanVersionCreated уведомляет о создании новой версии
//func (c *GrowingClient) NotifyCropPlanVersionCreated(
//	ctx context.Context,
//	originalPlanID string,
//	newPlanID string,
//	version int,
//) error {
//	event := CropPlanVersionCreatedEvent{
//		OriginalPlanID: originalPlanID,
//		NewPlanID:      newPlanID,
//		Version:        version,
//		CreatedAt:      time.Now(),
//	}
//
//	if err := c.eventBus.Publish(ctx, event); err != nil {
//		return fmt.Errorf("failed to publish crop plan version created event: %w", err)
//	}
//
//	return nil
//}
//
//// ========== Конвертеры ==========
//
//func convertStages(stages []cropplan.GrowthStage) []GrowthStageDTO {
//	result := make([]GrowthStageDTO, len(stages))
//	for i, s := range stages {
//		result[i] = GrowthStageDTO{
//			Order:         s.Order,
//			Name:          s.Name,
//			Duration:      s.Duration,
//			MinTemp:       s.MinTemp,
//			MaxTemp:       s.MaxTemp,
//			OptimalTemp:   s.OptimalTemp,
//			WaterPerDay:   s.WaterPerDay,
//			NitrogenReq:   s.NitrogenReq,
//			PhosphorusReq: s.PhosphorusReq,
//			PotassiumReq:  s.PotassiumReq,
//		}
//	}
//	return result
//}
//
//func convertEnvironment(env cropplan.EnvironmentalRequirements) EnvironmentDTO {
//	return EnvironmentDTO{
//		MinTemp:               env.MinTemp,
//		MaxTemp:               env.MaxTemp,
//		MinHumidity:           env.MinHumidity,
//		MaxHumidity:           env.MaxHumidity,
//		MinPH:                 env.MinPH,
//		MaxPH:                 env.MaxPH,
//		SoilTypes:             env.SoilTypes,
//		MinLightHours:         env.MinLightHours,
//		TotalWaterRequirement: env.TotalWaterRequirement,
//	}
//}
//
//func convertNutrients(nut cropplan.NutrientRequirements) NutrientsDTO {
//	return NutrientsDTO{
//		Nitrogen:   nut.Nitrogen,
//		Phosphorus: nut.Phosphorus,
//		Potassium:  nut.Potassium,
//		Calcium:    nut.Calcium,
//		Magnesium:  nut.Magnesium,
//		Sulfur:     nut.Sulfur,
//	}
//}
//
//func convertRotationRules(rules []cropplan.RotationRule) []RotationRuleDTO {
//	result := make([]RotationRuleDTO, len(rules))
//	for i, r := range rules {
//		result[i] = RotationRuleDTO{
//			PredecessorCropTypeID: r.PredecessorCropTypeID,
//			MinYears:              r.MinYears,
//			Recommended:           r.Recommended,
//			Notes:                 r.Notes,
//		}
//	}
//	return result
//}
//
//// ========== DTO для событий ==========
//
//// CropPlanPublishedEvent — событие публикации плана культуры
//type CropPlanPublishedEvent struct {
//	PlanID        string            `json:"plan_id"`
//	CropTypeID    string            `json:"crop_type_id"`
//	VarietyID     *string           `json:"variety_id"`
//	Name          string            `json:"name"`
//	Description   string            `json:"description"`
//	Duration      int               `json:"duration"`
//	Version       int               `json:"version"`
//	Stages        []GrowthStageDTO  `json:"stages"`
//	Environment   EnvironmentDTO    `json:"environment"`
//	Nutrients     NutrientsDTO      `json:"nutrients"`
//	RotationRules []RotationRuleDTO `json:"rotation_rules"`
//	PublishedAt   time.Time         `json:"published_at"`
//}
//
//func (e CropPlanPublishedEvent) EventName() string {
//	return "crop.plan.published"
//}
//
//// CropPlanDeprecatedEvent — событие деактивации плана
//type CropPlanDeprecatedEvent struct {
//	PlanID       string    `json:"plan_id"`
//	Reason       string    `json:"reason"`
//	DeprecatedAt time.Time `json:"deprecated_at"`
//}
//
//func (e CropPlanDeprecatedEvent) EventName() string {
//	return "crop.plan.deprecated"
//}
//
//// CropPlanVersionCreatedEvent — событие создания новой версии
//type CropPlanVersionCreatedEvent struct {
//	OriginalPlanID string    `json:"original_plan_id"`
//	NewPlanID      string    `json:"new_plan_id"`
//	Version        int       `json:"version"`
//	CreatedAt      time.Time `json:"created_at"`
//}
//
//func (e CropPlanVersionCreatedEvent) EventName() string {
//	return "crop.plan.version_created"
//}
//
//// ========== DTO для этапов и требований ==========
//
//type GrowthStageDTO struct {
//	Order         int     `json:"order"`
//	Name          string  `json:"name"`
//	Duration      int     `json:"duration"`
//	MinTemp       float64 `json:"min_temp"`
//	MaxTemp       float64 `json:"max_temp"`
//	OptimalTemp   float64 `json:"optimal_temp"`
//	WaterPerDay   float64 `json:"water_per_day"`
//	NitrogenReq   float64 `json:"nitrogen_req"`
//	PhosphorusReq float64 `json:"phosphorus_req"`
//	PotassiumReq  float64 `json:"potassium_req"`
//}
//
//type EnvironmentDTO struct {
//	MinTemp               float64  `json:"min_temp"`
//	MaxTemp               float64  `json:"max_temp"`
//	MinHumidity           float64  `json:"min_humidity"`
//	MaxHumidity           float64  `json:"max_humidity"`
//	MinPH                 float64  `json:"min_ph"`
//	MaxPH                 float64  `json:"max_ph"`
//	SoilTypes             []string `json:"soil_types"`
//	MinLightHours         int      `json:"min_light_hours"`
//	TotalWaterRequirement float64  `json:"total_water_requirement"`
//}
//
//type NutrientsDTO struct {
//	Nitrogen   float64 `json:"nitrogen"`
//	Phosphorus float64 `json:"phosphorus"`
//	Potassium  float64 `json:"potassium"`
//	Calcium    float64 `json:"calcium"`
//	Magnesium  float64 `json:"magnesium"`
//	Sulfur     float64 `json:"sulfur"`
//}
//
//type RotationRuleDTO struct {
//	PredecessorCropTypeID string `json:"predecessor_crop_type_id"`
//	MinYears              int    `json:"min_years"`
//	Recommended           bool   `json:"recommended"`
//	Notes                 string `json:"notes"`
//}
