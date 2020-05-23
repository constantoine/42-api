package api

import (
	"context"
	"fmt"
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

// API encapsulates Auth as to not pollute Token
type API struct {
	Auth Auth
}

type Token struct {
	Token            string
	ResourceOwnerID  int       `json:"resource_owner_id"`
	Scopes           []string  `json:"scopes"`
	ExpiresInSeconds int       `json:"expires_in_seconds"`
	CreatedAt        time.Time `json:"created_at"`
}

type Auth struct {
	conf  *oauth2.Config
	Token string
}

func NewAPI(id string, secret string, redirect string, scope []string) API {
	var api API = API{
		Auth: Auth{
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
		},
	}
	return api
}

func (api *API) ConnectWithCode(code string) error {
	tok, err := api.Auth.conf.Exchange(context.Background(), code)
	if err != nil {
		return err
	}
	api.Auth.Token = tok.AccessToken
	return nil
}

func main() {
	api := NewAPI(
		"XXXXXXXXXXXXX",
		"XXXXXXXXXXXXX",
		"http://localhost:8080/redirect",
		[]string{
			"public",
		},
	)
	var r *http.Request
	if err := api.ConnectWithCode(r.FormValue("code")); err != nil {
		panic(err)
	}
	fmt.Println(api.Auth.Token)
}
