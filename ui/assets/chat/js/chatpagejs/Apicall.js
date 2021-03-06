let listFriends = function (currentUserId,currentUserProfilePicture){
    $.ajax({
        url: `/user/friends/${currentUserId}`,
        type: 'GET',
        dataType: 'json',
        success: function (data, status) {
            console.log(data)
            let profile = $(".profile");
            profile.empty();
            data.forEach(function(item,index){
                friendBox(item.FriendId,item.Username,item.ProfilePicture,item.UserStatus,currentUserId,currentUserProfilePicture);
            });
        },
        error: function (jqXhr, textStatus, errorMessage) {
            // console.log("hello");
        }
    });

}

let listMessage = function(user1_id,user2_id,profile_picture1,profile_picture2){
    $.ajax({
        url: `/chats/user/${user1_id}/friends/${user2_id}`,
        type: 'GET',
        dataType: 'json',
        success: function (data, status) {
            let users = $(".users");
            users.empty();
            let list = data;
            if (list.length == 0){
                let m = $("<div id='no_message' style='color:grey; font-size: 20px; text-align: center'>There is no messages<br> in this chat yet</div>");
                users.append(m);
            }
            // console.log("....")
            //console.log("hi"+profile_picture2)
            //console.log(profile_picture1)
            list.forEach(function(item,index){

                if (item.FromId == user1_id){
                    let user1 = user1Handler(item.Message,item.SendTime,profile_picture1);
                    let height = $(".users")[0].scrollHeight;
                    users.append(user1);
                    $(".users").animate({scrollTop: height}, 0);

                }else{
                    // console.log(profile_picture2)
                    let user2 = user2Handler(item.Message,item.SendTime,profile_picture2);
                    let height = $(".users")[0].scrollHeight;
                    users.append(user2);
                    $(".users").animate({scrollTop: height}, 100);
                }

            });

        },
        error: function (jqXhr, textStatus, errorMessage) {
            console.log(errorMessage);
        }
    });
}


let updateTime =function (id) {
    $.ajax({
        url:"/user/"+id+"/updatelogin",
        type:"GET",
        success : function (data) {
            console.log("success");
        }
    })
}

let updateStatus = function (id) {
    $.ajax({
        url: `/user/friends/${id}`,
        type: 'GET',
        dataType: 'json',
        success: function (data, status) {
            data.forEach(function(item,index){
               let user =  $("#"+item.FriendId+" .section_right img")
                switch(item.UserStatus){
                    case "offline":
                        user.attr("src","")
                        break;
                    case "online":
                        user.attr("src",`../assets/images/iconfinder_status_46254.ico`)
                        break;
                }
            });
        },
        error: function (jqXhr, textStatus, errorMessage) {
            // console.log("hello");
        }
    });

}