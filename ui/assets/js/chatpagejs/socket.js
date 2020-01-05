let currentUserId = 1;
let currentUserProfilePicture = "profil.png";

listFriends(currentUserId,currentUserProfilePicture);
changeSelected()

let socket = new WebSocket('ws://localhost:8081/ws');
socket.addEventListener( 'message',function(event){
    let id = $(".selected")
    let msg = JSON.parse(event.data)
    console.log(msg)
    // console.log(currentUserId)
    if(id.attr("id")==msg.SenderId || currentUserId==msg.SenderId ){
        // if receiver is selected or sender
        // id.children()
        let messageBox = user1Handler(msg.MessageText, msg.Time, currentUserProfilePicture);
        $('#no_message').remove()
        $(".users").append(messageBox);
        let height = $(".users")[0].scrollHeight;
        $("#message_field").val("");
        $(".users").animate({scrollTop: height}, 500);
    }else{
        let usr = $("#"+msg.ReceiverId+" .section_right .message_notification")
        if(usr.length){
            usr.text(parseInt(usr.text())+1)
        }else{
            $("#"+ msg.ReceiverId +" .section_right").append($(`<div><span class="message_notification">1</span></div>`))

        }
    }
});
    // console.log(event.data);








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