package entity

import (
	"time"
)

type User struct {
	UserId      uint64
	ProfPic     uint
	ProfPicPath []string
	FirstName   string
	LastName    string
	Country     string
	City        string
	Bio         string
	Sex         string
	Dob         time.Time
}
