package main

import (
	"log"
	"net/http"
)

func listMilestonesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		processBadRequest(w)
		return
	}

	appState.Title = "Milestones"
	if err = mainTemplate.ExecuteTemplate(w, "listMilestones", appState); err != nil {
		log.Fatal(err)
	}
}
