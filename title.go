package api

type Title struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TitlesUser struct {
	ID       int  `json:"id"`
	Selected bool `json:"selected"` // You can only select one title at a time. So tells wether it's selected or not
	TitleID  int  `json:"title_id"`
	UserID   int  `json:"user_id"`
}
