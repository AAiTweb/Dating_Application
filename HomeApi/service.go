package HomeApi

import "github.com/Eyosi-G/Dating_Application/HomeApi/Models"

type HomeApiService interface {
	GetMatches(id int)([]Models.UserMatch,error)
}