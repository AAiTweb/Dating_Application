package message

import "github.com/Eyosi-G/Dating_Application/entity"

type MessageService interface {
	SaveMessage(message entity.Message) error
	DeleteMessage(message entity.Message) error
	Messages(user1 int, user2 int)[]entity.Message
}