package repository

import (
	"database/sql"
	"github.com/AAiTweb/Dating_Application/entity"
)

type UserRepositoryInstance struct {
	conn *sql.DB
}

func RegisterUser(user entity.UserPro) bool {
	//actual implementation goes here
	return true
}
func DeleteUser(user entity.UserPro) bool {
	//actual implementation goes here
	return true
}
func UpdateUser(user entity.UserPro) bool {
	//actual implementation goes here
	return true
}
