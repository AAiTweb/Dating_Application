package repository

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/AAiTweb/Dating_Application/entity"
	"github.com/AAiTweb/Dating_Application/form"
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
		err = form.MailToken(token, email)
		if err != nil {
			_ = pr.RemoveUser(username)
			return err
		}
	}
	return nil
}

func (pr *Psqlrepo) RemoveUser(username string) error {
	query := `DELETE FROM users WHERE username=$1`
	_,err := pr.Conn.Exec(query,username)
	if err != nil{
		return err
	}
	return nil
}

func (pr *Psqlrepo) CheckLogin(username string, password string) (int, string, string, error) {
	row := pr.Conn.QueryRow(`select t1.user_id,t1.username,t3.picture_path
    from (select user_id, username from users where username=$1) as t1
      inner join (select picture_owner_id,picture_path from user_profile inner join gallery on picture_id=profile_picture) t3 on
        t3.picture_owner_id=t1.user_id;`, username)

	usr := struct {
		Id int
		UserName,
		ProfilePicture string
	}{}
	err := row.Scan(&usr.Id, &usr.UserName, &usr.ProfilePicture)
	log.Println(usr)

	if err != nil {
		return -1, "", "", errors.New("User Doesn't Exist")
	}
	var hashedPassword string
	fmt.Println(hashedPassword)
	query := `SELECT password FROM users WHERE username=$1`
	row = pr.Conn.QueryRow(query,username)
	err = row.Scan(&hashedPassword)
	isEqual := form.CompareHashToPassword(password,hashedPassword)
	if !isEqual{
		log.Println(err)
		return -1, "", "", err
	}
	return usr.Id, usr.UserName, usr.ProfilePicture, nil
}

func (pr *Psqlrepo) CheckReset(username string, password string) (int, string, string, error) {
	row := pr.Conn.QueryRow(`select t1.user_id,t1.username,t3.picture_path
    from (select user_id, username from users where username=$1) as t1
      inner join (select picture_owner_id,picture_path from user_profile inner join gallery on picture_id=profile_picture) t3 on
        t3.picture_owner_id=t1.user_id;`, username)

	usr := struct {
		Id int
		UserName,
		ProfilePicture string
	}{}
	err := row.Scan(&usr.Id, &usr.UserName, &usr.ProfilePicture)
	log.Println(usr)

	if err != nil {
		return -1, "", "", errors.New("User Doesn't Exist")
	}
	return usr.Id, usr.UserName, usr.ProfilePicture, nil
}

func (pr *Psqlrepo) ValidateToken(vkey string) error {
		squery := `UPDATE users SET is_activated = 1 WHERE confirmation_token=$1`
		_,err := pr.Conn.Exec(squery, vkey)
		if err != nil {
			log.Println("Token validation failed")
			return err
		}
		return nil
}

func (pr *Psqlrepo) DeleteUser(username string) error {
	querystat := `delete from users where username=$1`
	_, err := pr.Conn.Query(querystat, username)
	return err
}
func (pr *Psqlrepo) UpdateUser(user entity.User) bool {
	//actual implementation goes here
	return true
}

func (pr *Psqlrepo) Checkemail(email string) error {
	query := `SELECT confirmation_token FROM users WHERE email=$1`
	row := pr.Conn.QueryRow(query,email)
	var confToken string
	switch err := row.Scan(&confToken); err {
	case sql.ErrNoRows:
		err = errors.New("1") //no rows error
		return err
	}
	err := form.MailResetPassword(email,confToken)
	if err != nil {
		query := `UPDATE users SET password_reset=0 WHERE email=$1`
		_,err := pr.Conn.Exec(query,email)
		err = errors.New("2")
		return err
	}
	return nil
}

func (pr *Psqlrepo) ConfirmReset(key string) error {
	query := `UPDATE users SET password_reset=1 WHERE confirmation_token=$1`
	_,err := pr.Conn.Exec(query,key)
	if err != nil {
		return err
	}
	return nil
}

func (pr *Psqlrepo) ResetPassword(Id int,password string) error {
	log.Println("Hashed password from repo",password)
	query := `UPDATE users SET password=$1,password_reset=0  WHERE user_id=$2`
	_,err := pr.Conn.Exec(query,password,Id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (pr *Psqlrepo) GetUser(token string) (string,string,error) {
	query := `Select username,password from users WHERE confirmation_token=$1`
	var username string
	var password string
	row := pr.Conn.QueryRow(query,token)
	switch err := row.Scan(&username,&password); err {
	case sql.ErrNoRows:
		return username,password,err
	}
	return username,password,nil
}

func (pr *Psqlrepo) QueFilled(username string,password string) (bool,error) {
	query := `SELECT quefilled FROM users WHERE username=$1`
	var CheckFilled int
	row := pr.Conn.QueryRow(query,username)
	switch err := row.Scan(&CheckFilled); err{
	case sql.ErrNoRows:
		return false,err
	}
	return true, nil
	//var hashedPassword string
	//fmt.Println(hashedPassword)
	//query = `SELECT password FROM users WHERE username=$1`
	//row = pr.Conn.QueryRow(query,username)
	//err := row.Scan(&hashedPassword)
	//isEqual := form.CompareHashToPassword(password,hashedPassword)
	//if !isEqual || err != nil{
	//	return isEqual,err
	//}
	//return true,nil
}
