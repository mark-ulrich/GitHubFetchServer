package main

import (
	"net/http"
)

// The update-repo handler. A GET request displays a form to enter a GitHub
// username and repository name. A POST request updates the server's
// currentRepoInfo structure.
func updateRepoInfoHandler(w http.ResponseWriter, r *http.Request) {

	if !stringInSlice([]string{"GET", "POST"}, r.Method) {
		processBadRequest(w)
		return
	}

	appState.Request = r
	appState.Title = "Select Repository"

	if err := mainTemplate.ExecuteTemplate(w, "updateRepo", appState); err != nil {
		panic(err)
	}
}
