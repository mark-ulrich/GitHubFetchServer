package main

import (
	"html/template"
	"log"
	"net/http"
)

// Root handler for the server.
func rootHandler(w http.ResponseWriter, r *http.Request) {

	// We must have a valid repo selected
	if appState.CurrentRepo == nil {
		http.Redirect(w, r, UpdateRepoHandlerPath, http.StatusTemporaryRedirect)
	}

	pageTemplate, err := template.ParseFiles("../html/layout.gohtml")
	if err != nil {
		log.Println(err)
	}
	if err = pageTemplate.Execute(w, appState); err != nil {
		log.Println(err)
	}

}
