package HomeApi

import "github.com/Eyosi-G/Dating_Application/HomeApi/Models"

type HomeApiRepository interface {
	GetMatches(id int)([]Models.UserMatch,error)
}
