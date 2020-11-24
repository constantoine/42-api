package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func (api *API) UserCampusUsers(userID uint, query *Search) ([]CampusUser, error) {
	var usr []CampusUser
	q, err := query.QueryString()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("https://api.intra.42.fr/v2/users/%d/campus_users%s", userID, q)
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

func (api *API) CampusCampusUsers(campusID uint, query *Search) ([]CampusUser, error) {
	var usr []CampusUser
	q, err := query.QueryString()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("https://api.intra.42.fr/v2/campus/%d/users%s", campusID, q)
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

func (api *API) CampusUsers(query *Search) ([]CampusUser, error) {
	var usr []CampusUser
	q, err := query.QueryString()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("https://api.intra.42.fr/v2/campus_users%s", q)
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

func (api *API) CampusUser(campusUserID uint) (CampusUser, error) {
	var usr CampusUser
	uri := fmt.Sprintf("https://api.intra.42.fr/v2/campus_users/%d", campusUserID)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return CampusUser{}, err
	}
	api.lock()
	defer api.unlock()
	resp, err := api.conf.Client(context.Background(), api.tok).Do(req)
	if err != nil {
		return CampusUser{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return CampusUser{}, err
	}
	if err = json.Unmarshal(body, &usr); err != nil {
		return CampusUser{}, err
	}
	return usr, nil
}
