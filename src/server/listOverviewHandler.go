package main

import (
	"fmt"
	"log"
	"net/http"
)

func listOverviewHandler(w http.ResponseWriter, r *http.Request) {

	// We must have a valid repo selected
	if appState.CurrentRepo == nil {
		http.Redirect(w, r, UpdateRepoInfoRoutePath, http.StatusTemporaryRedirect)
		return
	}

	if r.Method != "GET" {
		processBadRequest(w)
		return
	}

	appState.Title = "Repository Overview"

	if err = mainTemplate.ExecuteTemplate(w, "listOverview", appState); err != nil {
		log.Fatal(err)
	}

	// TODO: REMOVE
	for _, bug := range appState.CurrentRepo.Bugs {
		fmt.Println(bug.Title)
	}
}
