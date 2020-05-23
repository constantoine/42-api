package api

import (
	"time"
)

type ProjectsUser struct {
	CurrentTeamID int       `json:"current_team_id"`
	CursusIds     []int     `json:"cursus_ids"`
	FinalMark     int       `json:"final_mark"`
	ID            int       `json:"id"`
	Marked        bool      `json:"marked"`
	MarkedAt      time.Time `json:"marked_at"`
	Occurrence    int       `json:"occurrence"`
	Project       Project   `json:"project"`
	RetriableAt   time.Time `json:"retriable_at"`
	Status        string    `json:"status"`
	Validated     bool      `json:"validated?"`
}

type Project struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	ParentID interface{} `json:"parent_id"`
	Slug     string      `json:"slug"`
}
