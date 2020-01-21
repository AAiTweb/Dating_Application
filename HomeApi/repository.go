package HomeApi

import "github.com/AAiTweb/Dating_Application/HomeApi/Models"

type HomeApiRepository interface {
	GetMatches(id int)([]Models.UserMatch,error)
}
