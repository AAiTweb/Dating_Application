package repository

import (
	"errors"
	"github.com/AAiTweb/Dating_Application/ChatApi/Models"
	"time"
)

type FakeApiRepository struct {
	FakeFriends *map[int][]Models.FriendLoadInformation
	LoginDetails *[]LoginDetail
}

func (f FakeApiRepository) LoadFriendInformation(id int) ([]Models.FriendLoadInformation, error) {
	if data,ok := (*f.FakeFriends)[id];ok{
				return data,nil
			}
			return nil,errors.New("Invalid Id")
}

func (f FakeApiRepository) UpdateLoginInformation(id int) error {
	now := time.Now();
		for _,val :=  range *(f.LoginDetails){
			if val.UserId == id{
				val.LastActivity = now
				return nil
			}
		}
		return errors.New("Invalid Id")
}

func NewFakeApiRepository (fakefriends *map[int][]Models.FriendLoadInformation,logindetails *[]LoginDetail)FakeApiRepository{
	return FakeApiRepository{fakefriends,logindetails}
}



