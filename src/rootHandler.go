package main

import "net/http"

// Root handler for the server.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// We must have a valid repo selected
	if currentRepo == nil {
		http.Redirect(w, r, "/update-repo", http.StatusTemporaryRedirect)
	}
}
