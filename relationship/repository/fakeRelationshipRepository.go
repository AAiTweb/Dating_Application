package repository

import (
)

type FakeRelationshipRepository struct {
	//repository map[int]entity.Relationship
}
func NewFakeRelationshipRepo ()FakeRelationshipRepository{
	return FakeRelationshipRepository{}
}

func (f FakeRelationshipRepository) SendRequest(sender, reciever int) int {
	//last := len(f.repository)
	//_newRepo := entity.Relationship{last,sender,reciever,1}
	//f.repository[last] = _newRepo;
	//return last
	return -1
}

func (f FakeRelationshipRepository) AcceptRequest(sender, reciever int) int {
	return -1;
}

