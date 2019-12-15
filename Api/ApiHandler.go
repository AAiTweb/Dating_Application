package Api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type APIHandler struct {
	Db *sql.DB
}

func (ap *APIHandler)friends (w http.ResponseWriter, r *http.Request){
	id, _ := strconv.ParseInt(r.URL.Query().Get("id"),0,0)
	if r.Method == http.MethodGet{
		row ,err := ap.Db.Query(`select login_details.login_user_id,users.username,login_details.last_activity,
		profile.profile_picture
		from (select user_two_id from relationship where user_one_id=$1) t1
			inner join login_details on login_details.login_user_id=t1.user_two_id
			inner join profile on profile.profile_user_id = t1.user_two_id
			inner join users on users.id = t1.user_two_id;`,int(id))
		if err!=nil{
			return
		}
		var FriendsList = []FriendLoadInformation{}
		for row.Next(){
			fli := FriendLoadInformation{}
			err = row.Scan(&fli.Login_user_id,&fli.Username,&fli.Last_activity,&fli.Profile_picture)
			if err!=nil{
				return
			}
			t:=fli.Last_activity
			tNow := time.Now()
			tNow = tNow.Add(time.Second*-10)
			tdiff := tNow.Sub(t)
			result := ap.timeDefference(tdiff)
			switch result {
			case 0:
				fli.UserStatus = "offline"
			case 1:
				fli.UserStatus = "online"
			}
			FriendsList = append(FriendsList,fli)
		}
		FriendsListJson := json.NewEncoder(w)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		err = FriendsListJson.Encode(FriendsList)
		if err!=nil{
			return
		}

	}

	//ap.Db.Query("select * from from where user_id=$1",id)


}

func (ap *APIHandler) timeDefference(t time.Duration) int{
	tHour:= t.Hours()
	tMin := t.Minutes()
	tSec := t.Seconds()

	if tHour<=0{
		if tMin<=0{
			if tSec<=0{
				return 1
			}else{
				//offline
				return 0
			}
		}else{
			//offline
			return 0
		}
	}else{
		//offline
		return 0
	}



}
