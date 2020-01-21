package user

import "github.com/AAiTweb/Dating_Application/entity"

type UserService interface {
	RegisterUser(user entity.UserPro)bool
	DeleteUser(user entity.UserPro)bool
	UpdateUser(user entity.UserPro)bool
}