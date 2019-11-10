function usernameSearch(username) {
  console.log("username search" + username);
}
function locationSearch(location) {
  console.log("location search" + location);
}

function search(selectedValue, input) {
  switch (selectedValue) {
    case "1":
      usernameSearch(input);
      break;
    case "2":
      locationSearch(input);
      break;
  }
}
