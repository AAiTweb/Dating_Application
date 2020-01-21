package repository

import "github.com/AAiTweb/Dating_Application/entity"

type FakeMessageRepository struct {
	FakeMessages *[]entity.Message
}
func NewFakeMessageRepository(fakeMessages *[]entity.Message)FakeMessageRepository{
	return FakeMessageRepository{FakeMessages:fakeMessages}
}
func (f FakeMessageRepository) SaveMessage(message entity.Message) error {
	*f.FakeMessages  = append(*f.FakeMessages,message)
		return nil
}

func (f FakeMessageRepository) DeleteMessage(message entity.Message) error {
	return nil
}

func (f FakeMessageRepository) Messages(user1 int, user2 int) []entity.Message {
	messages := []entity.Message{}
		for _,val := range *f.FakeMessages{
			if (val.ToId == user1 && val.FromId == user2) ||(val.ToId==user2 && val.FromId==user1){
				messages = append(messages,val)
			}
		}
		return messages
}
