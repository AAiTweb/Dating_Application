package HomeApi

import "github.com/AAiTweb/Dating_Application/HomeApi/Models"

type HomeApiService interface {
	GetMatches(id int) ([]Models.UserMatch, error)
	SearchByName(id int,name string)([]Models.UserMatch, error)
}
