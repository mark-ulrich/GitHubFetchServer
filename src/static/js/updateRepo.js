let isUsernameValid = false;
let isRepoNameValid = false;

document.addEventListener("DOMContentLoaded", function (e) {
  updateButton = document.querySelector("#update-button");
  usernameInput = document.querySelector("#username-input");
  repoNameInput = document.querySelector("#repo-name-input");

  updateUpdateButton();

  usernameInput.addEventListener("blur", function (e) {
    validateGitHubUsername(e.target.value);
  })

  repoNameInput.addEventListener("input", function (e) {
    validateGitHubRepoName(e.target.value);
  })
})

// Verify a given username exists on GitHub
function validateGitHubUsername(username) {
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

  let payload = JSON.stringify(repoName)

  const xhr = new XMLHttpRequest();
  xhr.open("POST", "/repo");
  xhr.addEventListener("load", function (e) {
    // Ugly hack to avoid bugs if the current user is presently being updated
    if (xhr.status == 202) {
      xhr.open("POST", "/repo");
      xhr.send(payload)
    }
    isRepoNameValid = (xhr.status === 200);
    repoNameErrorDiv = document.querySelector("#repo-name-error");
    repoNameErrorDiv.style.display = isRepoNameValid ? "none" : "block";
    updateUpdateButton();
  });
  xhr.send(payload);
}

function updateUpdateButton() {
  updateButton = document.querySelector("#update-button");
  updateButton.disabled = !(isUsernameValid && isRepoNameValid)
}