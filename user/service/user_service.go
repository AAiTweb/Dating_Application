package service

import (
	"github.com/Eyosi-G/Dating_Application/entity"
	"github.com/Eyosi-G/Dating_Application/user/repository"
)

type UserServiceInstance struct {
	RepositoryInstance repository.UserRepositoryInstance
}
func RegisterUser(user entity.User)bool{
	return true
}
func DeleteUser(user entity.User)bool{
	return true
}
func UpdateUser(user entity.User)bool{
	return true
}