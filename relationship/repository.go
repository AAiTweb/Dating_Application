package relationship

type RelationshipRepository interface {
	SendRequest(sender, reciever int) int
	AcceptRequest(sender, reciever int) int
}
