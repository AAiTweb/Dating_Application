package main

import (
	"database/sql"
	"fmt"
	"github.com/biniyam112/Dating_Application/delivery/http/handler"
	notif_repo "github.com/biniyam112/Dating_Application/notification/repositry"
	notif_serv "github.com/biniyam112/Dating_Application/notification/service"
	login_repo "github.com/biniyam112/Dating_Application/user/repository"
	login_serv "github.com/biniyam112/Dating_Application/user/service"
	"html/template"
	"log"
	"net/http"

	//"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
)

var templ = template.Must(template.ParseGlob("/root/go_projects/src/github.com/biniyam112/Dating_Application/ui/assets/loginpage/*.html"))

func main() {
	connvalue := "postgres://admin:kali@localhost/dating_app?sslmode=disable"
	db, err := sql.Open("postgres", connvalue)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()

	UserrepoInstance := login_repo.NewUserRepo(db)
	userserviceInstance := login_serv.NewUserServe(UserrepoInstance)
	UserhandlerInstance := handler.NewUserHandler(userserviceInstance, templ)
	NotifrepoInstance := notif_repo.NewUserRepo(db)
	NotifServiceInstance := notif_serv.NewNotifServe(NotifrepoInstance)
	NotifhandlerInstance := handler.NewMainHandler(NotifServiceInstance, templ)

	//api
	//Handler := Api.APIHandler{Db:db}
	//router := mux.NewRouter()
	//fs := http.FileServer(http.Dir("/root/go_projects/src/github.com/biniyam112/log_in/assets"))
	//router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/",fs))
	//router.HandleFunc("/login",Handler.Login)
	//router.HandleFunc("/signup",Handler.Signup)
	//router.HandleFunc("/validate",Handler.Validatelogin)
	//router.HandleFunc("/user/{id}/friends",Handler.Friends)
	//router.HandleFunc("/chats/user/{uid}/friends/{fid}",Handler.Messages)
	////router.HandleFunc("/ws",socketHandler)
	//fmt.Println("connection established!")
	//http.ListenAndServe("localhost:8081",router)

	//Login server
	fmt.Println("connection established!")
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("/root/go_projects/src/github.com/biniyam112/Dating_Application/ui/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/login", UserhandlerInstance.Login)
	mux.HandleFunc("/signup", UserhandlerInstance.Signup)
	mux.HandleFunc("/signup/validate", UserhandlerInstance.ValidateSignup)
	mux.HandleFunc("/validate", UserhandlerInstance.Validatelogin)
	mux.HandleFunc("/notification", NotifhandlerInstance.SeeNotification)
	mux.HandleFunc("/accept", NotifhandlerInstance.AcceptNotification)
	mux.HandleFunc("/reject", NotifhandlerInstance.RejectNotification)
	fmt.Println("starting servier...")
	err = http.ListenAndServe(":8088", mux)
	if err != nil {
		log.Fatal(err)
	}

}
