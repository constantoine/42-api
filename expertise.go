package api

import (
	"time"
)

type ExpertisesUser struct {
	ContactMe   bool      `json:"contact_me"`
	CreatedAt   time.Time `json:"created_at"`
	ExpertiseID int       `json:"expertise_id"`
	ID          int       `json:"id"`
	Interested  bool      `json:"interested"`
	UserID      int       `json:"user_id"`
	Value       int       `json:"value"`
}
