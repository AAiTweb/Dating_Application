package ChatApi

import "github.com/Eyosi-G/Dating_Application/ChatApi/Models"

type APIRepository interface {
	LoadFriendInformation(id int)([]Models.FriendLoadInformation,error)
	UpdateLoginInformation(id int)error
}


//type FakeAPIRepository interface{
//	LoadFriendInformation(id int)([]Models.FriendLoadInformation,error)
//	UpdateLoginInformation(id int)error
//}