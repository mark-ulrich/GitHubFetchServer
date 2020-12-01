package main

import (
	"html/template"
	"log"
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

	pageTemplate, err := template.ParseFiles("../html/layout.gohtml", "../html/updateRepo.gohtml")
	if err != nil {
		log.Println(err)
	}
	err = pageTemplate.Execute(w, appState)
	if err != nil {
		log.Println(err)
	}
}
