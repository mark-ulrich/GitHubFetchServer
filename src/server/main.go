// A simple web server which queries GitHub for a given repository, and allows
// navigation of the list of bug reports, milestones, and users.
package main

import (
	"log"
	"net/http"
	"text/template"
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
	templatePath = "../templates/"
)

// Route paths
const (
	ListRoutePath           = "/list/"
	ListOverviewRoutePath   = "/list/overview"
	ListBugsRoutePath       = "/list/bugs"
	ListMilestonesRoutePath = "/list/milestones"
	ListUsersRoutePath      = "/list/users"
	UpdateRepoInfoRoutePath = "/update-repo"
	UpdateRepoRoutePath     = "/repo"
	UpdateUserRoutePath     = "/user"
)

var (
	appState     AppState
	mainTemplate *template.Template
	err          error
)

func main() {

	loadTemplates()

	// Serve static assets
	fileServer := http.FileServer(http.Dir("../static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// ===============================================
	//									Routes
	// ===============================================

	// Root handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/list/overview", http.StatusMovedPermanently)
	})

	// Update user and repo
	http.HandleFunc(UpdateRepoInfoRoutePath, updateRepoInfoHandler)
	http.HandleFunc(UpdateUserRoutePath, updateUserHandler)
	http.HandleFunc(UpdateRepoRoutePath, updateRepoHandler)

	// List repo info
	http.HandleFunc(ListRoutePath, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, ListOverviewRoutePath, http.StatusMovedPermanently)
	})
	http.HandleFunc(ListOverviewRoutePath, listOverviewHandler)
	http.HandleFunc(ListBugsRoutePath, listBugsHandler)
	http.HandleFunc(ListMilestonesRoutePath, listMilestonesHandler)
	http.HandleFunc(ListUsersRoutePath, listUsersHandler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func loadTemplates() {
	mainTemplate, err = template.ParseGlob(templatePath + "*.gohtml")
	if err != nil {
		panic(err)
	}
}
