package cropplan

//type GetCurrentPhenologyQuery struct {
//	PlanID string `json:"plan_id"`
//}
//
//type GetCurrentPhenologyResponse struct {
//	PlanID      string `json:"plan_id"`
//	PlanName    string `json:"plan_name"`
//	VarietyID   string `json:"variety_id"`
//	VarietyName string `json:"variety_name"`
//
//	// GDD
//	AccumulatedGDD     float64 `json:"accumulated_gdd"`
//	RequiredGDDForNext float64 `json:"required_gdd_for_next"`
//
//	// Текущая фаза
//	CurrentPhaseCode string  `json:"current_phase_code"`
//	CurrentPhaseName string  `json:"current_phase_name"`
//	ProgressPercent  float64 `json:"progress_percent"`
//
//	// Прогноз
//	EstimatedDaysToNextPhase int     `json:"estimated_days_to_next_phase"`
//	EstimatedHarvestDate     *string `json:"estimated_harvest_date,omitempty"`
//
//	// Отклонения
//	DeviationDays   int    `json:"deviation_days"`
//	DeviationReason string `json:"deviation_reason"`
//
//	// Критичность
//	IsCritical bool `json:"is_critical"`
//
//	// Рекомендации
//	RecommendedActions []phenology.RecommendedAction `json:"recommended_actions"`
//
//	// Доступные этапы для выполнения
//	AvailableStages []dto.StageDTO `json:"available_stages"`
//}
//
//func (h *QueryHandler) GetCurrentPhenolog(ctx context.Context, query any) (any, error) {
//	q, ok := query.(*GetCurrentPhenologyQuery)
//	if !ok {
//		return nil, errors.New("invalid query type")
//	}
//
//	// Получаем план
//	plan, err := h.provider.CropPlans().GetByID(ctx, q.PlanID)
//	if err != nil {
//		return nil, fmt.Errorf("plan not found: %w", err)
//	}
//
//	// Получаем сорт
//	variety, err := h.provider.Catalog().GetVariety(ctx, plan.VarietyID)
//	if err != nil {
//		return nil, fmt.Errorf("variety not found: %w", err)
//	}
//
//	// Получаем текущую фенологию
//	current, err := h.provider.PhenologyService().GetCurrentPhenology(
//		ctx,
//		plan.ID(),
//		plan.Variety().GetId(),
//		plan.PlantingDate(),
//		30.5,
//		40.2,
//		//plan.Latitude(),
//		//plan.Longitude(),
//	)
//	if err != nil {
//		return nil, fmt.Errorf("failed to get phenology: %w", err)
//	}
//
//	// Получаем BBCH код из текущей фазы
//	currentBBCH := extractBBCHCode(current.CurrentPhaseCode)
//
//	// Получаем этапы, доступные для выполнения
//	applicableStages := plan.plan.VarietyID(currentBBCH)
//	availableStages := make([]dto.StageDTO, len(applicableStages))
//	for i, stage := range applicableStages {
//		availableStages[i] = dto.ToStageDTO(stage, currentBBCH)
//	}
//
//	// Форматируем дату сбора
//	var harvestDateStr *string
//	if current.EstimatedHarvestDate != nil {
//		str := current.EstimatedHarvestDate.Format("2006-01-02")
//		harvestDateStr = &str
//	}
//
//	return &GetCurrentPhenologyResponse{
//		PlanID:                   plan.ID(),
//		PlanName:                 plan.Name(),
//		VarietyID:                variety.ID,
//		VarietyName:              variety.Name,
//		AccumulatedGDD:           current.AccumulatedGDD,
//		RequiredGDDForNext:       current.RequiredGDDForNext,
//		CurrentPhaseCode:         current.CurrentPhaseCode,
//		CurrentPhaseName:         current.CurrentPhaseName,
//		ProgressPercent:          current.ProgressPercent,
//		EstimatedDaysToNextPhase: current.EstimatedDaysToNextPhase,
//		EstimatedHarvestDate:     harvestDateStr,
//		DeviationDays:            current.DeviationDays,
//		DeviationReason:          current.DeviationReason,
//		IsCritical:               current.IsCritical,
//		RecommendedActions:       current.RecommendedActions,
//		AvailableStages:          availableStages,
//	}, nil
//}
//
//// extractBBCHCode извлекает числовой код из "BBCH-61"
//func extractBBCHCode(code string) int {
//	var bbch int
//	fmt.Sscanf(code, "BBCH-%d", &bbch)
//	return bbch
//}
