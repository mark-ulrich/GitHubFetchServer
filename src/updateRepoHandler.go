package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func updateRepoHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		processBadRequest(w)
		return
	}

	if currentUser == nil {
		processBadRequest(w)
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
	log.Println(len(*currentUser.Repos))
	for _, repo := range *currentUser.Repos {
		if repo.Name == repoName {
			currentRepo = &repo
			return
		}
	}

	// Not found
	w.WriteHeader(http.StatusNotFound)

}
