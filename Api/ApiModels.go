package Api

import "time"

type FriendLoadInformation struct {
	Login_user_id int
	Username string
	Last_activity time.Time
	Profile_picture string
	UserStatus string
}

