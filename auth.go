package api

import (
	"context"
	"time"

	"golang.org/x/oauth2"
)

const (
	AUTH_URL  = "https://api.intra.42.fr/oauth/authorize"
	TOKEN_URL = "https://api.intra.42.fr/oauth/token"
)

type Application struct {
	UID string `json:"uid"`
}

// API is the main object
type API struct {
	conf *oauth2.Config
	tok  *oauth2.Token
	Sync
}

type Token struct {
	Token            string
	ResourceOwnerID  int       `json:"resource_owner_id"`
	Scopes           []string  `json:"scopes"`
	ExpiresInSeconds int       `json:"expires_in_seconds"`
	CreatedAt        time.Time `json:"created_at"`
}

func NewAPI(id string, secret string, redirect string, scope []string) API {
	var api API = API{
		conf: &oauth2.Config{
			ClientID:     id,
			ClientSecret: secret,
			Scopes:       scope,
			RedirectURL:  redirect,
			Endpoint: oauth2.Endpoint{
				AuthURL:  AUTH_URL,
				TokenURL: TOKEN_URL,
			},
		},
	}
	return api
}

func (api *API) ConnectWithCode(code string, state *string) error {
	var options []oauth2.AuthCodeOption
	if state != nil {
		options = append(options, oauth2.SetAuthURLParam("state", *state))
	}
	tok, err := api.conf.Exchange(context.Background(), code, options...)
	if err != nil {
		return err
	}
	api.tok = tok
	return nil
}

func (api *API) GenerateURI(state string) string {
	return api.conf.AuthCodeURL(state, oauth2.AccessTypeOffline)
}
