package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type reposResponse struct {
	Id          int64                  `json:"id"`
	Name        string                 `json:"name"`
	LanguageUrl string                 `json:"languages_url"`
	Fork        bool                   `json:"fork"`
	Rest        map[string]interface{} `json:"~"`
}

type LanguageType struct {
	Repo     string
	Language map[string]int
}

func (lang *LanguageType) toString() string {
	b := new(bytes.Buffer)
	fmt.Fprintf(b, "Repo name: %s, Languages:\n", lang.Repo)
	for key, value := range lang.Language {
		fmt.Fprintf(b, "\t%s: %d \n", key, value)
	}
	return b.String()
}

func getGithubClientId() string {
	githubClientId, exists := os.LookupEnv("CLIENT_ID")
	if !exists {
		log.Fatal("Github Client ID not defined in .env file")
	}
	return githubClientId
}

func getGithubClientSecret() string {
	getGithubClientSecret, exists := os.LookupEnv("CLIENT_SECRET")
	if !exists {
		log.Fatal("Github Client Secret not defined in .env file")
	}
	return getGithubClientSecret
}

func getPort() string {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		log.Fatal("Port not defined in .env file")
	}
	return port
}

func getGithubAccessToken(code string) string {
	clientId := getGithubClientId()
	clientSecret := getGithubClientSecret()
	requestBodyMap := map[string]string{
		"client_id":     clientId,
		"client_secret": clientSecret,
		"code":          code,
	}
	requestJSON, _ := json.Marshal(requestBodyMap)

	// POST request to set URL
	req, reqerr := http.NewRequest("POST",
		"https://github.com/login/oauth/access_token",
		bytes.NewBuffer(requestJSON))

	if reqerr != nil {
		log.Panic("Request creation failed")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Get response
	resp, resperr := http.DefaultClient.Do(req)
	if resperr != nil {
		log.Panic("Request failed")
	}

	respbody, _ := ioutil.ReadAll(resp.Body)
	// log.Printf("respbody: %s\n", respbody)

	type githubAccessTokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}

	// Convert stringified JSON to struct object
	var ghresp githubAccessTokenResponse
	json.Unmarshal(respbody, &ghresp)
	return ghresp.AccessToken
}

func getGithubData(accessToken string) string {
	req, reqerr := http.NewRequest(
		"GET",
		"https://api.github.com/user",
		nil,
	)
	if reqerr != nil {
		log.Panic("API Request failed to send")
	}

	authorizationHeaderValue := fmt.Sprintf("Bearer %s", accessToken)
	req.Header.Set("Authorization", authorizationHeaderValue)

	resp, resperr := http.DefaultClient.Do(req)
	if resperr != nil {
		log.Panic("Request failed")
	}
	defer resp.Body.Close()

	respbody, _ := io.ReadAll(resp.Body)
	return string(respbody)
}

func getLanguagesFromRepoUrl(reposUrl string) []string {
	req, err := http.NewRequest(
		"GET",
		reposUrl,
		nil,
	)
	if err != nil {
		log.Fatal("Malformed repos url")
	}

	resp, resperr := http.DefaultClient.Do(req)
	if resperr != nil {
		log.Panic("Request failed")
	}
	defer resp.Body.Close()

	respbody, _ := io.ReadAll(resp.Body)
	
	var respJson []reposResponse
	json.Unmarshal(respbody, &respJson)

	var languageUrls []string
	for _, repoResponse := range respJson {
		languageUrls = append(languageUrls, repoResponse.LanguageUrl)
	}
	return languageUrls

	// return getLanguages(respJson)
}

func getRepoNameFromUrl(url string) string {
	log.Printf("repo url =  %s", url)
	parts := strings.Split(url, "/")
	if len(parts) < 2 {
		log.Fatal("Malformed language url")
	}
	repoName := parts[len(parts) - 2]
	return repoName
}

func getLanguages(reposUrl string) []LanguageType {
	// repos := getRepos(reposUrl)
	// var reposJson []reposResponse
	// if err := json.Unmarshal(repos, &reposJson); err != nil {
	// 	log.Fatal("Malformed repos reponse")
	// }

	languageUrls := getLanguagesFromRepoUrl(reposUrl)
	
	var languages []LanguageType
	// for _, repo := range reposJson {
	for _, langUrl := range languageUrls {
		resp, err := http.Get(langUrl)
		if err != nil {
			log.Fatalf("Unable to fetch language url: %s\n", err)
		}

		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		languageMap := make(map[string]int)
		json.Unmarshal(body, &languageMap)

		language := LanguageType{Repo: getRepoNameFromUrl(langUrl), Language: languageMap}
		languages = append(languages, language)
	}

	return languages
}
