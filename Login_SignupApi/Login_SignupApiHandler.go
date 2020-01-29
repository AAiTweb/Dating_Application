package Login_SignupApi

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/AAiTweb/Dating_Application/entity"
	"io/ioutil"
	"net/http"
	"time"
)

type APIHandler struct {
	Db *sql.DB
}

func (ap *APIHandler) Validatelogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("Lusername")
		password := r.FormValue("Lpassword")
		var email, conftoken, saved_password string
		querystat := `select email,confirmation_token,password from users where username=$1`
		User_row := ap.Db.QueryRow(querystat, username)
		switch err := User_row.Scan(&email, &conftoken, &saved_password); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
			return
		case nil:
			fmt.Println("info retrived")
		default:
			panic(err)
		}
		if password != saved_password {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(w).Encode("Incorrect password")
			w.Write([]byte("signup"))
			return
		}
		user := entity.User{
			Id:                0,
			UserName:          username,
			Password:          password,
			Email:             email,
			ConfirmationToken: conftoken,
		}
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &user)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(&user)
		w.Write([]byte("signup"))

	}
}

func (ap *APIHandler) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		Fusername := r.FormValue("f_username")
		Femail := r.FormValue("f_email")
		Fpassword := r.FormValue("f_password")
		time := time.Now().String()
		confToken := time + Fusername
		fmt.Println(Fusername, confToken)
		fmt.Println(Femail, Fpassword)
		user := entity.User{
			Id:                0,
			UserName:          Fusername,
			Password:          Fpassword,
			Email:             Femail,
			ConfirmationToken: confToken,
		}
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &user)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(&user)
		w.Write([]byte("Validate login"))
	}
}
