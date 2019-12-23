package Api

import (
	"database/sql"
	"encoding/json"
	"github.com/Eyosi-G/Dating_Application/entity"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

type APIHandler struct {
	Db *sql.DB
}
func (ap *APIHandler)Messages( w http.ResponseWriter , r *http.Request){
	path := mux.Vars(r)
	userId,_ := strconv.ParseInt(path["uid"],0,0)
	friendId,_ := strconv.ParseInt(path["fid"],0,0)
	messages := []entity.Message{}
	if r.Method == http.MethodGet{
		row,_ := ap.Db.Query(`select * from messages where message_sender_id=$1 and message_reciever_id=$2 
or message_reciever_id=$1 and message_sender_id=$2;`,userId,friendId)
		for row.Next(){
			message := entity.Message{};
			err := row.Scan(&message.MessageId,&message.FromId,&message.ToId,&message.Message,&message.SendTime)
			if err!=nil{
				return
			}
			messages = append(messages,message)
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(messages)

	}




}
func (ap *APIHandler)Friends (w http.ResponseWriter, r *http.Request){
	// /chat/friends/{id}
	id, _ := strconv.ParseInt(mux.Vars(r)["id"],0,0)
	if r.Method == http.MethodGet{
		row ,err := ap.Db.Query(`select  users.user_id,users.username,login_details.last_activity,t3.picture_path
from (select user_sender_id from relationship where user_reciever_id=$1 and relationship_status=2 union select user_reciever_id from relationship where user_sender_id=$1 and relationship_status=2) as t1
inner join users on users.user_id = t1.user_sender_id 
inner join login_details on login_details.user_login_id = t1.user_sender_id
inner join (select picture_owner_id,picture_path from user_profile inner join gallery on picture_id=profile_picture) t3 on
t3.picture_owner_id=t1.user_sender_id;`,int(id))
		if err!=nil{
			return
		}
		var FriendsList []FriendLoadInformation
		for row.Next(){
			eachFriend := FriendLoadInformation{}
			err = row.Scan(&eachFriend.FriendId,&eachFriend.Username,&eachFriend.LastActivity,&eachFriend.ProfilePicture)
			if err!=nil{
				return
			}
			friendLastActivity:=eachFriend.LastActivity
			tNow := time.Now()
			tNow = tNow.Add(time.Second*-10)
			timeDifference := tNow.Sub(friendLastActivity)
			result := ap.timeDifference(timeDifference)
			switch result {
			case 0:
				eachFriend.UserStatus = "offline"
			case 1:
				eachFriend.UserStatus = "online"
			}
			FriendsList = append(FriendsList,eachFriend)
		}
		a := make(map[string]interface{})
		a["friends"] = FriendsList;
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(a)

	}

	//ap.Db.Query("select * from from where user_id=$1",id)

}
func (ap *APIHandler) timeDifference(t time.Duration) int{
	tHour:= t.Hours()
	tMin := t.Minutes()
	tSec := t.Seconds()
	if tHour<=0 && tMin<=0 && tSec<=0{
		return 1
	}
	return 0
}


func (ap APIHandler)LastSeenUpdater(w http.ResponseWriter, r *http.Request){
	// /user/update_login/{id}
	if r.Method == http.MethodPost{
		id, _ := strconv.ParseInt(r.URL.Query().Get("id"),0,0)
		now := time.Now()
		_,err := ap.Db.Exec("update login_details set last_activity=$1 where login_user_id=$2",now,id)
		if err!=nil{
			return
		}
	}

}
