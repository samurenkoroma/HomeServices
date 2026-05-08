package cropplan

import (
	"errors"
	"samurenkoroma/services/internal/core/domain/aggregate"
	"time"
)

type Status string

const (
	StatusDraft    Status = "draft"
	StatusPlanned  Status = "planned"
	StatusActive   Status = "active"
	StatusComplete Status = "complete"
)

type CultivationSnapshot struct {
	ID      string            `json:"id"`
	Name    string            `json:"name"`
	Version int               `json:"version"`
	Steps   []CultivationStep `json:"steps"`
}

type CultivationStep struct {
	ID      uint8       `json:"id"`
	Type    string      `json:"type"`
	Trigger StepTrigger `json:"trigger"`
}

type StepTrigger struct {
	Type  string         `json:"type"` // date | bbch
	Value map[string]any `json:"value"`
}

type Plan struct {
	aggregate.BaseAggregate
	ID           string
	Organization string

	CropKey   string
	VarietyID *string

	AreaID   string
	SeasonID string

	StartDate time.Time
	Status    Status

	CultivationPlanID      string
	CultivationPlanVersion int
	Snapshot               CultivationSnapshot

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Plan) Activate() error {
	if p.Status != StatusPlanned {
		return errors.New("cannot activate")
	}

	if len(p.Snapshot.Steps) == 0 {
		return errors.New("no cultivation steps")
	}

	p.Status = StatusActive
	p.AddEvent(NewCropPlanActivated(p.ID))

	return nil
}
