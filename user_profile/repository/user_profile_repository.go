package repository

import (
	"database/sql"
	"github.com/AAiTweb/Dating_Application/entity"
	"log"
	"time"
)

type UserProfileRepositoryImpl struct {
	conn *sql.DB
}

func NewUserProfileRepoImpl(conn *sql.DB) *UserProfileRepositoryImpl {
	return &UserProfileRepositoryImpl{conn: conn}
}
func (pfl *UserProfileRepositoryImpl) UsersProfile() ([]entity.UserPro, error) {
	query := `SELECT * FROM user_profile`
	rows, err := pfl.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []entity.UserPro{}
	for rows.Next() {
		user := entity.UserPro{}
		err = rows.Scan(&user.UserId, &user.FirstName, &user.LastName, &user.Country, &user.City, &user.Bio, &user.Dob, &user.ProfPic, &user.Sex,&user.Dob)
		if err != nil {
			return nil, err
		}
		users = append(users, user)

	}

	return users, nil

}
func (pfl *UserProfileRepositoryImpl) UserProfile(id uint) (*entity.UserPro, error) {
	query := `SELECT * FROM user_profile WHERE profile_user_id=$1`

	galleryQuery := `SELECT picture_path FROM gallery WHERE picture_owner_id=$1`
	anonymousUser := struct {
		ProfId  uint
		UserId  uint64
		ProfPic uint

		FirstName string
		LastName  string
		Country   string
		City      string
		Bio       string
		Sex       string
		Dob       time.Time
	}{}

	row := pfl.conn.QueryRow(query, id)
	err := row.Scan(&anonymousUser.ProfId, &anonymousUser.UserId, &anonymousUser.FirstName, &anonymousUser.LastName, &anonymousUser.Country, &anonymousUser.City, &anonymousUser.Bio, &anonymousUser.Dob, &anonymousUser.ProfPic, &anonymousUser.Sex)

	if err != nil {
		return nil, err
	}

	gImages := []string{}

	rows, err := pfl.conn.Query(galleryQuery, id)
	for rows.Next() {
		var imgPath string
		err = rows.Scan(&imgPath)

		if err != nil {
			return nil, err
		}
		gImages = append(gImages, imgPath)
	}
	user := &entity.UserPro{anonymousUser.UserId, anonymousUser.ProfPic, gImages, anonymousUser.FirstName, anonymousUser.LastName, anonymousUser.Country, anonymousUser.City, anonymousUser.Bio, anonymousUser.Sex, anonymousUser.Dob}

	if err != nil {
		return user, err
	}
	log.Println(user)
	return user, nil

}
func (pfl *UserProfileRepositoryImpl) UpdateProfile(user *entity.UserPro) (*entity.UserPro, error) {
	id:= pfl.conn.QueryRow("INSERT INTO gallery(picture_owner_id,picture_path) values($1,$2) RETURNING picture_id", user.UserId, user.ProfPicPath[0])
	var returning uint

	id.Scan(&returning)

	log.Println(returning,user.ProfPicPath,user.UserId)




	//picId := pfl.conn.QueryRow("SELECT picture_id FROM gallery WHERE picture_owner_id = $1", user.UserId)

	//err = picId.Scan(&user.ProfPic)
	//if err != nil {
	//	return nil, err
	//}

	_, err := pfl.conn.Exec("UPDATE user_profile SET first_name=$1,second_name=$2,country=$3,city=$4,bio=$5,dof=$6,profile_picture=$7,sex=$8 WHERE profile_user_id=$9", user.FirstName, user.LastName, user.Country, user.City, user.Bio, user.Dob, returning, user.Sex,user.UserId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (pfl *UserProfileRepositoryImpl) AddProfile(user *entity.UserPro) (*entity.UserPro, error) {

	default_picture_path := "default.jpg"
	// log.Println("add user")
	_, err := pfl.conn.Exec("INSERT INTO gallery(picture_owner_id,picture_path) values($1,$2)", user.UserId, default_picture_path)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	// log.Println("added")
	picId := pfl.conn.QueryRow("SELECT picture_id FROM gallery WHERE picture_owner_id = $1", user.UserId)

	err = picId.Scan(&user.ProfPic)
	if err != nil {
		return nil, err
	}

	_, err = pfl.conn.Exec("INSERT INTO user_profile(profile_user_id,first_name,second_name,country,city,bio,dof,profile_picture,sex) values($1,$2,$3,$4,$5,$6,$7,$8,$9)", user.UserId, user.FirstName, user.LastName, user.Country, user.City, user.Bio, user.Dob, user.ProfPic, user.Sex)
	// log.Println("added")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}

func (pfl *UserProfileRepositoryImpl) DeleteProfile(id uint) (uint, error) {
	_, err := pfl.conn.Exec("DELETE FROM user_profile WHERE profile_user_id=$1", id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
