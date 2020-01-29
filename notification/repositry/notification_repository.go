package repository

import (
	"database/sql"
	"fmt"
	"github.com/AAiTweb/Dating_Application/entity"
	"log"
	"strconv"

)

type Psqlrepo struct {
	Conn *sql.DB
}

func NewUserRepo(db *sql.DB) *Psqlrepo {
	return &Psqlrepo{db}
}

func (pr *Psqlrepo) AddNotification(userId int) ([]entity.Notification, error) {
	query := `SELECT user_sender_id FROM relationship WHERE relationship_status=1 AND user_reciever_id=$1`
	var NotifObj []entity.Notification
	var sender_id int
	var userProfile_pic string
	rows, err := pr.Conn.Query(query,userId)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&sender_id); err != nil {
			fmt.Println(sender_id)
			fmt.Println(err)
		}
		fmt.Println("sender id",sender_id)
		fmt.Println("receiver id ",userId)
		Profilequery := `SELECT picture_path from gallery WHERE picture_owner_id=$1`
		row := pr.Conn.QueryRow(Profilequery,sender_id)
		err = row.Scan(&userProfile_pic)
		fmt.Println("The profile picture",userProfile_pic)
		if err != nil {
			fmt.Println(err)
		}

		notifInstance := entity.Notification{
			SenderName:         pr.GetName(sender_id),
			ProfilePicturePath: "../assets/images/"+userProfile_pic	,
		}
		NotifObj = append(NotifObj,notifInstance)

		fmt.Println(notifInstance.ProfilePicturePath)
	}

	fmt.Println(len(NotifObj))
	return NotifObj, err
}

func (pr *Psqlrepo) GetName(user_id int) string {
	query := `SELECT username from users WHERE user_id=$1`
	rows := pr.Conn.QueryRow(query, user_id)
	var user_name string
	switch err := rows.Scan(&user_name); err {
	case sql.ErrNoRows:
		log.Fatal(err)
	}
	return user_name
}
func (pr *Psqlrepo) GetId(user_name string) int {
	query := `SELECT user_id from users WHERE username=$1`
	rows := pr.Conn.QueryRow(query, user_name)
	var user_id string
	switch err := rows.Scan(&user_id); err {
	case sql.ErrNoRows:
		log.Fatal(err)
	}
	id_value, _ := strconv.Atoi(user_id)
	fmt.Println(id_value)
	return id_value
}

func (pr *Psqlrepo) AcceptNotification(sender_name string,receiverId int) error {
	query := `UPDATE relationship SET relationship_status=$1 WHERE user_sender_id=$2 AND user_reciever_id=$3`
	_, err := pr.Conn.Exec(query, 2, pr.GetId(sender_name),receiverId)
	if err != nil {
		fmt.Print(err)
	}
	return err
}
func (pr *Psqlrepo) RejectNotification(sender_name string,receiverId int) error {
	query := `UPDATE relationship SET relationship_status=$1 WHERE user_sender_id=$2 AND user_reciever_id=$3`
	_, err := pr.Conn.Exec(query, 0, pr.GetId(sender_name),receiverId)
	if err != nil {
		fmt.Print(err)
	}
	return err
}
