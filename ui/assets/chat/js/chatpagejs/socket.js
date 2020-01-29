let token = document.cookie.split("=")[1]
// console.log(token);
let tkobject = JSON.parse(atob(token.split(".")[1]))
console.log(tkobject)
let currentUserId = parseInt(tkobject.id);
let currentUserProfilePicture = tkobject.profile_picture;



listFriends(currentUserId,currentUserProfilePicture);
setInterval(function(){
    updateTime(currentUserId);
    updateStatus(currentUserId)
},11000)


changeSelected()



let socket = new WebSocket('ws://localhost:8081/ws');
socket.addEventListener( 'message',function(event){
    let id = $(".selected")
    let msg = JSON.parse(event.data)
    console.log(msg)
    // console.log(currentUserId)

    if(id.attr("id")==msg.SenderId ) {

        // if receiver is selected or sender

        let messageBox = user2Handler(msg.MessageText, msg.Time, msg.SenderPicture);
        $('#no_message').remove()
        $(".users").append(messageBox);
        let height = $(".users")[0].scrollHeight;
        $("#message_field").val("");
        $(".users").animate({scrollTop: height}, 500);
    }else if(currentUserId==msg.SenderId ){
        let messageBox = user1Handler(msg.MessageText, msg.Time, currentUserProfilePicture);
        $('#no_message').remove()
        $(".users").append(messageBox);
        let height = $(".users")[0].scrollHeight;
        $("#message_field").val("");
        $(".users").animate({scrollTop: height}, 500);

    }else{
        console.log(msg)
        let usr = $("#"+msg.SenderId+" .section_right .message_notification")
        if(usr.length){
            usr.text(parseInt(usr.text())+1)
        }else{
            $("#"+ msg.SenderId +" .section_right").append($(`<div><span class="message_notification">1</span></div>`))
        }
    }
});
    // console.log(event.data);


socket.addEventListener("close",function(evert){
    console.log("socket closing .......")
})


// Connection opened
// socket.addEventListener('open', function (event) {
//     socket.send('Hello Server!');
// });

// Listen for messages
// socket.addEventListener('message', function (event) {
//     // console.log(event.data)
//     let id = $(".selected").attr("id")
//
//     // console.log(id)
//
//     let msg = JSON.parse(event.data)

    // console.log(msg)

    // if(id==msg.SenderId ||  ){
    //     messageBox = user1Handler(msg.MessageText, msg.Time, "profil.png");
    //     $('#no_message').remove()
    //     $(".users").append(messageBox);
    //     let height = $(".users")[0].scrollHeight;
    //     $("#message_field").val("");
    //     $(".users").animate({scrollTop: height}, 500);
    // }
    // console.log(msg)

// });