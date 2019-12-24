package service

import (
	"github.com/Eyosi-G/Dating_Application/Api/Models"
	"github.com/Eyosi-G/Dating_Application/Api/repository"
)

type ApiService struct{
	apiRepository repository.ApiRepository
}
func NewApiService(apiRepository repository.ApiRepository)ApiService{
	return ApiService{apiRepository}
}
func (apiservice *ApiService) LoadFriendInformation(id int)([]Models.FriendLoadInformation,error){
	return apiservice.apiRepository.LoadFriendInformation(id)
}