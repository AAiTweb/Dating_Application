package handler

import (
	"fmt"
	"github.com/AAiTweb/Dating_Application/entity"
	"github.com/AAiTweb/Dating_Application/form"
	"github.com/AAiTweb/Dating_Application/session"
	"github.com/AAiTweb/Dating_Application/user/service"
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

//var logintemp = template.Must(template.ParseFiles("/root/go_projects/src/github.com/biniyam112/Dating_Application/ui/assets/loginpage/logincopy.html"))
//var hometemp = template.Must(template.ParseFiles("/root/go_projects/src/github.com/biniyam112/Dating_Application/ui/assets/dashboard/homepage.html"))

func (mh UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Print("this is bini's MF")
	mh.Templ.ExecuteTemplate(w,"logincopy.html", nil)
	w.Write([]byte("Login"))
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
			mh.Templ.ExecuteTemplate(w, "logincopy.html","Passwords doesn't match, re-enter password")
			return
		}
		if form.ValidatePassword(Fpassword) != "" {
			mh.Templ.ExecuteTemplate(w,"logincopy.html", form.ValidatePassword(Fpassword))
			return
		}

		if len(Fusername) < 5 {
			println("short username")
			mh.Templ.ExecuteTemplate(w, "logincopy.html","Username must be longer than 5 characters")
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
			mh.Templ.ExecuteTemplate(w,"logincopy.html", "Unstable conn or already used username or email")
		} else {
			mh.Templ.ExecuteTemplate(w,"logincopy.html" ,nil)
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
		userId,UserName,profilePic, err := handlerserivce.CheckLogin(user)
		//log.Println(UserName)
		if err != nil {
			mh.Templ.ExecuteTemplate(w,"logincopy.html", "incorrect username or password, try again")
			fmt.Println("incorrect username or password")
		} else {
			fmt.Println("input: " + password + " saved: " + password)
				println("Correct username and password")
				//r.URL, _ = r.URL.Parse("/home")
				//hometemp.Execute(w, nil)
					tokenString,err := session.Generate(userId,UserName,profilePic)
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
