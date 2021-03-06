package user

import "github.com/AAiTweb/Dating_Application/entity"

type UserService interface {
	RegisterUser(user entity.User) error
	CheckLogin(user entity.User) error
	DeleteUser(user entity.User) error
	UpdateUser(user entity.User) error
}
