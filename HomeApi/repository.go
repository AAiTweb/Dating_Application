package HomeApi

import "github.com/biniyam112/TheDatingApp/Dating_Application/HomeApi/Models"

type HomeApiRepository interface {
	GetMatches(id int)([]Models.UserMatch,error)
}
