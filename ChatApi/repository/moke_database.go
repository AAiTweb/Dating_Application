package repository

import (
	"github.com/AAiTweb/Dating_Application/ChatApi/Models"
	"github.com/AAiTweb/Dating_Application/entity"
	"time"
)

type  LoginDetail struct {
	LogInID int
	UserId int
	LastActivity time.Time
}


var Messages = []entity.Message{
	{1,1,2,"mesaage 1",  time.Now().Add(time.Minute*-40)},
	{2,5,1,"message 2",  time.Now().Add(time.Minute*-30)},
	{3,1,5,"message 3",  time.Now().Add(time.Minute*-29)},
	{4,9,1,"message 4",  time.Now().Add(time.Minute*-20)},
	{5,1,5,"message 5",  time.Now().Add(time.Minute*-18)},
}


var FriendList = map[int][]Models.FriendLoadInformation{
	   1: {Models.FriendLoadInformation{FriendId: 7, Username: "Sophia", LastActivity: time.Now().Add(time.Minute * -40), ProfilePicture: "p1ic7.jpg", UserStatus: "offline"},
		   Models.FriendLoadInformation{FriendId: 9, Username: "Rosmunda", LastActivity: time.Now().Add(time.Minute * -40), ProfilePicture: "default.jpg", UserStatus: "offline"},
		   Models.FriendLoadInformation{FriendId: 5, Username: "Gabriella", LastActivity: time.Now().Add(time.Minute * -40), ProfilePicture: "p3ic5.jpg", UserStatus: "offline"},
		   Models.FriendLoadInformation{FriendId: 10, Username: "Isabella", LastActivity: time.Now().Add(time.Second * -3), ProfilePicture: "p1ic10.jpg", UserStatus: "online"},
	      },
	   10:{Models.FriendLoadInformation{FriendId:3,Username:"Stephen",LastActivity:time.Now().Add(time.Minute * -40),ProfilePicture: "default.jpg",UserStatus: "offline"},
	   	   Models.FriendLoadInformation{FriendId: 1,Username: "Eddie",LastActivity: time.Now().Add(time.Second * -3),ProfilePicture: "default.jpg",UserStatus: "online"},
	      },
}
var User1FriendList = []Models.FriendLoadInformation{
	{FriendId: 7, Username: "Sophia", LastActivity: time.Now().Add(time.Minute*-40), ProfilePicture: "p1ic7.jpg", UserStatus: "offline"},
	{FriendId: 9, Username: "Rosmunda", LastActivity: time.Now().Add(time.Minute*-40), ProfilePicture: "default.jpg", UserStatus: "offline"},
	{FriendId: 5, Username: "Gabriella", LastActivity: time.Now().Add(time.Minute*-40), ProfilePicture: "p3ic5.jpg", UserStatus: "offline"},
	{FriendId: 10, Username: "Isabella", LastActivity: time.Now().Add(time.Second*-3), ProfilePicture: "p1ic10.jpg", UserStatus: "online"},
}

var LoginDetails = []LoginDetail{
	{LogInID:1,UserId:7,LastActivity:time.Now().Add(time.Minute * -40)},
	{LogInID:2,UserId:9, LastActivity: time.Now().Add(time.Minute * -40)},
	{LogInID:3,UserId:10,LastActivity:time.Now().Add(time.Second * -3)},
	{LogInID:4,UserId:1,LastActivity:time.Now().Add(time.Second * -3)},
}



