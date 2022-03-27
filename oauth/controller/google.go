package controller

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"oauth/config"
)

func Home(w http.ResponseWriter, r *http.Request) {
	var html = `<html><body><a href="/google/login">Google Login </a></body></html>`
	fmt.Fprint(w, html)
}

func GoogleLogin(res http.ResponseWriter, req *http.Request) {

	googleConfig := config.SetupConfig()
	url := googleConfig.AuthCodeURL("randomstate")
	http.Redirect(res, req, url, http.StatusTemporaryRedirect)
}

func GoogleCallback(res http.ResponseWriter, req *http.Request) {
	state := req.URL.Query()["state"][0]
	if state != "randomstate" {
		fmt.Fprintln(res, "state dont match")
		return
	}
	// code := req.URL.Query()["code"][0]
	// log.Println(state)
	googleConfig := config.SetupConfig()
	// log.Println(googleConfig)
	token, err := googleConfig.Exchange(context.Background(), req.FormValue("code"))
	if err != nil {
		fmt.Fprintln(res, "code token exchange Failed")
	}
	// log.Println(token.AccessToken)
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	// log.Println(resp)

	if err != nil {
		fmt.Fprintln(res, "User data fetch failed")
	}
	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(res, "json parsing failed")
	}

	fmt.Fprintln(res, string(userData))
}
