package season

import (
	"samurenkoroma/services/internal/core/domain/event"
	"time"
)

// SeasonCreated — событие создания сезона
type SeasonCreated struct {
	event.BaseEvent
	SeasonID  string    `json:"season_id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

func (e SeasonCreated) EventName() string {
	return "growing.season.created"
}

// SeasonActivated — событие активации сезона
type SeasonActivated struct {
	event.BaseEvent
	SeasonID string `json:"season_id"`
	Name     string `json:"name"`
}

func (e SeasonActivated) EventName() string {
	return "growing.season.activated"
}

// SeasonCompleted — событие завершения сезона
type SeasonCompleted struct {
	event.BaseEvent
	SeasonID string `json:"season_id"`
	Name     string `json:"name"`
}

func (e SeasonCompleted) EventName() string {
	return "growing.season.completed"
}

// SeasonArchived — событие архивации сезона
type SeasonArchived struct {
	event.BaseEvent
	SeasonID string `json:"season_id"`
}

func (e SeasonArchived) EventName() string {
	return "growing.season.archived"
}
