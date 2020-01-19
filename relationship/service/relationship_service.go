package service

import "github.com/biniyam112/TheDatingApp/Dating_Application/relationship"


type RelationshipService struct {
	relationshipRepository relationship.RelationshipRepository
}

func NewRelationshipService(repository relationship.RelationshipRepository)relationship.RelationshipService{
	return RelationshipService{repository}
}

func (r RelationshipService) SendRequest(sender, reciever int) int {
	return r.relationshipRepository.SendRequest(sender,reciever)
}

func (r RelationshipService) AcceptRequest(sender, reciever int) int {
	return r.relationshipRepository.AcceptRequest(sender,reciever)
}


