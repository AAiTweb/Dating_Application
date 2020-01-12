package user

import "github.com/biniyam112/Dating_Application/entity"

type UserRepository interface {
	RegisterUser(user entity.User) error
	CheckLogin(user entity.User) error
	DeleteUser(user entity.User) error
	UpdateUser(user entity.User) error
}
