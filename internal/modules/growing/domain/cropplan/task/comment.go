package task

import (
	"time"

	"github.com/google/uuid"
)

// Comment комментарий к заданию
type Comment struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	UserName  string    `json:"user_name"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

// NewComment создает новый комментарий
func NewComment(userID, userName, text string) Comment {
	return Comment{
		ID:        uuid.New().String(),
		UserID:    userID,
		UserName:  userName,
		Text:      text,
		CreatedAt: time.Now(),
	}
}
