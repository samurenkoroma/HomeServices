package cultivation

import (
	"samurenkoroma/services/internal/core/domain/types"
	"time"
)

type TriggerType string

const (
	TriggerDayOffset   TriggerType = "day_offset"
	TriggerDayInterval TriggerType = "day_interval"
	TriggerBBCH        TriggerType = "bbch"
)

type Trigger struct {
	Type  TriggerType
	Value map[string]any
}

type Step struct {
	ID      uint8
	Type    string
	Trigger Trigger
}

func NewStep(ID uint8, Type string, trigger Trigger) *Step {
	return &Step{ID: ID, Type: Type, Trigger: trigger}
}

type CultivationPlan struct {
	ID        string
	Name      string
	CropKey   string
	VarietyID *string
	Version   int
	Steps     []Step
	CreatedAt time.Time
}

func NewCultivationPlan(name string, cropKey string, varietyID *string, version int) *CultivationPlan {
	return &CultivationPlan{
		ID:        types.NewUUID(),
		Name:      name,
		CropKey:   cropKey,
		VarietyID: varietyID,
		Version:   version,
		CreatedAt: time.Now(),
	}
}

func (c *CultivationPlan) AddStep(step *Step) {
	c.Steps = append(c.Steps, *step)
}
