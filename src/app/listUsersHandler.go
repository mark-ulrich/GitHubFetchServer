package main

import (
	"html/template"
	"log"
	"net/http"
)

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		processBadRequest(w)
		return
	}

	appState.Title = "Users"
	pageTemplate, err := template.ParseFiles("../html/layout.gohtml", "../html/listUsers.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	if err = pageTemplate.Execute(w, appState); err != nil {
		log.Fatal(err)
	}
}
