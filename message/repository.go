package message

import "github.com/AAiTweb/Dating_Application/entity"


type MessageRepository interface {
	SaveMessage(message entity.Message) error
	DeleteMessage(message entity.Message) error
	Messages(user1 int, user2 int)[]entity.Message
}

//type FakeMessageRepository interface {
//	SaveMessage(message entity.Message) error
//	DeleteMessage(message entity.Message) error
//	Messages(user1 int, user2 int)[]entity.Message
//}