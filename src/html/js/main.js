// Force a redirect to given url after delaySeconds seconds. The selector
// references a DOM counter element which should display the remaining time
// in seconds until the redirect will occur. 
function redirect(delaySeconds, url, selector) {
  let secondsRemaining = Math.floor(delaySeconds)

  let counter = document.querySelector(selector);
  if (secondsRemaining >= 0)
    counter.innerText = `Redirecting in ${secondsRemaining}...`;
  else
    counter.innerText = "Redirecting..."

  let redirectInterval = setInterval(function () {
    --secondsRemaining;
    counter = document.querySelector(selector);
    if (secondsRemaining === 0) {
      counter.innerText = "Redirecting...";
      window.location.href = url;
      window.clearInterval(redirectTimeout);
    } else {
      counter.innerText = `Redirecting in ${secondsRemaining}...`;
    }

  }, 1000)
}