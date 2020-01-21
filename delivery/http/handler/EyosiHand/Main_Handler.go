package EyosiHand

import (
	"database/sql"
	"errors"
	"github.com/AAiTweb/Dating_Application/session"
	"html/template"
	"log"
	"net/http"
)

type MainHandler struct {
	Templ *template.Template
	Db *sql.DB
}

func NewMainHandler(t *template.Template, db *sql.DB)MainHandler{
	return MainHandler{t,db}
}


func (m MainHandler)Index(w http.ResponseWriter, r *http.Request){
	m.Templ.ExecuteTemplate(w,"chatpage.html",nil)

}

func (m MainHandler)Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet{
		m.Templ.ExecuteTemplate(w,"login.html",nil)
	}else{
		uname := r.FormValue("username")
		pword := r.FormValue("password")
		row := m.Db.QueryRow(`select t1.user_id,t1.username,t3.picture_path
		from (select user_id, username from users where username=$1 and password=$2) as t1
			inner join (select picture_owner_id,picture_path from user_profile inner join gallery on picture_id=profile_picture) t3 on
				t3.picture_owner_id=t1.user_id;`,uname,pword)
		usr := struct {
			Id int
			UserName,
			ProfilePicture string
		}{}
		err := row.Scan(&usr.Id,&usr.UserName,&usr.ProfilePicture)
		log.Println(usr)


		if err!=nil && err==sql.ErrNoRows {
			errors.New("User Doesn't Exist")
		}else{
			tokenString,err := session.Generate(usr.Id,usr.UserName,usr.ProfilePicture)
			if err!=nil{
				return
			}
			http.SetCookie(w, &http.Cookie{
				Name:    "token",
				Value:   tokenString,
			})
			http.Redirect(w,r,"/chat",http.StatusSeeOther)
		}
	}
}

func (m MainHandler)Home(w http.ResponseWriter, r *http.Request){
	m.Templ.ExecuteTemplate(w,"dashboard.html",nil)
}