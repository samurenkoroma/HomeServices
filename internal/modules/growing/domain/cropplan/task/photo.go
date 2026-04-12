package task

import (
	"time"

	"github.com/google/uuid"
)

// Photo фотография, прикрепленная к заданию
type Photo struct {
	ID        string    `json:"id"`
	URL       string    `json:"url"`       // оригинал
	Thumbnail string    `json:"thumbnail"` // миниатюра
	TakenAt   time.Time `json:"taken_at"`
	Notes     string    `json:"notes"` // подпись к фото
}

// NewPhoto создает новую фотографию
func NewPhoto(url, thumbnail, notes string) Photo {
	return Photo{
		ID:        uuid.New().String(),
		URL:       url,
		Thumbnail: thumbnail,
		TakenAt:   time.Now(),
		Notes:     notes,
	}
}
