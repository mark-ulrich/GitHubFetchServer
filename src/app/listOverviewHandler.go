package main

import (
	"html/template"
	"log"
	"net/http"
)

func listOverviewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		processBadRequest(w)
		return
	}

	appState.Title = "Repository Overview"
	pageTemplate, err := template.ParseFiles("../html/layout.gohtml", "../html/listOverview.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	if err = pageTemplate.Execute(w, appState); err != nil {
		log.Fatal(err)
	}
}
