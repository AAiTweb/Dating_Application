package repository

import (
	"database/sql"
	"github.com/Eyosi-G/Dating_Application/entity"
)

type UserRepositoryInstance struct {
	conn *sql.DB
}

func RegisterUser(user entity.User)bool{
	//actual implementation goes here
return true
}
func DeleteUser(user entity.User)bool{
	//actual implementation goes here
return true
}
func UpdateUser(user entity.User)bool{
	//actual implementation goes here
return true
}