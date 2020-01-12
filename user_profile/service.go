package user_profile

import (
	"github.com/betse/Dating_Application-master/entity"
)

type ProfileService interface {
	UsersProfile() ([]entity.User, error)
	UserProfile(id uint) (*entity.User, error)
	DeleteProfile(id uint) (uint, error)
	UpdateProfile(user *entity.User) (*entity.User, error)
	AddProfile(user *entity.User) (*entity.User, error)
}
