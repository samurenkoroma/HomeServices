package inmemory

import (
	"context"
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/cropplan"
	"sync"
	"time"
)

// CropPlanRepo in-memory реализация репозитория планов
type CropPlanRepo struct {
	mu    sync.RWMutex
	plans map[string]*cropplan.CropPlan
}

func NewCropPlanRepo() cropplan.Repository {
	return &CropPlanRepo{
		plans: make(map[string]*cropplan.CropPlan),
	}
}

// Save сохраняет новый план
func (r *CropPlanRepo) Save(ctx context.Context, plan *cropplan.CropPlan) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.plans[plan.ID()]; exists {
		return ErrDuplicatePlan
	}

	r.plans[plan.ID()] = plan
	return nil
}

// Update обновляет существующий план
func (r *CropPlanRepo) Update(ctx context.Context, plan *cropplan.CropPlan) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.plans[plan.ID()]; !exists {
		return cropplan.ErrPlanNotFound
	}

	r.plans[plan.ID()] = plan
	return nil
}

// FindByID находит план по ID
func (r *CropPlanRepo) FindByID(ctx context.Context, id string) (*cropplan.CropPlan, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	plan, exists := r.plans[id]
	if !exists {
		return nil, cropplan.ErrPlanNotFound
	}

	return plan, nil
}

// Delete удаляет план
func (r *CropPlanRepo) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.plans[id]; !exists {
		return cropplan.ErrPlanNotFound
	}

	delete(r.plans, id)
	return nil
}

// FindByBed находит все планы на грядке
func (r *CropPlanRepo) FindByArea(ctx context.Context, bedID string) ([]*cropplan.CropPlan, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*cropplan.CropPlan
	for _, plan := range r.plans {
		if plan.Area().GetId() == bedID {
			result = append(result, plan)
		}
	}
	return result, nil
}

// FindByVariety находит все планы по сорту
func (r *CropPlanRepo) FindByVariety(ctx context.Context, varietyID string) ([]*cropplan.CropPlan, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*cropplan.CropPlan
	for _, plan := range r.plans {
		if plan.Variety().GetId() == varietyID {
			result = append(result, plan)
		}
	}
	return result, nil
}

// FindByStatus находит все планы по статусу
func (r *CropPlanRepo) FindByStatus(ctx context.Context, status cropplan.Status) ([]*cropplan.CropPlan, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*cropplan.CropPlan
	for _, plan := range r.plans {
		if plan.Status() == status {
			result = append(result, plan)
		}
	}
	return result, nil
}

// FindByAssignedTo находит планы, назначенные агроному
func (r *CropPlanRepo) FindByAssignedTo(ctx context.Context, userID string) ([]*cropplan.CropPlan, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*cropplan.CropPlan
	for _, plan := range r.plans {
		if plan.AssignedTo() == userID {
			result = append(result, plan)
		}
	}
	return result, nil
}

// FindActiveByDate находит активные планы на указанную дату
func (r *CropPlanRepo) FindActiveByDate(ctx context.Context, date time.Time) ([]*cropplan.CropPlan, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*cropplan.CropPlan
	for _, plan := range r.plans {
		if plan.Status() != cropplan.StatusActive {
			continue
		}
		// Проверяем, что дата в пределах сезона
		if (date.Equal(plan.Season().GetStartDate()) || date.After(plan.Season().GetStartDate())) &&
			(date.Equal(plan.Season().GetEndDate()) || date.Before(plan.Season().GetEndDate())) {
			result = append(result, plan)
		}
	}
	return result, nil
}

// FindBySeason находит планы на грядке за сезон
func (r *CropPlanRepo) FindBySeason(ctx context.Context, bedID string, seasonStart, seasonEnd time.Time) ([]*cropplan.CropPlan, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*cropplan.CropPlan
	for _, plan := range r.plans {
		if plan.Area().GetId() != bedID {
			continue
		}
		// Проверяем пересечение сезонов
		if plan.Season().GetStartDate().Before(seasonEnd) && seasonStart.Before(plan.Season().GetEndDate()) {
			result = append(result, plan)
		}
	}
	return result, nil
}

// GetStatistics возвращает статистику по планам
func (r *CropPlanRepo) GetStatistics(ctx context.Context, filter cropplan.StatisticsFilter) (*cropplan.Statistics, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	stats := &cropplan.Statistics{}

	for _, plan := range r.plans {
		// Фильтр по грядке
		if filter.BedID != "" && plan.Area().GetId() != filter.BedID {
			continue
		}
		// Фильтр по сорту
		if filter.VarietyID != "" && plan.Variety().GetId() != filter.VarietyID {
			continue
		}
		// Фильтр по дате
		if !filter.DateFrom.IsZero() && plan.CreatedAt().Before(filter.DateFrom) {
			continue
		}
		if !filter.DateTo.IsZero() && plan.CreatedAt().After(filter.DateTo) {
			continue
		}

		stats.TotalPlans++

		switch plan.Status() {
		case cropplan.StatusActive:
			stats.ActivePlans++
		case cropplan.StatusCompleted:
			stats.CompletedPlans++
			stats.TotalHarvestKg += plan.HarvestKg()
		}
	}

	return stats, nil
}

// Clear очищает репозиторий (для тестов)
func (r *CropPlanRepo) Clear() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.plans = make(map[string]*cropplan.CropPlan)
}

// Count возвращает количество планов (для тестов)
func (r *CropPlanRepo) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.plans)
}
