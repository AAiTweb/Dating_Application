package entity

import (
	"time"
)

type User struct {
	UserId    uint
	ProfPic   uint
	FirstName string
	LastName  string
	Country   string
	City      string
	Bio       string
	Sex       string
	Dob       time.Time
}
