package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

type Psqlrepo struct {
	Conn *sql.DB
}

func NewUserRepo(db *sql.DB) *Psqlrepo {
	return &Psqlrepo{db}
}

func (pr *Psqlrepo) AddNotification() ([]string, error) {

	query := `SELECT user_sender_id from relationship WHERE relationship_status=$1`
	var username []string
	var user_id int
	rows, err := pr.Conn.Query(query, 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&user_id); err != nil {
			log.Fatal(err)
		}
		username = append(username, pr.GetName(user_id))

	}
	return username, err
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

func (pr *Psqlrepo) AcceptNotification(sender_name string) error {
	query := `UPDATE relationship SET relationship_status=$1 WHERE user_sender_id=$2`
	_, err := pr.Conn.Exec(query, 2, pr.GetId(sender_name))
	if err != nil {
		fmt.Print(err)
	}
	return err
}
func (pr *Psqlrepo) RejectNotification(sender_name string) error {
	query := `UPDATE relationship SET relationship_status=$1 WHERE user_sender_id=$2`
	_, err := pr.Conn.Exec(query, 0, pr.GetId(sender_name))
	if err != nil {
		fmt.Print(err)
	}
	return err
}
