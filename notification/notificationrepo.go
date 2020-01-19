package notification

import "github.com/biniyam112/TheDatingApp/Dating_Application/entity"

type NotifRepository interface {
	AddNotification(user entity.UserPro) error
	AcceptNotification(user entity.UserPro) error
	RejectNotification(user entity.UserPro) error
}