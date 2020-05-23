package api

import (
	"time"
)

type CursusUser struct {
	BeginAt      time.Time       `json:"begin_at"`
	BlackholedAt time.Time       `json:"blackholed_at"` // Can sometimes be null, maybe because cursus was ended? Anyway, it might get set to the Unix Epoch in that case
	Cursus       Cursus          `json:"cursus"`
	CursusID     int             `json:"cursus_id"`
	EndAt        time.Time       `json:"end_at"`
	Grade        string          `json:"grade"` // Can sometimes be null, maybe because cursus was ended? Grade is to be understood as in a molitay grade, not a numerical one.
	HasCoalition bool            `json:"has_coalition"`
	ID           int             `json:"id"`
	Level        float64         `json:"level"`
	Skills       []Skill         `json:"skills"`
	User         CursusUserInner `json:"user"`
}

type Skill struct {
	ID    int     `json:"id"`
	Level float64 `json:"level"`
	Name  string  `json:"name"`
}

// CursusUserInner type doesn't match the api name (User) because its type would collide with other user types
type CursusUserInner struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
	URL   string `json:"url"`
}

type Cursus struct {
	CreatedAt time.Time `json:"created_at"`
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
}
