package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func LoggedinHandler(w http.ResponseWriter, r *http.Request, githubData string) {
	if githubData == "" {
		// Unauthorized request
		fmt.Fprintf(w, "Unauthorized!")
		return
	}
	w.Header().Set("Content-Type", "application/json")

	var prettyJSON bytes.Buffer
	parseErr := json.Indent(&prettyJSON, []byte(githubData), "", "\t")
	if parseErr != nil {
		log.Panic("JSON parse error")
	}

	fmt.Fprint(w, prettyJSON.String())
}

func GithubLoginHandler(w http.ResponseWriter, r *http.Request) {
	githubClientId := getGithubClientId()
	port := getPort()
	redirectURL := fmt.Sprintf(
		"https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s",
		githubClientId,
		fmt.Sprintf("http://localhost:%s/login/github/callback", port))


	http.Redirect(w, r, redirectURL, http.StatusMovedPermanently) // TODO: check status code
}

func GithubCallbackHandler(w http.ResponseWriter, r *http.Request) {
	
	code := r.URL.Query().Get("code")
	githubAccessToken := getGithubAccessToken(code)
	githubData := getGithubData(githubAccessToken)
	LoggedinHandler(w, r, githubData)
}