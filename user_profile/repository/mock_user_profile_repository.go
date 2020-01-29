package repository

import (
	"database/sql"
	"errors"
	"github.com/AAiTweb/Dating_Application/entity"
	"github.com/AAiTweb/Dating_Application/user_profile"
)

type MockUserProfileRepository struct {
	conn *sql.DB
}
func NewMockUserProfileRepo(db *sql.DB) user_profile.ProfileRepository{
	return &MockUserProfileRepository{db}
}

func (m *MockUserProfileRepository) UsersProfile() ([]entity.UserPro, error) {
	users:=[]entity.UserPro{entity.UserProfileMock}
	return users,nil
}

func (m *MockUserProfileRepository) UserProfile(id uint) (*entity.UserPro, error) {
	user:=entity.UserProfileMock
	if id==1{
		return &user,nil
	}
	return nil, errors.New("not found")
}

func (m *MockUserProfileRepository) DeleteProfile(id uint) (uint, error) {
	user:=entity.UserProfileMock
	if id!=1{
		return 0,errors.New("not found")
	}
	return uint(user.UserId),nil
}

func (m *MockUserProfileRepository) UpdateProfile(user *entity.UserPro) (*entity.UserPro, error) {
	usr :=entity.UserProfileMock
	return  &usr,nil


}

func (m *MockUserProfileRepository) AddProfile(user *entity.UserPro) (*entity.UserPro, error) {
	usr :=entity.UserProfileMock
	return  &usr,nil
}

