package season

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"samurenkoroma/services/internal/core/domain/types"
	"time"
)

type SeasonStatus string
type SeasonID string

const (
	SeasonStatusPlanning  SeasonStatus = "planning"
	SeasonStatusActive    SeasonStatus = "active"
	SeasonStatusCompleted SeasonStatus = "completed"
	SeasonStatusArchived  SeasonStatus = "archived"
)

// Season - агрономический сезон
type Season struct {
	aggregate.Entity[SeasonID]
	name        string
	startDate   time.Time
	endDate     time.Time
	status      SeasonStatus
	description string
	createdBy   string

	// План на сезон
	plan *SeasonPlan
}

func NewSeason(
	name string,
	startDate, endDate time.Time,
	status SeasonStatus,
	createdBy string,
	description string,
) (*Season, error) {
	if name == "" {
		return nil, ErrInvalidName
	}
	if startDate.After(endDate) {
		return nil, ErrInvalidPeriod
	}
	if createdBy == "" {
		return nil, ErrInvalidCreatedBy
	}

	s := &Season{
		Entity:      aggregate.NewEntity(SeasonID(types.NewUUID())),
		name:        name,
		startDate:   startDate,
		endDate:     endDate,
		status:      status,
		createdBy:   createdBy,
		description: description,
	}

	//s.AddEvent(SeasonCreated{
	//	SeasonID:  string(s.Id),
	//	Name:      s.name,
	//	StartDate: s.startDate,
	//	EndDate:   s.endDate,
	//})

	return s, nil
}

// Activate активирует сезон
func (s *Season) Activate() error {
	if s.status != SeasonStatusPlanning {
		return ErrInvalidStatusTransition
	}

	now := time.Now()
	if now.Before(s.startDate) {
		return ErrSeasonNotStarted
	}

	s.status = SeasonStatusActive
	s.Update()

	s.AddEvent(SeasonActivated{
		SeasonID: string(s.Id),
		Name:     s.name,
	})

	return nil
}

// Complete завершает сезон
func (s *Season) Complete() error {
	if s.status != SeasonStatusActive {
		return ErrInvalidStatusTransition
	}

	s.status = SeasonStatusCompleted
	s.Update()

	s.AddEvent(SeasonCompleted{
		SeasonID: string(s.Id),
		Name:     s.name,
	})

	return nil
}

// Archive архивирует сезон
func (s *Season) Archive() error {
	if s.status == SeasonStatusArchived {
		return ErrAlreadyArchived
	}

	s.status = SeasonStatusArchived
	s.Update()

	s.AddEvent(SeasonArchived{
		SeasonID: string(s.Id),
	})

	return nil
}

// SetPlan устанавливает план на сезон
func (s *Season) SetPlan(plan *SeasonPlan) {
	s.plan = plan
	s.Update()
}

// GetPlan возвращает план на сезон
func (s *Season) GetPlan() *SeasonPlan {
	return s.plan
}

// IsActive проверяет, активен ли сезон в указанную дату
func (s *Season) IsActiveAt(date time.Time) bool {
	return !date.Before(s.startDate) && !date.After(s.endDate)
}

// Getters
func (s *Season) GetId() string           { return string(s.Id) }
func (s *Season) IsFinished() bool        { return time.Now().After(s.endDate) }
func (s *Season) GetName() string         { return s.name }
func (s *Season) GetStartDate() time.Time { return s.startDate }
func (s *Season) GetEndDate() time.Time   { return s.endDate }
func (s *Season) GetDescription() string  { return s.description }
func (s *Season) GetStatus() SeasonStatus { return s.status }
func (s *Season) GetCreatedBy() string    { return s.createdBy }
func (s *Season) GetCreatedAt() time.Time { return s.CreatedAt }
func (s *Season) GetUpdatedAt() time.Time { return s.UpdatedAt }

// Duration возвращает длительность сезона в днях
func (s *Season) Duration() int {
	return int(s.endDate.Sub(s.startDate).Hours() / 24)
}

func (s *Season) Delete() {
	now := time.Now()
	s.DeletedAt = &now
}

// Rehydrate восстанавливает тип культуры из БД
func Rehydrate(id SeasonID, status SeasonStatus, createdBy, name, description string, startDate, endDate, createdAt, updatedAt time.Time) *Season {
	return &Season{
		Entity: aggregate.Entity[SeasonID]{
			Id:        id,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
		name:        name,
		startDate:   startDate,
		endDate:     endDate,
		status:      status,
		description: description,
		createdBy:   createdBy,
	}
}
