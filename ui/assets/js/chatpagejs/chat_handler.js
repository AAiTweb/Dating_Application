let user1Handler =  function (message,sendtime,user_profile){
    let user1 = $("<div class=\"user1\"></div>");
    let user_container = $("<div class=\"user-container\"></div>");
    let send_time = `<p  class=\"user1-send-time\">${sendtime}</p>`;
    let image_container = $("<div class=\"user1-image\"></div>");
    let profile_image = `<img src="../assets/images/${user_profile}" class="circle-responsive-img">`;
    let message_container = $("<div class=\"chat chat_left message-container\"></div>");
    let message_paragraph =$("<p style=\"flex: 80%\"></p>");
    message_paragraph.text(message);
    message_container.append(message_paragraph)
    image_container.append(profile_image);
    user_container.append(image_container)
    user_container.append(message_container)
    user1.append(user_container,send_time)
    return user1;
}

let user2Handler = function(message,sendtime,user_profile){
    // console.log(user_profile)
    let user2 = $("<div class=\"user2\"></div>");
    let user_container = $("<div class=\"user-container\"></div>");
    let send_time = `<p  class="user2-send-time">${sendtime}</p>`;
    let image_container = $("<div  class=\"user2-image \"></div>");
    let profile_image = `<img src="../assets/images/${user_profile}" class="circle-responsive-img" >`;
    // console.log(profile_image)
    let message_container = $("<div class=\"chat chat_right message-container\"></div>");
    let message_paragraph = $("<p style='flex:80%'></p>");
    message_paragraph.text(message);
    message_container.append(message_paragraph)
    image_container.append(profile_image);
    user_container.append(message_container)
    user_container.append(image_container)
    user2.append(user_container)
    user2.append(send_time)

    return user2;
}
   //  <div class="user2">
   //     <div class="user-container">
   //          <div class="chat chat_right message-container">
   //               <p style="flex:80%">Hey! I'm fine. Thanks for asking!</p>
   //          </div>
   //          <div  class="user2-image">
   //            <img src="../assets/images/profil.png">
   //          </div>
   //    </div>
   //    <p  class="user2-send-time">11:01</p>
   // </div>






// let u = new _users();
// u.useruser1Handler();
// users.append(user1Handler("hi","12:00","profil.png"));