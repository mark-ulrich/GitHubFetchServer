package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Repository struct {
	Name     string
	ID       int
	Owner    *User
	HTMLURL  string `json:"html_url"`
	Issues   *[]Issue
	Bugs     []Issue
	BugCount int
}

type Issue struct {
	Title  string
	Labels *[]Label
}

type Label struct {
	Name string
}

func updateRepoHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		processBadRequest(w)
		return
	}

	// Aliases for brevity
	currentUser := appState.CurrentUser

	// If we're still processing an update to the current user, return status
	// 202 Accepted; otherwise, bail if we don't have a current user.
	if isUpdateUserProcessing {
		w.WriteHeader(http.StatusAccepted)
		return
	} else if currentUser == nil {
		// w.WriteHeader(http.StatusNotFound)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get repo name from JSON POST data
	var repoName string
	err := json.NewDecoder(r.Body).Decode(&repoName)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Does user have a repo matching the given name?
	for i, repo := range *currentUser.Repos {
		if repo.Name == repoName {
			appState.CurrentRepo = &(*currentUser.Repos)[i]
			fetchIssues(appState.CurrentRepo)
			fetchBugs(appState.CurrentRepo)
			return
		}
	}

	// Not found
	w.WriteHeader(http.StatusNotFound)
}
