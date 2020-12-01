package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type GitHubUserSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []User
}

type User struct {
	Login string
	ID    int
	Repos *[]Repository
}

var (
	isUpdateUserProcessing bool
)

func finishedUserProcessing() {
	isUpdateUserProcessing = false
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	if r.Method != "POST" {
		processBadRequest(w)
		return
	}

	isUpdateUserProcessing = true
	defer finishedUserProcessing()
	appState.CurrentUser = nil

	// Search GitHub for username
	var username string
	err = json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		log.Println(err)
	}
	resp, err := http.Get("https://api.github.com/search/users?q=" + username)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		var body []byte
		w.WriteHeader(resp.StatusCode)
		_, err = resp.Body.Read(body)
		if err != nil {
			log.Println(err)
		}
		_, err = w.Write(body)
		if err != nil {
			log.Println(err)
		}
	}

	var result GitHubUserSearchResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for i, user := range result.Items {
		if strings.EqualFold(username, user.Login) {
			// Update current user server-wide
			appState.CurrentUser = &result.Items[i]
			break
		}
	}

	if appState.CurrentUser == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	appState.CurrentUser.Repos, err = fetchRepos(appState.CurrentUser.Login)
	if err != nil {
		log.Println(err)
	}

}
