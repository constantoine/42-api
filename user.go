package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type User struct {
	Achievements    []Achievement    `json:"achievements"`
	Campus          []Campus         `json:"campus"`
	CampusUsers     []CampusUser     `json:"campus_users"`
	CorrectionPoint int              `json:"correction_point"`
	CursusUsers     []CursusUser     `json:"cursus_users"`
	Displayname     string           `json:"displayname"`
	Email           string           `json:"email"`
	ExpertisesUsers []ExpertisesUser `json:"expertises_users"`
	FirstName       string           `json:"first_name"`
	Groups          []interface{}    `json:"groups"` // I have no fucking idea yet, the prequel
	ID              int              `json:"id"`
	ImageURL        string           `json:"image_url"`
	LanguagesUsers  []LanguagesUser  `json:"languages_users"`
	LastName        string           `json:"last_name"`
	Location        string           `json:"location"`
	Login           string           `json:"login"`
	Partnerships    []interface{}    `json:"partnerships"` // I have no fucking idea yet
	Patroned        []Patronage      `json:"patroned"`     // List of patronages, not a boolean value as its name may suggest
	Patroning       []Patronage      `json:"patroning"`
	Phone           string           `json:"phone"`
	PoolMonth       string           `json:"pool_month"`
	PoolYear        string           `json:"pool_year"`
	ProjectsUsers   []ProjectsUser   `json:"projects_users"`
	Staff           bool             `json:"staff?"` // The fuck, why would you even use my shitty library
	Titles          []Title          `json:"titles"`
	TitlesUsers     []TitlesUser     `json:"titles_users"`
	URL             string           `json:"url"`
	Wallet          int              `json:"wallet"`
}

func (usr User) IsStud() bool {
	for i := range usr.CursusUsers {
		if usr.CursusUsers[i].ID == 21 {
			return true
		}
	}
	return false
}

func (usr User) GetBH() time.Time {
	for i := range usr.CursusUsers {
		if usr.CursusUsers[i].ID == 21 {
			return usr.CursusUsers[i].BlackholedAt
		}
	}
	return time.Time{}
}

func (api *API) User(userID uint) (User, error) {
	var usr User
	uri := fmt.Sprintf("https://api.intra.42.fr/v2/users/%d", userID)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return User{}, err
	}
	api.lock()
	defer api.unlock()
	resp, err := api.conf.Client(context.Background(), api.tok).Do(req)
	if err != nil {
		return User{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return User{}, err
	}
	if err = json.Unmarshal(body, &usr); err != nil {
		return User{}, err
	}
	return usr, nil
}

func (api *API) Users(query *Search) ([]User, error) {
	var usr []User
	q, err := query.QueryString()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("https://api.intra.42.fr/v2/users%s", q)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	api.lock()
	defer api.unlock()
	resp, err := api.conf.Client(context.Background(), api.tok).Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &usr); err != nil {
		return nil, err
	}
	return usr, nil
}

func (api *API) Me() (User, error) {
	var usr User
	req, err := http.NewRequest("GET", "https://api.intra.42.fr/v2/me", nil)
	if err != nil {
		return User{}, err
	}
	api.lock()
	defer api.unlock()
	resp, err := api.conf.Client(context.Background(), api.tok).Do(req)
	if err != nil {
		return User{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return User{}, err
	}
	if err = json.Unmarshal(body, &usr); err != nil {
		return User{}, err
	}
	return usr, nil
}
