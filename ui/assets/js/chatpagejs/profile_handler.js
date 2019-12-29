let sendMessage = function() {
    let messagefield = $("#message_field");
    messagefield.keypress(function(event){
        let keycode = (event.keyCode ? event.keyCode : event.which);
        if(keycode == '13'){
            let message = messagefield.val()
            socket.send(message)
        }
    });
    let messageSubmitButton = $("#submit");
    // let height = users.height();
    messageSubmitButton.on("click", function () {
        let message = messagefield.val()
        socket.send(message)
    });
}

let friendBox = function(id,uname,profile_picture,status,currentUserId,currentUserProfilePicture){
    let profile = $(".profile");
    let section = $(`<div class=\"section\" id=${id}></div>`);
    let left_section = $("<div class=\"section_left\"></div>");
    let image = `<img class=\"circle-responsive-img\" src=\"../assets/images/${profile_picture}\"/>`;
    let username =  `<p>${uname}</p>`;
    let right_section = $("<div class=\"section_right\"></div>");
    let status_image = `<img src="#"/>`;
    switch(status){
        case "offline":
            status_image = `<img src=""/>`
            break;
        case "online":
            status_image = `<img src=\"../assets/images/iconfinder_status_46254.ico\"/>`
            break;
    }
    left_section.append(image,username)
    right_section.append(status_image)
    section.append(left_section,right_section)
    section.on("click",function(){
        // console.log(profile_picture)
        let sections = $(".section");
        for(let i=0;i<sections.length;i++){
            $(sections[i]).css({"background-color":"rgb(250, 248, 248)"})
        }
        // sections.forEach(function(item,index){
        //     item.css({"background-color":"blue"})
        // })
        // let sections = document.getElementsByClassName("section");
        // console.log(sections)
        // sections.forEach(function(item,index){
        //     item.style.backgroundColor = "red"
        // });

        $(this).css({"background-color":"#ececec"})
        listMessage(currentUserId,id,currentUserProfilePicture,profile_picture);
    })
    profile.append(section)
}

let changeSelected = function () {
    let sections = $(".section");
    sections.on("click",function(){
        $(this).css({"background-color":"red"})
    })

}








// let chatContainer = function(users, height){
//     this.profile = null;
//     height = users.users.height();
//     this.
//     this.
//
//
//
//
// }

