package user

import "github.com/biniyam112/TheDatingApp/Dating_Application/entity"

type UserService interface {
	RegisterUser(user entity.UserPro)bool
	DeleteUser(user entity.UserPro)bool
	UpdateUser(user entity.UserPro)bool
}