package notification

import "github.com/AAiTweb/Dating_Application/entity"

type NotifService interface {
	AddNotification(user entity.UserPro) error
	AcceptNotification(user entity.UserPro) error
	RejectNotification(user entity.UserPro) error
}
