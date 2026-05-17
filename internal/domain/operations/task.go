package operations

import "time"

// Task - Operational task for agronomist
type Task struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	ProductionUnitID string   `json:"production_unit_id"`
	DueDate         time.Time `json:"due_date"`
	Status          string    `json:"status"`
	Priority        int       `json:"priority"`
}
