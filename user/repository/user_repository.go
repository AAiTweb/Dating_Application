package repository

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/biniyam112/Dating_Application/entity"
	"github.com/biniyam112/Dating_Application/form"
	"log"
)

type Psqlrepo struct {
	Conn *sql.DB
}

func NewUserRepo(db *sql.DB) *Psqlrepo {
	return &Psqlrepo{db}
}

//savetodatabase is changed to RegisterUser
func (pr *Psqlrepo) RegisterUser(username string, email string, password string, confirmationtoken string) error {
	encription := md5.New()
	encription.Write([]byte(confirmationtoken))
	token := hex.EncodeToString(encription.Sum(nil))
	fmt.Println(token)
	querystatement := `INSERT INTO users (username,email,password,confirmation_token) VALUES ($1,$2,$3,$4)`
	_, err := pr.Conn.Exec(querystatement, username, email, password, token)
	if err == nil {
		err2 := form.MailToken(token, email)
		if err2 != nil {
			return err2
		}
	}
	return err
}

func (pr *Psqlrepo) CheckLogin(username string) (string, error) {
	querystat := `select password from users where username=$1`
	userinforow := pr.Conn.QueryRow(querystat, username)
	var userpassword string
	err := userinforow.Scan(&userpassword)
	return userpassword, err
}

func (pr Psqlrepo) ValidateToken(vkey string) error {
	var global_error error
	squery := `SELECT is_activated,confirmation_token FROM users WHERE is_activated=0 AND confirmation_token=$1 LIMIT 1`
	row := pr.Conn.QueryRow(squery, vkey)
	var isact int
	var conftok string
	switch err := row.Scan(&isact, &conftok); err {
	case sql.ErrNoRows:
		global_error = err
		return err
		log.Printf("no row found")
	case nil:
		fmt.Println("values retrived are :" + string(isact) + string(conftok))
		squery = `UPDATE users SET is_activated = 1 WHERE confirmation_token=$1`
		_, err = pr.Conn.Exec(squery, conftok)
		global_error = err
		return err
	default:
		panic(err)
	}
	return global_error
}

func (pr Psqlrepo) DeleteUser(username string) error {
	querystat := `delete from users where username=$1`
	_, err := pr.Conn.Query(querystat, username)
	return err
}
func (db *Psqlrepo) UpdateUser(user entity.User) bool {
	//actual implementation goes here
	return true
}
