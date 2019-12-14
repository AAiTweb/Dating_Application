function searchBoth(username,lists) {
  let usernameslist = lists;

  for (var usernames of usernameslist) {
    let content = usernames.innerHTML.toLowerCase();
    if (username == "") {
      let card =
        usernames.parentElement.parentElement.parentElement.parentElement;
      
      card.style.display = "";
    } else if (!content.includes(username)) {
      let card =
        usernames.parentElement.parentElement.parentElement.parentElement;
      card.style.display = "none";
    }else{
      let card =
        usernames.parentElement.parentElement.parentElement.parentElement;
      card.style.display = "";
    }
  }
}

function search(selectedValue, input) {
  let list;
  switch (selectedValue) {
    case "1":
      list = document.querySelectorAll(".uname");
      searchBoth(input,list);
      break;
    case "2":
      list = document.querySelectorAll(".country");
      searchBoth(input,list);
      break;
  }
}
