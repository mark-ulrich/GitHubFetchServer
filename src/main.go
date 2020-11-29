// A simple web server which queries GitHub for a given repository, and allows
// navigation of the list of bug reports, milestones, and users.
package main

import (
	"log"
	"net/http"
)

var (
	currentUser *User
	currentRepo *Repository
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/update-repo", updateRepoInfoHandler)
	http.HandleFunc("/user", updateUserHandler)
	http.HandleFunc("/repo", updateRepoHandler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
