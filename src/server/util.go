package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Writes a 400 Bad Request response to an http.ResponseWriter
func processBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("<h2>400 Bad Request</h2>"))
}

// Search for string s in slice of strings.
func stringInSlice(slice []string, s string) bool {
	for _, str := range slice {
		if str == s {
			return true
		}
	}
	return false
}

// Fetch a list of repositories for a given GitHub username
func fetchRepos(username string) (repos *[]Repository, err error) {
	resp, err := http.Get("https://api.github.com/users/mark-ulrich/repos?per_page=100000")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to fetch repos: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&repos)
	return repos, err
}
