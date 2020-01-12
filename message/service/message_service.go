package service

import (
	"fmt"
	"github.com/biniyam112/Dating_Application/entity"
	"github.com/biniyam112/Dating_Application/message/repository"
)

type MessageService struct {
	RepoMessage repository.RepositoryMessage
}

func (ms *MessageService) SaveMessage(message entity.Message) error {
	err := ms.RepoMessage.SaveMessage(message)

	fmt.Println("Inserted Successfully")
	return err
}
func (ms *MessageService) DeleteMessage(message entity.Message) error {
	return nil
}

func (ms *MessageService) Messages(user1 int, user2 int) ([]entity.Message, error) {
	messages, err := ms.RepoMessage.Messages(user1, user2)
	return messages, err
}
