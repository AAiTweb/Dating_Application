package ChatApi

import "github.com/biniyam112/TheDatingApp/Dating_Application/ChatApi/Models"

type APIRepository interface {
	LoadFriendInformation(id int)([]Models.FriendLoadInformation,error)
	UpdateLoginInformation(id int)error
}
