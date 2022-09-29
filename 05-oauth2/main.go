package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// key is github ID, value is user ID
var githubConnections map[string]string

// json layout
type githubResponse struct {
	Data struct {
		Viewer struct {
			ID string `json:"id"`
		} `json:"viewer"`
	} `json:"data"`
}

var githubOauthConfig = &oauth2.Config{
	ClientID:     "a4482d5ac404bed1dd68",
	ClientSecret: "71c2510e4c9f3f123fe6fd3aaa19aafb99fac97b",
	Endpoint:     github.Endpoint,
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/oauth/github", startGithubOauth)
	http.HandleFunc("/oauth2/receive", completeGithubOauth)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Document</title>
  </head>
  <body>
    <form action="/oauth/github" method="post">
      <input type="submit" value="Login with Github" />
    </form>
  </body>
</html>`)
}

func startGithubOauth(w http.ResponseWriter, r *http.Request) {
	redirectURL := githubOauthConfig.AuthCodeURL("0000")
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

func completeGithubOauth(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	state := r.FormValue("state")

	if state != "0000" {
		http.Error(w, "State is incorrect", http.StatusBadRequest)
		return
	}

	token, err := githubOauthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Couldn't login", http.StatusInternalServerError)
		return
	}

	ts := githubOauthConfig.TokenSource(r.Context(), token)
	client := oauth2.NewClient(r.Context(), ts)

	requestBody := strings.NewReader(`{"query": "query {viewer {id}}"}`)
	resp, err := client.Post("https://api.github.com/graphql", "application/json", requestBody)
	if err != nil {
		http.Error(w, "Couldn't get user", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Couldn't read github information", http.StatusInternalServerError)
		return
	}

	log.Println(string(bs))

	var gr githubResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		http.Error(w, "Github invalid response", http.StatusInternalServerError)
		return
	}

	githubID := gr.Data.Viewer.ID
	userID, ok := githubConnections[githubID]
	if !ok {
		// New User - create account
	}

	// Login to account userID using JWT
}
