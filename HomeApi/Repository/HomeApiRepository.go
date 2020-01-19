package Repository

import (
	"database/sql"
	"github.com/biniyam112/TheDatingApp/Dating_Application/HomeApi/Models"
	"math"
	"time"
)

type HomeApiRepository struct {
	Db *sql.DB
}

func NewHomeApiRepository(db *sql.DB)HomeApiRepository{
	return HomeApiRepository{db}
}

func (h HomeApiRepository) GetMatches(id int) ([]Models.UserMatch, error) {
	query := `select profile_user_id,dof,country,city,t3.picture_path,t4.username,t1.match_point from user_profile
			  inner join (select match_usertwo_id,match_point from matches where match_userone_id=$1 union select match_userone_id,match_point from matches where match_usertwo_id=$1) t1
			  inner join (select picture_owner_id,picture_path from user_profile inner join gallery on picture_id=profile_picture) t3 
          	  inner join users as t4
			  on t4.user_id = t3.picture_owner_id
			  on t3.picture_owner_id=t1.match_usertwo_id
			  on t1.match_usertwo_id = user_profile.profile_user_id;`
	row,err := h.Db.Query(query,id)
	if err!=nil{
		return nil,err
	}
	usermatches := []Models.UserMatch{}
	for row.Next(){
		usermatchtemp := struct {
			UserId int
			DateOfBirth time.Time
			Country string
			City string
			PicturePath string
			UserName string
			MatchPercentage int
		}{}
		err = row.Scan(&usermatchtemp.UserId,&usermatchtemp.DateOfBirth,&usermatchtemp.Country,&usermatchtemp.City,&usermatchtemp.PicturePath,&usermatchtemp.UserName,&usermatchtemp.MatchPercentage)
		if err!=nil{
			return nil,err
		}
		dob := usermatchtemp.DateOfBirth
		birthday := time.Date(dob.Year(), dob.Month(), dob.Day(), 0, 0, 0, 0, time.UTC)
		today := time.Now()
		age := math.Floor(today.Sub(birthday).Hours() / 24 / 365)
		usermatch := Models.UserMatch{
			UserId:          usermatchtemp.UserId,
			Age:             int(age),
			Country:         usermatchtemp.Country,
			City:            usermatchtemp.City,
			PicturePath:     usermatchtemp.PicturePath,
			UserName:        usermatchtemp.UserName,
			MatchPercentage: usermatchtemp.MatchPercentage,
		}
		usermatches = append(usermatches,usermatch)
	}
	return usermatches,nil
}
