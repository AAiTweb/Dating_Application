package notification

import "github.com/biniyam112/Dating_Application/entity"

type NotifService interface {
	AddNotification(user entity.User) error
	AcceptNotification(user entity.User) error
	RejectNotification(user entity.User) error
}