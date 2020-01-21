package service

import (
	"github.com/AAiTweb/Dating_Application/ChatApi"
	"github.com/AAiTweb/Dating_Application/ChatApi/Models"
)

type ApiService struct {
	apiRepository ChatApi.APIRepository
}

func NewApiService(apiRepository ChatApi.APIRepository) ChatApi.APIService {
	return ApiService{apiRepository}
}

func (a ApiService) LoadFriendInformation(id int) ([]Models.FriendLoadInformation, error) {
	return a.apiRepository.LoadFriendInformation(id)
}

func (a ApiService) UpdateLoginInformation(id int) error {
	return a.apiRepository.UpdateLoginInformation(id)
}
