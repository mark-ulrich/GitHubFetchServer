package main

const htmlHeader = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	
	<style>
	* {
		margin: 0;
		padding: 0;
	}

	body {
		margin: 1rem;
		font-family: sans-serif;
	}

	h2 {
		margin-bottom: .5rem;
	}

	#update-repo-form {
		padding: 1rem;
	}

	input,label {
		padding-top: .5rem;
		padding-bottom: .5rem;
	}

	input {
		padding-left: .5rem;
		padding-right: .5rem;
	}


	.error {
		display: block;
		color: #ff0000;
		margin: 0;
		font-size: .85rem;
		font-family: serif;
		padding: .4rem;
	}
	</style>

  <title>GitHub Fetch Server</title>
</head>
<body>`

const htmlFooter = `
</body>
</html>`

const htmlUpdateRepoGET = `
<h2>Update GitHub Repository</h2>
<form id="update-repo-form" method="POST">
  <label for="username-input">Username:</label>
  <input id="username-input" name="username" placeholder="GitHub username"><br>
  <div class="error" id="username-error"><p>Username not found</p></div>
  <label for="repo-name-input">Repository:</label>
  <input id="repo-name-input" name="repo-name" placeholder="Repository name"><br>
  <div class="error" id="repo-name-error"><p>Repository not found</p></div>
  <button id="update-button" type="submit">Update</button>
</form>

<script type="text/javascript">
let isUsernameValid = false;
let isRepoNameValid = false;
let currentUsername = "";


document.addEventListener("DOMContentLoaded", function(e) {
  updateButton = document.querySelector("#update-button");
  usernameInput = document.querySelector("#username-input");
	repoNameInput = document.querySelector("#repo-name-input");

	updateUpdateButton();

  usernameInput.addEventListener("blur", function(e) {
    validateGitHubUsername(e.target.value);
  })

  repoNameInput.addEventListener("input", function(e) {
    validateGitHubRepoName(e.target.value);
  })
})

// Verify a given username exists on GitHub
function validateGitHubUsername(username) {
  currentUsername = username;

  const xhr = new XMLHttpRequest();
  xhr.open("POST", "/user")
  xhr.setRequestHeader("Content-Type", 'application/json');

  xhr.onload = function (e) {
    isUsernameValid = (xhr.status === 200);
    usernameErrorDiv = document.querySelector("#username-error");
    usernameErrorDiv.style.display = isUsernameValid ? "none" : "block";
    updateUpdateButton();
  };

  xhr.send(JSON.stringify(username));
}

// Verify a given repository exists under the current username
function validateGitHubRepoName(repoName) {

  // We can't look up a repo without a username
  if (currentUsername === "")
    return;

  const xhr = new XMLHttpRequest();
  xhr.open("POST", "/repo");
  xhr.addEventListener("load", function(e) {
		isRepoNameValid = (xhr.status === 200);
		repoNameErrorDiv = document.querySelector("#repo-name-error");
		repoNameErrorDiv.style.display = isRepoNameValid ? "none" : "block";
		updateUpdateButton();
  });
  xhr.send(JSON.stringify(repoName));
}

function updateUpdateButton() {
  updateButton = document.querySelector("#update-button");
  updateButton.disabled = !(isUsernameValid && isRepoNameValid)
}

</script>`

const htmlUpdateRepoPOST = `
<h2>Updated Repo</h2>
<p>Using repo: {{.Username}}/{{.Name}}</p>
<a href="/">Go Back</a>`
