const socket = new WebSocket('ws://localhost:8081/ws');

// Connection opened
socket.addEventListener('open', function (event) {
    socket.send('Hello Server!');
});

// Listen for messages
socket.addEventListener('message', function (event) {
    console.log(event.data)
    let msg = JSON.parse(event.data)
    messageBox = user1Handler(msg.Message, msg.Time, "profil.png");
    $(".users").append(messageBox);
    let height = $(".users")[0].scrollHeight;
    $("#message_field").val("");
    $(".users").animate({scrollTop: height}, 500);
});