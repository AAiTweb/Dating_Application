package service

import (
	"fmt"
	"github.com/Eyosi-G/Dating_Application/entity"
	"github.com/Eyosi-G/Dating_Application/message"
)

type MessageService struct {
	RepoMessage message.MessageRepository
}

func NewMessageService(repomessage message.MessageRepository) message.MessageService{
	return  MessageService{RepoMessage:repomessage}
}

func (m MessageService) SaveMessage(message entity.Message) error {
	err := m.RepoMessage.SaveMessage(message)
		if err==nil{
			fmt.Println("Inserted Successfully")
		}
		return err
}

func (m MessageService) DeleteMessage(message entity.Message) error {
	return nil
}

func (m MessageService) Messages(user1 int, user2 int) []entity.Message {
	messages := m.RepoMessage.Messages(user1,user2)
	return messages
}







//func (ms *MessageService)SaveMessage(message entity.Message) error{
//	err := ms.RepoMessage.SaveMessage(message)
//	if err==nil{
//		fmt.Println("Inserted Successfully")
//	}
//	return err
//}
//func (ms *MessageService)DeleteMessage(message entity.Message)error{
//	return nil
//}
//
//func (ms *MessageService)Messages(user1 int, user2 int)([]entity.Message,error){
//	messages,err := ms.RepoMessage.Messages(user1,user2)
//	return messages,err
//}