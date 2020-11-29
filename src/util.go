package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Repository struct {
	Name string
	ID   int
}

// Return a 401 Bad Request response
func processBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("<h1>400 Bad Request</h1>"))
}

// Return a list of repositories for a user
func fetchRepos(username string) (*[]Repository, error) {
	var repos []Repository

	url := fmt.Sprintf("https://api.github.com/users/%s/repos?per_page=100000", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error fetching repos: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&repos)
	return &repos, err
}
