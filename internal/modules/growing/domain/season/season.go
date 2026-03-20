package season

import (
	"samurenkoroma/services/internal/core/domain/aggregate"
	"time"
)

type SeasonStatus string

const (
	SeasonStatusPlanning  SeasonStatus = "planning"
	SeasonStatusActive    SeasonStatus = "active"
	SeasonStatusCompleted SeasonStatus = "completed"
	SeasonStatusArchived  SeasonStatus = "archived"
)

// Season - агрономический сезон
type Season struct {
	aggregate.BaseAggregate

	ID          string
	Name        string
	StartDate   time.Time
	EndDate     time.Time
	Status      SeasonStatus
	Description string

	// План на сезон
	Plan *SeasonPlan

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewSeason(name string, startDate, endDate time.Time) *Season {
	return &Season{
		ID:        generateID(),
		Name:      name,
		StartDate: startDate,
		EndDate:   endDate,
		Status:    SeasonStatusPlanning,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (s *Season) Activate() error {
	if s.Status != SeasonStatusPlanning {
		return ErrInvalidStatusTransition
	}

	now := time.Now()
	if now.Before(s.StartDate) {
		return ErrSeasonNotStarted
	}

	s.Status = SeasonStatusActive
	s.UpdatedAt = now

	s.AddEvent(SeasonActivated{
		SeasonID: s.ID,
		Name:     s.Name,
	})

	return nil
}

func (s *Season) Complete() error {
	if s.Status != SeasonStatusActive {
		return ErrInvalidStatusTransition
	}

	s.Status = SeasonStatusCompleted
	s.UpdatedAt = time.Now()

	s.AddEvent(SeasonCompleted{
		SeasonID: s.ID,
	})

	return nil
}
