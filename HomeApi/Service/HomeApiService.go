package Service

import (
	"github.com/Eyosi-G/Dating_Application/HomeApi"
	"github.com/Eyosi-G/Dating_Application/HomeApi/Models"
)

type HomeApiService struct {
	HomeApiRepositroy HomeApi.HomeApiRepository
}

func NewHomeApiService(repository HomeApi.HomeApiRepository)HomeApi.HomeApiService{
	return HomeApiService{repository}
}

func (h HomeApiService) GetMatches(id int) ([]Models.UserMatch, error) {
	usermatches,err := h.HomeApiRepositroy.GetMatches(id)
	if err!=nil{
		return nil,err
	}
	return usermatches,nil
}

