package service

import (
	"fmt"
	"github.com/AAiTweb/Dating_Application/entity"
	"github.com/AAiTweb/Dating_Application/user/repository"
)

type UserService struct {
	UserInstance repository.Psqlrepo
}

func NewUserServe(repo *repository.Psqlrepo) *UserService {
	return &UserService{*repo}
}

func (us *UserService) RegisterUser(user entity.User) error {
	fmt.Printf("user name " + user.UserName + " password " + user.Password)
	newuser := us.UserInstance
	err := newuser.RegisterUser(user.UserName, user.Email, user.Password, user.ConfirmationToken)
	if err != nil {
		fmt.Println("Insertion failed")
	}
	return err
}
func (us *UserService) CheckLogin(user entity.User) (int, string, string, error) {
	newuser := us.UserInstance
	userId, username, profilePic, err := newuser.CheckLogin(user.UserName, user.Password)
	if err != nil {
		fmt.Println(err)
		return -1, "", "", err
	}
	return userId, username, profilePic, nil
}

func (us *UserService) ValidateToken(user entity.User) error {
	newuser := us.UserInstance
	err := newuser.ValidateToken(user.ConfirmationToken)
	if err != nil {
		fmt.Println("Token validation failed")
	}
	return err
}

func (us *UserService) DeleteUser(user entity.User) error {
	newuser := us.UserInstance
	err := newuser.DeleteUser(user.UserName)
	if err != nil {
		fmt.Println("UserPro deletion failed")
	}
	return err
}
func (us *UserService) UpdateUser(user entity.User) bool {
	return true
}
