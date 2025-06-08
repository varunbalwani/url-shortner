package models

import (
	"time"
)
type ClickEvent struct {
	ShortCode   string    `json:"short_code"`
	OriginalURL string    `json:"original_url"`
	Timestamp   time.Time `json:"timestamp"`
}