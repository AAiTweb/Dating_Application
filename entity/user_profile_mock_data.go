package entity

import "time"

var UserProfileMock=UserPro{
	UserId:      1,
	ProfPic:     1,
	ProfPicPath: []string{"new.jpg"},
	FirstName:   "mock user name",
	LastName:    "mock last name",
	Country:     "mock country",
	City:        "mock city",
	Bio:         "mock bio",
	Sex:         "mock sex",
	Dob:         time.Time{},
}