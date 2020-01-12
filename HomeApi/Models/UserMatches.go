package Models

import "time"

type UserMatch struct {
	UserId int
	DateOfBirth time.Time
	Country string
	City string
	PicturePath string
	UserName string
}