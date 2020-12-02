package main

import (
	"log"
	"net/http"
)

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		processBadRequest(w)
		return
	}

	appState.Title = "Users"
	if err = mainTemplate.ExecuteTemplate(w, "listUsers", appState); err != nil {
		log.Fatal(err)
	}
}
