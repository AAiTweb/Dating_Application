package Login_SignupApi

import "time"

type FriendLoadInformation struct {
	FriendId       int
	Username       string
	LastActivity   time.Time
	ProfilePicture string
	UserStatus     string
}
