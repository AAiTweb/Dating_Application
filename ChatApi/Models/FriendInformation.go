package Models

import "time"

type FriendLoadInformation struct {
	FriendId       int       `json:FriendId`
	Username       string    `json:Username`
	LastActivity   time.Time `json:LastActivity`
	ProfilePicture string    `json:ProfilePicture`
	UserStatus     string    `json:UserStatus`
}
