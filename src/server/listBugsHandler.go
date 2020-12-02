package main

import (
	"log"
	"net/http"
)

func listBugsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		processBadRequest(w)
		return
	}

	appState.Title = "Bugs"
	if err = mainTemplate.ExecuteTemplate(w, "listBugs", appState); err != nil {
		log.Fatal(err)
	}
}
