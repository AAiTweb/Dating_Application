package repository

import (
	"database/sql"

	"log"

	"github.com/betse/Dating_Application-master/entity"
)

type UserProfileRepositoryImpl struct {
	conn *sql.DB
}

func NewUserProfileRepoImpl(conn *sql.DB) *UserProfileRepositoryImpl {
	return &UserProfileRepositoryImpl{conn: conn}
}
func (pfl *UserProfileRepositoryImpl) UsersProfile() ([]entity.User, error) {
	query := `SELECT * FROM dating_app.user_profile`
	rows, err := pfl.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []entity.User{}
	for rows.Next() {
		user := entity.User{}
		err = rows.Scan(&user.UserId, &user.FirstName, &user.LastName, &user.Country, &user.City, &user.Bio, &user.Dob, &user.ProfPic, &user.Sex)
		if err != nil {
			return nil, err
		}
		users = append(users, user)

	}
	return users, nil

}
func (pfl *UserProfileRepositoryImpl) UserProfile(id uint) ([]entity.User, error) {
	query := `
	SELECT  dating_app.user_profile.profile_user_id,
			dating_app.user_profile.first_name,
			dating_app.user_profile.second_name,
			dating_app.user_profile.country,
			dating_app.user_profile.city,
			dating_app.user_profile.bio,
			dating_app.user_profile.dof,
			dating_app.user_profile.sex,
			dating_app.gallery.picture_path
	
	FROM dating_app.gallery
	INNER JOIN dating_app.user_profile ON dating_app.gallery.picture_owner_id=$1;
`
	rows, err := pfl.conn.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user := []entity.User{}

	for rows.Next() {
		userProfile := entity.User{}
		err = rows.Scan(&userProfile.UserId, &userProfile.FirstName, &userProfile.LastName, &userProfile.Country, &userProfile.City, &userProfile.Bio, &userProfile.Dob, &userProfile.Sex, &userProfile.ProfPicPath)
		if err != nil {
			return nil, err
		}
		user = append(user, userProfile)
	}
	// log.Println(user)

	// user := &entity.User{}

	if err != nil {
		return user, err
	}
	return user, nil

}
func (pfl *UserProfileRepositoryImpl) UpdateProfile(user *entity.User) (*entity.User, error) {
	_, err := pfl.conn.Exec("UPDATE dating_app.user_profile SET first_name=$1,second_name=$2,country=$3,city=$4,bio=$5,dof=$6,profile_picture=$7,sex=$8", user.FirstName, user.LastName, user.Country, user.City, user.Bio, user.Dob, user.ProfPic, user.Sex)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (pfl *UserProfileRepositoryImpl) AddProfile(user *entity.User) (*entity.User, error) {
	default_picture_path := "placeholder.png"
	log.Println("add user")
	_, err := pfl.conn.Exec("INSERT INTO dating_app.gallery(picture_owner_id,picture_path) values($1,$2)", user.UserId, default_picture_path)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("added")

	_, err = pfl.conn.Exec("INSERT INTO dating_app.user_profile(profile_user_id,first_name,second_name,country,city,bio,dof,profile_picture,sex) values($1,$2,$3,$4,$5,$6,$7,$8,$9)", user.UserId, user.FirstName, user.LastName, user.Country, user.City, user.Bio, user.Dob, user.ProfPic, user.Sex)
	log.Println("added")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}

func (pfl *UserProfileRepositoryImpl) DeleteProfile(id uint) (uint, error) {
	_, err := pfl.conn.Exec("DELETE FROM dating_app.user_profile WHERE profile_user_id=$1", id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
