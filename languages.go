package api

import (
	"time"
)

type LanguagesUser struct {
	CreatedAt  time.Time `json:"created_at"`
	ID         int       `json:"id"`
	LanguageID int       `json:"language_id"`
	Position   int       `json:"position"`
	UserID     int       `json:"user_id"`
}
