package api

import (
	"time"
)

type Campus struct {
	Address     string   `json:"address"`
	City        string   `json:"city"`
	Country     string   `json:"country"`
	Facebook    string   `json:"facebook"`
	ID          int      `json:"id"`
	Language    Language `json:"language"`
	Name        string   `json:"name"`
	TimeZone    string   `json:"time_zone"`
	Twitter     string   `json:"twitter"`
	UsersCount  int      `json:"users_count"`
	VogsphereID int      `json:"vogsphere_id"`
	Website     string   `json:"website"`
	Zip         string   `json:"zip"`
}

type Language struct {
	CreatedAt  time.Time `json:"created_at"`
	ID         int       `json:"id"`
	Identifier string    `json:"identifier"`
	Name       string    `json:"name"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CampusUser struct {
	CampusID  int  `json:"campus_id"`
	ID        int  `json:"id"`
	IsPrimary bool `json:"is_primary"`
	UserID    int  `json:"user_id"`
}
