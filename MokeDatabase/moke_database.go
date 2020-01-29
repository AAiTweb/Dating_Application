package MokeDatabase

import (
	"github.com/AAiTweb/Dating_Application/ChatApi/Models"
	HomeModels "github.com/AAiTweb/Dating_Application/HomeApi/Models"
	"github.com/AAiTweb/Dating_Application/entity"
	"time"
)

type LoginDetail struct {
	LogInID      int
	UserId       int
	LastActivity time.Time
}

var Messages = []entity.Message{
	{1, 1, 2, "mesaage 1", time.Now().Add(time.Minute * -40)},
	{2, 5, 1, "message 2", time.Now().Add(time.Minute * -30)},
	{3, 1, 5, "message 3", time.Now().Add(time.Minute * -29)},
	{4, 9, 1, "message 4", time.Now().Add(time.Minute * -20)},
	{5, 1, 5, "message 5", time.Now().Add(time.Minute * -18)},
}

var FriendList = map[int][]Models.FriendLoadInformation{
	1: {Models.FriendLoadInformation{FriendId: 7, Username: "Sophia", LastActivity: time.Now().Add(time.Minute * -40), ProfilePicture: "p1ic7.jpg", UserStatus: "offline"},
		Models.FriendLoadInformation{FriendId: 9, Username: "Rosmunda", LastActivity: time.Now().Add(time.Minute * -40), ProfilePicture: "default.jpg", UserStatus: "offline"},
		Models.FriendLoadInformation{FriendId: 5, Username: "Gabriella", LastActivity: time.Now().Add(time.Minute * -40), ProfilePicture: "p3ic5.jpg", UserStatus: "offline"},
		Models.FriendLoadInformation{FriendId: 10, Username: "Isabella", LastActivity: time.Now().Add(time.Second * -3), ProfilePicture: "p1ic10.jpg", UserStatus: "online"},
	},
	10: {Models.FriendLoadInformation{FriendId: 3, Username: "Stephen", LastActivity: time.Now().Add(time.Minute * -40), ProfilePicture: "default.jpg", UserStatus: "offline"},
		Models.FriendLoadInformation{FriendId: 1, Username: "Eddie", LastActivity: time.Now().Add(time.Second * -3), ProfilePicture: "default.jpg", UserStatus: "online"},
	},
}
var User1FriendList = []Models.FriendLoadInformation{
	{FriendId: 7, Username: "Sophia", LastActivity: time.Now().Add(time.Minute * -40), ProfilePicture: "p1ic7.jpg", UserStatus: "offline"},
	{FriendId: 9, Username: "Rosmunda", LastActivity: time.Now().Add(time.Minute * -40), ProfilePicture: "default.jpg", UserStatus: "offline"},
	{FriendId: 5, Username: "Gabriella", LastActivity: time.Now().Add(time.Minute * -40), ProfilePicture: "p3ic5.jpg", UserStatus: "offline"},
	{FriendId: 10, Username: "Isabella", LastActivity: time.Now().Add(time.Second * -3), ProfilePicture: "p1ic10.jpg", UserStatus: "online"},
}

var LoginDetails = []LoginDetail{
	{LogInID: 1, UserId: 7, LastActivity: time.Now().Add(time.Minute * -40)},
	{LogInID: 2, UserId: 9, LastActivity: time.Now().Add(time.Minute * -40)},
	{LogInID: 3, UserId: 10, LastActivity: time.Now().Add(time.Second * -3)},
	{LogInID: 4, UserId: 1, LastActivity: time.Now().Add(time.Second * -3)},
}

var UserProfile = map[int]entity.UserPro{
	1:{ UserId:1,FirstName:"Eddie",LastName:"EddieDad",Country:"Ethiopia",City:"Addis Ababa",Bio:"love love love",Sex:"Male",Dob :time.Date(1999,01,02,0,0,0,0,time.UTC),ProfPic:7},
	2:{ UserId:3,FirstName:"Stephen",LastName:"StephenDad",Country:"USA",City:"Texas",Bio:"love food",Sex:"Male",Dob :time.Date(1998,01,02,0,0,0,0,time.UTC),ProfPic:7},
	3:{ UserId:4,FirstName:"Emma",LastName:"EddieDad",Country:"EmmaDad",City:"New York",Bio:"people person",Sex:"Female",Dob :time.Date(2006,01,02,0,0,0,0,time.UTC),ProfPic:7},
}

var Gallery = map[int]entity.Gallery{
	1:{5,"p1ic5.jpg"},
	2:{1,"p1ic7.jpg"},
	3:{10,"p1ic10.jpg"},
	4:{5,"p2ic5.jpg"},
	5:{5,"p3ic5.jpg"},
}

var Matches = map[int]entity.Match{
	1:{1,3,10},
	2:{1,4,70},
	3:{1,5,90},
	4:{1,6,50},
}

var Users = map[int]entity.User{
	1:{1,"Eddie","aDAJ9169LC","bdisdel0@biglobe.ne.jp","hfkdklerfkklsxkwe"},
	2:{2,"Jimmy","aDAJ9169LC","bdisdel0@biglobe.ne.jp","hfkdklerfkklsxkwe"},
	3:{3,"Stephen","aDAJ9169LC","bdisdel0@biglobe.ne.jp","hfkdklerfkklsxkwe"},
	4:{4,"Emma","aDAJ9169LC","bdisdel0@biglobe.ne.jp","hfkdklerfkklsxkwe"},
	5:{5,"Gabriella","aDAJ9169LC","bdisdel0@biglobe.ne.jp","hfkdklerfkklsxkwe"},
	6:{6,"Adam","aDAJ9169LC","bdisdel0@biglobe.ne.jp","hfkdklerfkklsxkwe"},
	7:{7,"Sophia","aDAJ9169LC","bdisdel0@biglobe.ne.jp","hfkdklerfkklsxkwe"},

}

var Relationships = map[int]entity.Relationship{
	1:{ 1,7,3,1},
	2:{ 2,9,3,2},
	3:{ 3,3,10,1},
	4:{ 4,4,8,2},
	5:{ 5,4,6,2},
	6:{ 6,1,4,2},
	7:{ 7,1,7,2},
	8:{ 8,1,10,2},
}

var UsersMatches  = map[int][]HomeModels.UserMatch{
	1 : {{3 ,32, "USA" ,"Texas", "default.jpg", "Stephen", 70},
		 {5 ,29, "Mexico", "Mexico" ,"p3ic5.jpg", "Gabriella", 70},
	},
}

var Notification =[]entity.Notification{
		{"Eddie","p1ic5.jpg"},
		{"Stephen","Stephen1.png"},
		{"Emma","p1ic7.jpg"},
		{"Gabriella","p2ic5.jpg"},
		{"Sophia","p2ic5.jpg"},
}

type IdGetter struct{
	Id int
	Name string
}

var IdList = []IdGetter{
	{12,"Eddie"},{34,"Godzilla"},{34,"Mia"},
}
