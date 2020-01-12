package handler

import (
	"fmt"
	"github.com/biniyam112/Dating_Application/entity"
	"github.com/biniyam112/Dating_Application/form"
	"github.com/biniyam112/Dating_Application/user/service"
	"github.com/biniyam112/log_in/password"
	"html/template"
	"log"
	"net/http"
	"time"
)

type MainHandler struct {
	Templ    template.Template
	Uservice service.UserService
}

func NewMainHandler(servc *service.UserService, temp *template.Template) *MainHandler {
	return &MainHandler{
		Templ:    *temp,
		Uservice: *servc,
	}
}

var logintemp = template.Must(template.ParseFiles("/root/go_projects/src/github.com/biniyam112/log_in/login_copy.php"))
var hometemp = template.Must(template.ParseFiles("/root/go_projects/src/github.com/biniyam112/log_in/homepage.html"))

func (mh MainHandler) Login(w http.ResponseWriter, r *http.Request) {
	logintemp.Execute(w, nil)
	//w.Write([]byte("Login"))
}
func (mh MainHandler) Signup(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		Fusername := r.FormValue("f_username")
		Femail := r.FormValue("f_email")
		Fpassword := r.FormValue("f_password")
		Repassword := r.FormValue("confirm_f_password")
		time := time.Now().String()
		confToken := time + Fusername

		if Fpassword != Repassword {
			logintemp.Execute(w, "Passwords doesn't match, re-enter password")
			return
		}
		if password.ValidatePassword(Fpassword) != "" {
			logintemp.Execute(w, form.ValidatePassword(Fpassword))
			return
		}

		if len(Fusername) < 5 {
			println("short username")
			logintemp.Execute(w, "Username must be longer than 5 characters")
			return
		}
		fmt.Println(Fusername)
		fmt.Println(Femail, Fpassword)
		user := entity.User{
			Id:                0,
			UserName:          Fusername,
			Password:          Fpassword,
			Email:             Femail,
			ConfirmationToken: confToken,
		}
		handlerserivce := mh.Uservice
		err := handlerserivce.RegisterUser(user)
		if err != nil {
			fmt.Println(err)
			logintemp.Execute(w, "already used username or email")
			fmt.Println("already used username or email")
		} else {
			hometemp.Execute(w, nil)
		}

	}
}

func (mh MainHandler) Validatelogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("Lusername")
		password := r.FormValue("Lpassword")
		fmt.Println(username, password)
		user := entity.User{
			Id:                0,
			UserName:          username,
			Password:          password,
			Email:             "",
			ConfirmationToken: "",
		}
		handlerserivce := mh.Uservice
		Rpassword, err := handlerserivce.CheckLogin(user)
		fmt.Println(Rpassword)
		if err != nil {
			logintemp.Execute(w, "Invalid username or password, try again")
			fmt.Println("incorrect username or password")
		} else{
			fmt.Println("input: "+string(password)+" saved: "+string(Rpassword))
			if string(password) == string(Rpassword) {
				println("Correct username and password")
				//r.URL, _ = r.URL.Parse("/home")
				hometemp.Execute(w,nil)
				http.Redirect(w, r, "/root/go_projects/src/github.com/biniyam112/Dating_Application/ui/assets/dashboard/homepage.html", 301)
				fmt.Println(r.URL.String())
			} else {
				logintemp.Execute(w, "Invalid username or password, try again")
				fmt.Println("incorrect username or password")
			}
		}
	}
}
func (mh MainHandler) ValidateSignup(w http.ResponseWriter, r *http.Request) {
	vkey := r.URL.Query().Get("conftok")
	if !(len(vkey) > 0) {
		log.Println("verification key not found")
	}
	user := entity.User{
		Id:                0,
		UserName:          "",
		Password:          "",User validation successfuUser validation successfu
		Email:             "",
		ConfirmationToken: vkey,
	}
	handlerserivce := mh.Uservice
	handlerserivce.ValidateToken(user)
	w.Write([]byte("Validate sign up"))
}

func Notification(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Notification"))
}
func Message(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Message"))
}
func Profile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Profile"))
}
func (h *MainHandler) Home(w http.ResponseWriter, r *http.Request) {
	//h.Templ.ExecuteTemplate(w, "dashboard.html", nil)
}
func Questionnaire(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Questionnaire"))
}