package api

type NbrSuccess interface {
}

type Achievement struct {
	Description  string `json:"description"`
	ID           int    `json:"id"`
	Image        string `json:"image"`
	Kind         string `json:"kind"`
	Name         string `json:"name"`
	NbrOfSuccess int    `json:"nbr_of_success"` // If the API returns NULL for this value (often means it4's not relevant due to achievement type) this will be set to 0
	Tier         string `json:"tier"`
	UsersURL     string `json:"users_url"`
	Visible      bool   `json:"visible"`
}
