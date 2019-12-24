package repository

import (
	"database/sql"
	"github.com/Eyosi-G/Dating_Application/Api/Models"
	"time"
)

type ApiRepository struct {
	db *sql.DB
}
func  NewApiRepository (Db *sql.DB)ApiRepository{
	return ApiRepository{Db}
}
func(ap *ApiRepository)LoadFriendInformation(id int)([]Models.FriendLoadInformation,error){
	row ,err := ap.db.Query(`
select  users.user_id,users.username,login_details.last_activity,t3.picture_path
from (select user_sender_id from relationship where user_reciever_id=$1 and
relationship_status=2 union select user_reciever_id from relationship where user_sender_id=$1 and relationship_status=2) as t1
inner join users on users.user_id = t1.user_sender_id 
inner join login_details on login_details.user_login_id = t1.user_sender_id
inner join (select picture_owner_id,picture_path from user_profile inner join gallery on picture_id=profile_picture) t3 on
t3.picture_owner_id=t1.user_sender_id;`,int(id))
	if err!=nil{
		return nil,err
	}
	var FriendsList []Models.FriendLoadInformation
	for row.Next(){
		eachFriend := Models.FriendLoadInformation{}
		err = row.Scan(&eachFriend.FriendId,&eachFriend.Username,&eachFriend.LastActivity,&eachFriend.ProfilePicture)
		if err!=nil{
			return nil,err
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
	return FriendsList,nil
}
func (ap *ApiRepository) timeDifference(t time.Duration) int{
	tHour:= t.Hours()
	tMin := t.Minutes()
	tSec := t.Seconds()
	if tHour<=0 && tMin<=0 && tSec<=0{
		return 1
	}
	return 0
}