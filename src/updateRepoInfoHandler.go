package main

import (
	"html/template"
	"log"
	"net/http"
)

type RepoInfo struct {
	Username string
	Name     string
}

// The update-repo handler. A GET request displays a form to enter a GitHub
// username and repository name. A POST request updates the server's
// currentRepoInfo structure.
func updateRepoInfoHandler(w http.ResponseWriter, r *http.Request) {
	var htmlTemplate string

	switch r.Method {
	case "GET":
		htmlTemplate = htmlUpdateRepoGET
	case "POST":
		htmlTemplate = htmlUpdateRepoPOST
	case "HEAD":
		return
	default:
		processBadRequest(w)
		return
	}

	fullTemplate := htmlHeader + htmlTemplate + htmlFooter
	temp := template.Must(template.New("elem").Parse(fullTemplate))

	repoInfo := RepoInfo{}
	if currentUser != nil {
		repoInfo.Username = currentUser.Login
	}
	if currentRepo != nil {
		repoInfo.Name = currentRepo.Name
	}

	if err := temp.Execute(w, repoInfo); err != nil {
		log.Fatal(err)
	}
}
