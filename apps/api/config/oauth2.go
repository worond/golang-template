package config

import (
	"os"

	"golang.org/x/oauth2"
)

var OAuth2Config *oauth2.Config

func InitOAuth2Config() *oauth2.Config {
	OAuth2Config = &oauth2.Config{
		ClientID:     os.Getenv("OAUTH2_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH2_CLIENT_SECRET"),
		Scopes:       []string{"openid", "email", "profile"},
		RedirectURL:  os.Getenv("OAUTH2_REDIRECT_URI"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  os.Getenv("OAUTH2_ISSUER") + "/auth",
			TokenURL: os.Getenv("OAUTH2_ISSUER") + "/token",
		},
	}
	return OAuth2Config
}
