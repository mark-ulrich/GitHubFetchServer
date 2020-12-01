// A simple web server which queries GitHub for a given repository, and allows
// navigation of the list of bug reports, milestones, and users.
package main

import (
	"log"
	"net/http"
)

type PageInfo struct {
	Title   string
	Request *http.Request
}

type AppState struct {
	CurrentUser *User
	CurrentRepo *Repository
	PageInfo
}

const (
	UpdateRepoHandlerPath = "/update-repo"
	UpdateRepoPath        = "/repo"
	UpdateUserPath        = "/user"
)

var (
	appState AppState
)

func main() {

	// Serve static assets
	fileServer := http.FileServer(http.Dir("../html/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Set up routes
	http.HandleFunc("/", rootHandler)
	http.HandleFunc(UpdateRepoHandlerPath, updateRepoInfoHandler)
	http.HandleFunc(UpdateUserPath, updateUserHandler)
	http.HandleFunc(UpdateRepoPath, updateRepoHandler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
