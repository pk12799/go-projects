package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func SetupConfig() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     "46764723697-gp073p4qs0rugir0i6j4akhrbjtnccin.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:3000/google/callback",

		Scopes:   []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: google.Endpoint,
	}
	return conf
}
