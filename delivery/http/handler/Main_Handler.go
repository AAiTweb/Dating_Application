package handler

import (
	"github.com/Eyosi-G/Dating_Application/user"
	"html/template"
	"net/http"
)

type MainHandler struct {
	Templ *template.Template
	Uservice user.UserService

}

func Login(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Login"))
}

func Signup(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Sign up"))
}
func Notification(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Notification"))
}
func Message(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Message"))
}
func Profile(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Profile"))
}
func(h *MainHandler) Home(w http.ResponseWriter, r *http.Request){
	h.Templ.ExecuteTemplate(w,"dashboard.html",nil)
}
func Questionnaire(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Questionnaire"))
}