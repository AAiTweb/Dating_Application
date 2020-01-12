package repository

import (
	"database/sql"
	"github.com/biniyam112/Dating_Application/entity"
)

type RepositoryMessage struct {
	Db *sql.DB
}

func (rm *RepositoryMessage) SaveMessage(message entity.Message) error {
	_, err := rm.Db.Exec(`INSERT  INTO messages (
		from_id,
		to_id ,
		messages,
		send_time) 
		VALUES ($1,$2,$3,$4)`, message.FromId, message.ToId, message.Message, message.SendTime)
	return err
}
func (rm *RepositoryMessage) DeleteMessage(message entity.Message) error {
	_, err := rm.Db.Exec("DELETE from messages WHERE messages_id=$1", message.MessageId)
	return err
}
func (rm *RepositoryMessage) Messages(user1 int, user2 int) ([]entity.Message, error) {
	rows, _ := rm.Db.Query("SELECT  * FROM messages WHERE from_id=$1 AND to_id=$2", user1, user2)
	messages := []entity.Message{}
	for rows.Next() {
		message := entity.Message{}
		err := rows.Scan(&message.MessageId, &message.FromId, &message.ToId, &message.Message, &message.SendTime)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}
