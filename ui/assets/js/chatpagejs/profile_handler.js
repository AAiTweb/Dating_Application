let sendMessage = function(senderId,reciverId,currentUserProfilePicture) {
    // console.log('fdfd'+currentUserProfilePicture)

   let messagefield = $("#message_field")
    // console.log()
    messagefield.keypress(function(event){

        let keycode = (event.keyCode ? event.keyCode : event.which);
        if(keycode == '13'){
            let messageText = messagefield.val();
                let jmessage = {
                    SenderId:senderId,
                    ReceiverId:reciverId,
                    MessageText:messageText,
                    SenderPicture:currentUserProfilePicture
                }
                socket.send(JSON.stringify(jmessage))



        }
    });
    let messageSubmitButton = $("#submit");
    messageSubmitButton.on("click", function () {
       let messageText = messagefield.val();
           let jmessage = {
               SenderId:senderId,
               ReceiverId:reciverId,
               MessageText:messageText,
               SenderPicture:currentUserProfilePicture
           }

           socket.send(JSON.stringify(jmessage))




        // socket.send(message)
    });
}

let friendBox = function(id,uname,profile_picture,status,currentUserId,currentUserProfilePicture){
    let profile = $(".profile");
    let section = $(`<div class=\"section\" id=${id}></div>`);
    let left_section = $("<div class=\"section_left\"></div>");
    let image = `<img class=\"circle-responsive-img\" src=\"../assets/images/${profile_picture}\"/>`;
    let username =  `<p>${uname}</p>`;
   // let messageNotification  = $(`<div><span class="message_notification">12</span></div>`);
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
    // right_section.append(messageNotification)
    section.append(left_section,right_section)


    section.on("click",function(){
        // console.log(profile_picture)
        let sections = $(".section");
        for(let i=0;i<sections.length;i++){
            if ($(sections[i]).children().hasClass("section-right-div")){
                $(sections[i]).children().removeClass("section-right-div")
            }
            $(sections[i]).css({"background-color":"rgb(250, 248, 248)"}).attr({"class":"section"})

            let messageSubmitButton = $("#submit");
            let messagefield = $("#message_field");
            messageSubmitButton.unbind()
            messagefield.unbind()

        }
        $(this).css({"background-color":"#ececec"}).attr({"class":"section selected"})
        let rightSelectedDiv = $('<div class="section-right-div"></div>');
        $(this).prepend(rightSelectedDiv);
        let div = $(".selected > .section_right > div");
        div.remove();

        sendMessage(currentUserId,id,currentUserProfilePicture)
        // sendMessage(currentUserId,id)
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

