package api

import (
	"time"
)

// Patronage type doesn't match the api name (Patroned/Patroning) because it's the same thing and I don't want two copies of the same type
type Patronage struct {
	CreatedAt   time.Time `json:"created_at"`
	GodfatherID int       `json:"godfather_id"` // Is it actually how it's called or a funny translation?
	ID          int       `json:"id"`
	Ongoing     bool      `json:"ongoing"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      int       `json:"user_id"`
}
