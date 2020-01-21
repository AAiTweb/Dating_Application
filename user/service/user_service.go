package service

import (
	"github.com/AAiTweb/Dating_Application/entity"
	"github.com/AAiTweb/Dating_Application/user/repository"
)

type UserServiceInstance struct {
	RepositoryInstance repository.UserRepositoryInstance
}
func RegisterUser(user entity.UserPro)bool{
	return true
}
func DeleteUser(user entity.UserPro)bool{
	return true
}
func UpdateUser(user entity.UserPro)bool{
	return true
}