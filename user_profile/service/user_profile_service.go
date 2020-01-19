package service

import (
	// "log"

	"github.com/biniyam112/TheDatingApp/Dating_Application/entity"
	"github.com/biniyam112/TheDatingApp/Dating_Application/user_profile"
)

type UserProfileServiceImpl struct {
	userRepo user_profile.ProfileRepository
}

func NewUserProfileServiceImpl(usrRepo user_profile.ProfileRepository) user_profile.ProfileService {
	return &UserProfileServiceImpl{userRepo: usrRepo}
}

func (ups *UserProfileServiceImpl) UserProfile(id uint) (*entity.UserPro, error) {
	user, err := ups.userRepo.UserProfile(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (ups *UserProfileServiceImpl) UsersProfile() ([]entity.UserPro, error) {
	users, err := ups.UsersProfile()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ups *UserProfileServiceImpl) AddProfile(user *entity.UserPro) (*entity.UserPro, error) {
	user, err := ups.userRepo.AddProfile(user)
	if err != nil {
		// log.Println(err)
		return nil, err

	}
	return user, nil

}
func (ups *UserProfileServiceImpl) DeleteProfile(id uint) (uint, error) {
	id, err := ups.userRepo.DeleteProfile(id)
	if err != nil {
		return 0, err
	}
	return id, nil

}
func (ups *UserProfileServiceImpl) UpdateProfile(user *entity.UserPro) (*entity.UserPro, error) {
	user, err := ups.userRepo.UpdateProfile(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
