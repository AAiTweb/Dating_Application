package entity

import "time"

type UserProfile struct {
	Id int
	FirstName,
	LastName,
	ProfilePicture,
	Sex string
	DateOfBirth         time.Time
	OnlineOfflineStatus bool
}
type HeadData struct {
	Name string
	ProfilePicture string
}
