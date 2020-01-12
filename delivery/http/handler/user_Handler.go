package handler

import (
	"fmt"
	"github.com/biniyam112/Dating_Application/entity"
	"github.com/biniyam112/Dating_Application/form"
	"github.com/biniyam112/Dating_Application/user/service"
	"html/template"
	"log"
	"net/http"
	"time"
)

type UserHandler struct {
	Templ    template.Template
	Uservice service.UserService
}

func NewUserHandler(servc *service.UserService, temp *template.Template) *UserHandler {
	return &UserHandler{
		Templ:    *temp,
		Uservice: *servc,
	}
}

var logintemp = template.Must(template.ParseFiles("/root/go_projects/src/github.com/biniyam112/Dating_Application/ui/assets/loginpage/logincopy.html"))
var hometemp = template.Must(template.ParseFiles("/root/go_projects/src/github.com/biniyam112/Dating_Application/ui/assets/dashboard/homepage.html"))

func (mh UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	logintemp.Execute(w, nil)
	//w.Write([]byte("Login"))
}
func (mh UserHandler) Signup(w http.ResponseWriter, r *http.Request) {

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
		if form.ValidatePassword(Fpassword) != "" {
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
			fmt.Println("the error is ", err)
			logintemp.Execute(w, "Unstable conn or already used username or email")
		} else {
			hometemp.Execute(w, nil)
		}

	}
}

func (mh UserHandler) Validatelogin(w http.ResponseWriter, r *http.Request) {
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
			logintemp.Execute(w, "incorrect username or password, try again")
			fmt.Println("incorrect username or password")
		} else {
			fmt.Println("input: " + password + " saved: " + Rpassword)
			if password == Rpassword {
				println("Correct username and password")
				//r.URL, _ = r.URL.Parse("/home")
				hometemp.Execute(w, nil)
				//http.Redirect(w, r, "/root/go_projects/src/github.com/biniyam112/Dating_Application/ui/assets/dashboard/homepage.html", 301)
				fmt.Println(r.URL.String())
			} else {
				logintemp.Execute(w, "user doesn't exist, try again")
				fmt.Println("incorrect username or password")
			}
		}
	}
}
func (mh UserHandler) ValidateSignup(w http.ResponseWriter, r *http.Request) {
	vkey := r.URL.Query().Get("conftok")
	if !(len(vkey) > 0) {
		log.Println("verification key not found")
	}
	user := entity.User{
		Id:                0,
		UserName:          "",
		Password:          "",
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
func (h *UserHandler) Home(w http.ResponseWriter, r *http.Request) {
	//h.Templ.ExecuteTemplate(w, "dashboard.html", nil)
}
func Questionnaire(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Questionnaire"))
}
