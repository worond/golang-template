package config

import (
	"golang.org/x/oauth2"
)

var OAuth2Config *oauth2.Config

func InitOAuth2Config() *oauth2.Config {
	OAuth2Config = &oauth2.Config{
		ClientID:     GetString("OAUTH2_CLIENT_ID"),
		ClientSecret: GetString("OAUTH2_CLIENT_SECRET"),
		Scopes:       []string{"openid", "email", "profile"},
		RedirectURL:  GetString("OAUTH2_REDIRECT_URI"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  GetString("OAUTH2_ISSUER") + "/auth",
			TokenURL: GetString("OAUTH2_ISSUER") + "/token",
		},
	}
	return OAuth2Config
}
