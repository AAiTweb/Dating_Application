package service

import (
	"fmt"
	"github.com/AAiTweb/Dating_Application/entity"
	"github.com/AAiTweb/Dating_Application/notification/repositry"
)

type NotifService struct {
	NotifInstance repository.Psqlrepo
}

func NewNotifServe(repo *repository.Psqlrepo) *NotifService {
	return &NotifService{*repo}
}

func (ns NotifService) AddNotification(userId int) ([]entity.Notification, error) {
	newNotif := ns.NotifInstance
	user_array, err := newNotif.AddNotification(userId)
	if err != nil {
		fmt.Print(err)
	}
	return user_array, err
}
func (ns NotifService) AcceptNotification(relation entity.Relationship,receiverId int) error {
	newNotif := ns.NotifInstance
	err := newNotif.AcceptNotification(ns.NotifInstance.GetName(relation.SenderId),receiverId)
	if err != nil {
		fmt.Print(err)
	}
	return err
}
func (ns NotifService) RejectNotification(relation entity.Relationship,receiverId int) error {
	newNotif := ns.NotifInstance
	err := newNotif.RejectNotification(ns.NotifInstance.GetName(relation.SenderId),receiverId)
	if err != nil {
		fmt.Print(err)
	}
	return err
}
