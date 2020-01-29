package entity

import "time"

type Message struct {
	MessageId,
	FromId,
	ToId int
	Message  string
	SendTime time.Time
}
