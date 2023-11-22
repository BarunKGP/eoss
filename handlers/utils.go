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
)

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

	respbody, _ := io.ReadAll(resp.Body)
	return string(respbody)
}
