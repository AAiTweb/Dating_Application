package user

import "github.com/Eyosi-G/Dating_Application/entity"

type UserRepository interface {
	RegisterUser(user entity.User)bool
	DeleteUser(user entity.User)bool
	UpdateUser(user entity.User)bool
}