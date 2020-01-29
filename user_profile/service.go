package user_profile

import (
	"github.com/AAiTweb/Dating_Application/entity"
)

type ProfileService interface {
	UsersProfile() ([]entity.UserPro, error)
	UserProfile(id uint) (*entity.UserPro, error)
	DeleteProfile(id uint) (uint, error)
	UpdateProfile(user *entity.UserPro) (*entity.UserPro, error)
	AddProfile(user *entity.UserPro) (*entity.UserPro, error)
}
