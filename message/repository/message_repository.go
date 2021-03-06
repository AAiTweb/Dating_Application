package repository

import (
	"database/sql"
	"github.com/AAiTweb/Dating_Application/entity"
)

type RepositoryMessage struct {
	db *sql.DB
}

func NewRepositoryMessage(Db *sql.DB) RepositoryMessage {
	return RepositoryMessage{Db}
}

func (r RepositoryMessage) SaveMessage(message entity.Message) error {
	_, err := r.db.Exec(`INSERT  INTO messages (
		message_sender_id,
		message_reciever_id ,
		text_message,
		send_time) 
		VALUES ($1,$2,$3,$4)`, message.FromId, message.ToId, message.Message, message.SendTime)
	return err
}

func (r RepositoryMessage) DeleteMessage(message entity.Message) error {
	_, err := r.db.Exec("DELETE from messages WHERE messages_id=$1", message.MessageId)
	return err
}

func (r RepositoryMessage) Messages(user1 int, user2 int) []entity.Message {
	row, _ := r.db.Query(`select * from messages where message_sender_id=$1 and message_reciever_id=$2 
								or message_reciever_id=$1 and message_sender_id=$2;`, user1, user2)
	messages := []entity.Message{}
	for row.Next() {
		message := entity.Message{}
		err := row.Scan(&message.MessageId, &message.FromId, &message.ToId, &message.Message, &message.SendTime)
		if err != nil {
			return nil
		}
		messages = append(messages, message)
	}
	return messages
}

//func (rm RepositoryMessage)SaveMessage(message entity.Message) error{
//	_, err := rm.db.Exec(`INSERT  INTO messages (
//		message_sender_id,
//		message_reciever_id ,
//		text_message,
//		send_time)
//		VALUES ($1,$2,$3,$4)`, message.FromId,message.ToId,message.Message,message.SendTime)
//	return err
//}
//func (rm RepositoryMessage)DeleteMessage(message entity.Message) error{
//	_,err := rm.db.Exec("DELETE from messages WHERE messages_id=$1",message.MessageId)
//	return err
//}
//func (rm RepositoryMessage)Messages(user1 int, user2 int)([]entity.Message,error){
//	row,_ := rm.db.Query(`select * from messages where message_sender_id=$1 and message_reciever_id=$2
//or message_reciever_id=$1 and message_sender_id=$2;`,user1,user2)
//	messages := []entity.Message{}
//	for row.Next(){
//		message := entity.Message{};
//		err := row.Scan(&message.MessageId,&message.FromId,&message.ToId,&message.Message,&message.SendTime)
//		if err!=nil{
//			return nil,err
//		}
//		messages = append(messages,message)
//	}
//	return messages,nil
//}
