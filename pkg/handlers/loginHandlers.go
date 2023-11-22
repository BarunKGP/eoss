package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func getRepoLanguages(ghData string) []string {
	m := make(map[string]any)
	err := json.Unmarshal([]byte(ghData), &m)
	if err != nil {
		log.Fatalf("Unable to unmarshal JSON: %s", err)
	}

	reposurl, ok := m["repos_url"]
	if !ok {
		log.Fatal("Unable to fetch repos_url from response")
	}

	reposUrl := reposurl.(string)
	languages := getLanguages(string(reposUrl))
	var languagesString []string
	for _, language := range languages {
		languagesString = append(languagesString, language.toString())
	}

	return languagesString
}

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

	languages := getRepoLanguages(prettyJSON.String())
	fmt.Fprint(w, strings.Join(languages, "\n"))
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
