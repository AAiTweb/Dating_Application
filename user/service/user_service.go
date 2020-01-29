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
	err := us.UserInstance.RegisterUser(user.UserName, user.Email, user.Password, user.ConfirmationToken)
	if err != nil {
		fmt.Println("Insertion failed")
	}
	return err
}
func (us *UserService) CheckReset(user entity.User) (int, string, string, error) {
	newuser := us.UserInstance
	userId, username, profilePic, err := newuser.CheckReset(user.UserName, user.Password)
	if err != nil {
		fmt.Println(err)
		return -1, "", "", err
	}
	return userId, username, profilePic, nil
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

func (us *UserService) Checkemail(email string) error {
	err := us.UserInstance.Checkemail(email)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) ConfirmReset(key string) error {
	err := us.UserInstance.ConfirmReset(key)
	if err != nil {
		fmt.Println("Confirm error")
		return err
	}
	return nil
}

func (us *UserService) ResetPassword(id int, password string) error {
	err := us.UserInstance.ResetPassword(id, password)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (us *UserService) GetUser(token string) (string, string, error) {
	username, password, err := us.UserInstance.GetUser(token)
	if err != nil {
		fmt.Println("Get user error")
		return username, password, err
	}
	return username, password, nil
}

func (us *UserService) QueFilled(user entity.User) (bool, error) {
	filled, err := us.UserInstance.QueFilled(user.UserName, user.Password)
	if err != nil {
		return filled, err
	}
	return filled, nil
}

func (us *UserService) GetUserId(vkey string) (int, error) {
	Id, err := us.UserInstance.GetUserId(vkey)
	if err != nil {
		return Id, err
	}
	return Id, nil
}

func (us *UserService) GetUserByUserName(username string) (int, error) {
	Id, err := us.UserInstance.GetUserIdbyUsername(username)
	if err != nil {
		return Id, err
	}
	return Id, nil
}
