package notification

import "github.com/AAiTweb/Dating_Application/entity"

type NotifRepository interface {
	AddNotification(userId int) ([]entity.Notification,error)
	AcceptNotification(relation entity.Relationship,receiverId int) error
	RejectNotification(relation entity.Relationship,receiverId int) error
	GetName(senderId int)
}
