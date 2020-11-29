package main

import (
	"encoding/json"
	"log"
	"net/http"
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

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	if r.Method != "POST" {
		processBadRequest(w)
		return
	}

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
		resp.Body.Read(body)
		w.Write(body)
	}

	var result GitHubUserSearchResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if result.TotalCount != 1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Update current user server-wide
	currentUser = &result.Items[0]
	currentUser.Repos, err = fetchRepos(currentUser.Login)
	if err != nil {
		log.Println(err)
	}

}
