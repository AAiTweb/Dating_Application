package relationship

type RelationshipService interface {
	SendRequest(sender, reciever int)int
	AcceptRequest(sender, reciever int)int
}