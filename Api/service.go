package Api

import "github.com/Eyosi-G/Dating_Application/Api/Models"

type APIService interface {
	LoadFriendInformation(id int)([]Models.FriendLoadInformation,error)
}
