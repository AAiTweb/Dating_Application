package repository

import (
	"database/sql"
	"errors"
	"github.com/AAiTweb/Dating_Application/MokeDatabase"
	"github.com/AAiTweb/Dating_Application/entity"
)

type MockNotificationRepo struct {
	conn *sql.DB
}
func FakeUserRepo(db *sql.DB) *Psqlrepo {
	return &Psqlrepo{db}
}


func (m *MockNotificationRepo) AddNotification(userId int) ([]entity.Notification, error) {
	cat := MokeDatabase.Notification
	if cat[1].SenderName != m.GetName(userId){
		err := errors.New("Adding notification failed ")
		return cat,err
	}
	return cat, nil
}

func (m *MockNotificationRepo) AcceptNotification(username string,senderId int) error {
	ctg := MokeDatabase.Notification
	if ctg[1].SenderName != username || ctg[1].SenderName !=m.GetName(senderId){
		err := errors.New("Doesn't match ")
		return  err
	}
	return nil
}

func (m *MockNotificationRepo) RejectNotification(username string, receiverId int) error {
	ctg := MokeDatabase.Notification
	if ctg[1].SenderName != username || ctg[1].SenderName !=m.GetName(receiverId){
		err := errors.New("Doesn't match")
		return  err
	}
	return nil
}

func (m *MockNotificationRepo) GetName(userId int) string{
	var name string
	ctg := MokeDatabase.IdList
	for _,UserLst := range ctg{
		if UserLst.Id == userId{
			name = UserLst.Name
		}
	}
	return name
}

