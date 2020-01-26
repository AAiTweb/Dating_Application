package repository

import (
	"database/sql"
	"log"
)

type RelationshipRepository struct {
	db *sql.DB
}

func NewRelationshipRepository(Db *sql.DB) RelationshipRepository {
	return RelationshipRepository{Db}
}

func (r RelationshipRepository) SendRequest(sender, reciever int) int {
	query := `insert into relationship(user_sender_id,user_reciever_id,relationship_status) values($1,$2,$3) returning relationship_id;`
	row := r.db.QueryRow(query, sender, reciever, 1)
	var relationshipId int
	err := row.Scan(&relationshipId)
	if err != nil {
		return -1
	}
	log.Println(relationshipId)
	return relationshipId
}

func (r RelationshipRepository) AcceptRequest(sender, reciever int) int {
	query := `update relationship set relationship_status=$1 where user_sender_id=$2 and user_reciever_id=$3 returning relationship_id;`
	row := r.db.QueryRow(query, 2, sender, reciever)
	var relationshipId int
	err := row.Scan(&relationshipId)
	log.Println(relationshipId)
	if err != nil {
		return -1
	}
	return relationshipId
}
