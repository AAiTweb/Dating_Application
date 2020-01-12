package Handler

import (
	"encoding/json"
	"github.com/Eyosi-G/Dating_Application/ChatApi"
	"github.com/Eyosi-G/Dating_Application/message"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

//
//import (
//	"encoding/json"
//	"github.com/Eyosi-G/Dating_Application/ChatApi/Models"
//	"github.com/Eyosi-G/Dating_Application/message/service"
//	"github.com/gorilla/mux"
//	"net/http"
//	"strconv"
//	"time"
//)
//
type APIHandler struct {
	msgService message.MessageService
	apiService ChatApi.APIService
}


func NewApiHandler(msgservice message.MessageService,apiService ChatApi.APIService) APIHandler{
	return APIHandler{msgservice,apiService}
}
func (Apihandler *APIHandler)GetFriends(w http.ResponseWriter, r *http.Request){
	id, _ := strconv.ParseInt(mux.Vars(r)["id"],0,0)
	friendsInformation,_ :=Apihandler.apiService.LoadFriendInformation(int(id))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(friendsInformation)
}
func (Apihandler *APIHandler)GetMessages(w http.ResponseWriter, r *http.Request){
	    if r.Method == http.MethodGet{
			path := mux.Vars(r)
			userId,_ := strconv.ParseInt(path["uid"],0,0)
			friendId,_ := strconv.ParseInt(path["fid"],0,0)
			messages := Apihandler.msgService.Messages(int(userId),int(friendId))
			type json_message struct {
				MessageId,
				FromId,
				ToId int
				Message string
				SendTime string
				Status int
			}
			json_messages := []json_message{}
			for _,msg:= range messages{
				jmessage := json_message{}
				jmessage.MessageId = msg.MessageId
				jmessage.FromId = msg.FromId
				jmessage.ToId = msg.ToId
				jmessage.Message = strings.TrimSuffix(strings.TrimPrefix(msg.Message,"'"),"'")
				jmessage.SendTime = ChatApi.MessageSendTimeChanger(msg.SendTime)
				json_messages = append(json_messages,jmessage)
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(w).Encode(json_messages)
		}


}
func (Apihandler *APIHandler) UpdateLoginDetails(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		path := mux.Vars(r)
		id,_  := strconv.ParseInt(path["id"],0,0)
		Apihandler.apiService.UpdateLoginInformation(int(id))
	}

}

//time := ""
//yearDifference := sendTime.Year() - rightNowTime.Year()
//monthDifference :=  sendTime.Month() - rightNowTime.Month()
//dayDifference := sendTime.Day() - rightNowTime.Day()
//month := MonthConverter(int(sendTime.Month()))
//day  := sendTime.Day()
//year := sendTime.Year()
//hour,min,_ := sendTime.Clock()
//formattedClock := ClockFormatter(hour,min)
//if yearDifference>0{
//time = fmt.Sprintf("%s %s %s %s, %s",month,day,year,formattedClock)
////include the year year-month-day time january 12  2017, 9:20 am
//}else{
//time = fmt.Sprintf("%s %s, %s",month,day,formattedClock)
//
//}

//func (ap *APIHandler)Messages( w http.ResponseWriter , r *http.Request){
//	path := mux.Vars(r)
//	userId,_ := strconv.ParseInt(path["uid"],0,0)
//	friendId,_ := strconv.ParseInt(path["fid"],0,0)
//	messages,_:= ap.msgService.Messages(int(userId),int(friendId))
////	messages := []entity.Message{}
////	if r.Method == http.MethodGet{
////		row,_ := ap.Db.Query(`select * from messages where message_sender_id=$1 and message_reciever_id=$2
////or message_reciever_id=$1 and message_sender_id=$2;`,userId,friendId)
////		for row.Next(){
////			message := entity.Message{};
////			err := row.Scan(&message.MessageId,&message.FromId,&message.ToId,&message.Message,&message.SendTime)
////			if err!=nil{
////				return
////			}
////			messages = append(messages,message)
////		}
//		w.Header().Set("Content-Type", "application/json; charset=utf-8")
//		json.NewEncoder(w).Encode(messages)
//
//	}
//
//
//
//
//}
//func (ap *APIHandler)Friends (w http.ResponseWriter, r *http.Request){
//	// /chat/friends/{id}
//	id, _ := strconv.ParseInt(mux.Vars(r)["id"],0,0)
//	if r.Method == http.MethodGet{
//		row ,err := ap.Db.Query(`select  users.user_id,users.username,login_details.last_activity,t3.picture_path
//from (select user_sender_id from relationship where user_reciever_id=$1 and relationship_status=2 union select user_reciever_id from relationship where user_sender_id=$1 and relationship_status=2) as t1
//inner join users on users.user_id = t1.user_sender_id
//inner join login_details on login_details.user_login_id = t1.user_sender_id
//inner join (select picture_owner_id,picture_path from user_profile inner join gallery on picture_id=profile_picture) t3 on
//t3.picture_owner_id=t1.user_sender_id;`,int(id))
//		if err!=nil{
//			return
//		}
//		var FriendsList []Models.FriendLoadInformation
//		for row.Next(){
//			eachFriend := Models.FriendLoadInformation{}
//			err = row.Scan(&eachFriend.FriendId,&eachFriend.Username,&eachFriend.LastActivity,&eachFriend.ProfilePicture)
//			if err!=nil{
//				return
//			}
//			friendLastActivity:=eachFriend.LastActivity
//			tNow := time.Now()
//			tNow = tNow.Add(time.Second*-10)
//			timeDifference := tNow.Sub(friendLastActivity)
//			result := ap.timeDifference(timeDifference)
//			switch result {
//			case 0:
//				eachFriend.UserStatus = "offline"
//			case 1:
//				eachFriend.UserStatus = "online"
//			}
//			FriendsList = append(FriendsList,eachFriend)
//		}
//		a := make(map[string]interface{})
//		a["friends"] = FriendsList;
//		w.Header().Set("Content-Type", "application/json; charset=utf-8")
//		json.NewEncoder(w).Encode(a)
//
//	}
//
//	//ap.Db.Query("select * from from where user_id=$1",id)
//
//}
//func (ap *APIHandler) timeDifference(t time.Duration) int{
//	tHour:= t.Hours()
//	tMin := t.Minutes()
//	tSec := t.Seconds()
//	if tHour<=0 && tMin<=0 && tSec<=0{
//		return 1
//	}
//	return 0
//}
//
//
//func (ap APIHandler)LastSeenUpdater(w http.ResponseWriter, r *http.Request){
//	// /user/update_login/{id}
//	if r.Method == http.MethodPost{
//		id, _ := strconv.ParseInt(r.URL.Query().Get("id"),0,0)
//		now := time.Now()
//		_,err := ap.Db.Exec("update login_details set last_activity=$1 where login_user_id=$2",now,id)
//		if err!=nil{
//			return
//		}
//	}
//
//}
