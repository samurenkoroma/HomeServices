package dto

import (
	"samurenkoroma/services/internal/modules/growing/domain/cropplan/task"
	"time"
)

// TaskDTO задание для агронома
type TaskDTO struct {
	ID     string  `json:"id"`
	PlanID *string `json:"plan_id,omitempty"`
	BedID  string  `json:"bed_id"`

	Type          string `json:"type"`
	TypeText      string `json:"type_text"`
	Status        string `json:"status"`
	StatusText    string `json:"status_text"`
	Priority      string `json:"priority"`
	PriorityText  string `json:"priority_text"`
	PriorityColor string `json:"priority_color"`

	Title        string `json:"title"`
	Description  string `json:"description"`
	Instructions string `json:"instructions"`

	ScheduledDate time.Time `json:"scheduled_date"`
	DueDate       time.Time `json:"due_date"`
	IsOverdue     bool      `json:"is_overdue"`

	EstimatedDuration int `json:"estimated_duration"`
	ActualDuration    int `json:"actual_duration,omitempty"`

	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	Photos   []PhotoDTO   `json:"photos,omitempty"`
	Comments []CommentDTO `json:"comments,omitempty"`

	CompletedAt *time.Time `json:"completed_at,omitempty"`
	CompletedBy string     `json:"completed_by,omitempty"`
}

// PhotoDTO фото
type PhotoDTO struct {
	ID        string    `json:"id"`
	URL       string    `json:"url"`
	Thumbnail string    `json:"thumbnail"`
	TakenAt   time.Time `json:"taken_at"`
	Notes     string    `json:"notes"`
}

// CommentDTO комментарий
type CommentDTO struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	UserName  string    `json:"user_name"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

// ToTaskDTO конвертирует доменный Task в DTO
func ToTaskDTO(t *task.Task) TaskDTO {
	return TaskDTO{
		ID:                t.ID,
		PlanID:            t.PlanID,
		BedID:             t.BedID,
		Type:              string(t.Type),
		TypeText:          t.Type.String(),
		Status:            string(t.Status),
		StatusText:        t.Status.String(),
		Priority:          string(t.Priority),
		PriorityText:      t.Priority.String(),
		PriorityColor:     t.Priority.Color(),
		Title:             t.Title,
		Description:       t.Description,
		Instructions:      t.Instructions,
		ScheduledDate:     t.ScheduledDate,
		DueDate:           t.DueDate,
		IsOverdue:         t.IsOverdue(),
		EstimatedDuration: t.EstimatedDuration,
		ActualDuration:    t.ActualDuration,
		Latitude:          t.Latitude,
		Longitude:         t.Longitude,
		CompletedAt:       t.CompletedAt,
		CompletedBy:       t.CompletedBy,
	}
}
