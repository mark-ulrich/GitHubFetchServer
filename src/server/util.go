package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Writes a 400 Bad Request response to an http.ResponseWriter
func processBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("<h2>400 Bad Request</h2>"))
}

// Search for string s in slice of strings.
func stringInSlice(slice []string, s string) bool {
	for _, str := range slice {
		if str == s {
			return true
		}
	}
	return false
}

// Fetch a list of repositories for a given GitHub username
func fetchRepos(username string) (repos *[]Repository, err error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos?per_page=100000", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to fetch repos: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&repos)
	return repos, err
}

func fetchBugs(repo *Repository) {
	if repo.Issues == nil {
		fetchIssues(repo)
	}

	for i, issue := range *repo.Issues {
		fmt.Println("BUG:", issue.Title)
		for _, label := range *issue.Labels {
			if label.Name == "bug" {
				repo.Bugs = append(repo.Bugs, (*repo.Issues)[i])
				repo.BugCount++
				break
			}
		}
	}
}

func fetchIssues(repo *Repository) {
	url := fmt.Sprintf(
		"https://api.github.com/repos/%s/%s/issues?per_page=1000000",
		repo.Owner.Login,
		repo.Name,
	)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("Failed to fetch issues from %s", url))
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&repo.Issues); err != nil {
		panic(err)
	}
	fmt.Println(len(*repo.Issues))
}
