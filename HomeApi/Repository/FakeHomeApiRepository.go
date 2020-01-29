package Repository

import (
	"errors"
	"github.com/AAiTweb/Dating_Application/HomeApi/Models"
	"github.com/AAiTweb/Dating_Application/MokeDatabase"
	"strings"
)

type FakeHomeApiRepo struct {
}

func NewFakeHomeApiRepo() FakeHomeApiRepo {
	return FakeHomeApiRepo{}
}
func (f FakeHomeApiRepo) GetMatches(id int) ([]Models.UserMatch, error) {
	    d,ok := MokeDatabase.UsersMatches[id]
	    if ok{
	    	return d,nil
		}
		return nil,errors.New("getMatches error")

}

func (f FakeHomeApiRepo) SearchByName(id int, name string) ([]Models.UserMatch, error) {
	d,ok := MokeDatabase.UsersMatches[id]
	newUserMatches := []Models.UserMatch{}
	if ok{
		for _,val :=  range d{
			lowerUserName := strings.ToLower(val.UserName)
			_lowerUserName  :=  strings.ToLower(name)
			if lowerUserName==_lowerUserName{
				newUserMatches = append(newUserMatches,val)

			}
		}
		return newUserMatches,nil
	}
	return nil,errors.New("searchByName error")
}
