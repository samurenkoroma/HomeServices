package worktask

import "time"

type Task struct {
	ID            string
	CropPlanID    string
	Type          string
	ScheduledDate time.Time
	Status        string
	Params        map[string]any
}
