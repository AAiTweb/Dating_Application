package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"

	//////////////////////////////////////betse//////////////////////////////////////////
	mainHand "github.com/AAiTweb/Dating_Application/delivery/http/handler"
	QueRep "github.com/AAiTweb/Dating_Application/questionnarie/repository"
	QueSer "github.com/AAiTweb/Dating_Application/questionnarie/service"
	usrProRepo "github.com/AAiTweb/Dating_Application/user_profile/repository"
	usrProServ "github.com/AAiTweb/Dating_Application/user_profile/service"
	//////////////////////////////////////Eyosi//////////////////////////////////////////
	Handler3 "github.com/AAiTweb/Dating_Application/ChatApi/Handler"
	repository3 "github.com/AAiTweb/Dating_Application/ChatApi/repository"
	service3 "github.com/AAiTweb/Dating_Application/ChatApi/service"
	Handler2 "github.com/AAiTweb/Dating_Application/HomeApi/Handler"
	HomeRep "github.com/AAiTweb/Dating_Application/HomeApi/Repository"
	HomeSer "github.com/AAiTweb/Dating_Application/HomeApi/Service"
	Socket "github.com/AAiTweb/Dating_Application/Socket"
	handler2 "github.com/AAiTweb/Dating_Application/delivery/http/handler/EyosiHand"
	MeseRep "github.com/AAiTweb/Dating_Application/message/repository"
	MesServe "github.com/AAiTweb/Dating_Application/message/service"
	repository2 "github.com/AAiTweb/Dating_Application/relationship/repository"
	service2 "github.com/AAiTweb/Dating_Application/relationship/service"
	session "github.com/AAiTweb/Dating_Application/session"

	//////////////////////////////////////Bini//////////////////////////////////////////
	//newHand "github.com/biniyam112/TheDatingApp/Dating_Application/delivery/http/handler"
	notif_repo "github.com/AAiTweb/Dating_Application/notification/repositry"
	notif_serv "github.com/AAiTweb/Dating_Application/notification/service"
	login_repo "github.com/AAiTweb/Dating_Application/user/repository"
	login_serv "github.com/AAiTweb/Dating_Application/user/service"
)

func main() {
	dbConne, err := sql.Open("postgres", "postgres://admin:kali@localhost/dating_app1?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer dbConne.Close()

	if err := dbConne.Ping(); err != nil {
		panic(err)
	}

	log.Println("database connected")

	var upgrader websocket.Upgrader
	var users = make(map[int]*websocket.Conn)
	//tmpl := template.Must(template.ParseGlob("../../../ui/template/*.html"))
	tmpl := template.Must(template.ParseGlob("github.com/biniyam112/TheDatingApp/Dating_Application/ui/template/*.html"))

	fs := http.FileServer(http.Dir("github.com/biniyam112/TheDatingApp/Dating_Application/ui/assets"))
	mux := mux.NewRouter()
	mux.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", fs))
	//////////////////////////////////////betse//////////////////////////////////////////
	//Questionnaire
	questionRespo := QueRep.NewQuestionnarieRepoImpl(dbConne)
	questionService := QueSer.NewQuestionnaireServiceImpl(questionRespo)
	questionnarieHandler := mainHand.NewUserQuestionnarieHandler(tmpl, questionService)

	userProfileRepo := usrProRepo.NewUserProfileRepoImpl(dbConne)
	userProfileServ := usrProServ.NewUserProfileServiceImpl(userProfileRepo)
	userProfHandler := mainHand.NewUserProfileHandler(tmpl, userProfileServ)

	mux.HandleFunc("/user/questionnarie/answers/{user_id}/{index}", questionnarieHandler.PostAnswers)
	mux.HandleFunc("/user/profile", userProfHandler.GetUser)
	mux.HandleFunc("/user/questionnarie", questionnarieHandler.MainQuestionnarie)
	mux.HandleFunc("/user/questionnarie/questions", questionnarieHandler.Questionnaire)
	mux.HandleFunc("/user/profile/addUser", userProfHandler.PostUser)
	mux.HandleFunc("/user/profile/update", userProfHandler.PutUser)
	//////////////////////////////////////betse//////////////////////////////////////////

	//fs = http.FileServer(http.Dir("../../ui/assets"))

	//Chat Api
	MessageRepository := MeseRep.NewRepositoryMessage(dbConne)
	MessageService := MesServe.NewMessageService(MessageRepository)
	ChatApiRepo := repository3.NewApiRepository(dbConne)
	ChatApiService := service3.NewApiService(ChatApiRepo)
	handler := Handler3.NewApiHandler(MessageService, ChatApiService)
	socketHandler := Socket.NewSocketHandler(upgrader, users, MessageService)

	//Home Api
	relationshipRepo := repository2.NewRelationshipRepository(dbConne)
	relationshipService := service2.NewRelationshipService(relationshipRepo)
	HomeApiRepository := HomeRep.NewHomeApiRepository(dbConne)
	HomeApiService := HomeSer.NewHomeApiService(HomeApiRepository)
	HomeApiHandler := Handler2.NewHomeApiHandler(HomeApiRepository, HomeApiService, relationshipService)

	//Main Handler
	mainHandler := handler2.NewMainHandler(tmpl, dbConne)

	mux.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))
	mux.Handle("/chat", session.IsAuthenticated(http.HandlerFunc(mainHandler.Index)))
	mux.Handle("/home", session.IsAuthenticated(http.HandlerFunc(mainHandler.Home)))

	//Chat Page EndPoints
	mux.HandleFunc("/user/{id}/friends", handler.GetFriends)
	mux.Handle("/chats/user/{uid}/friends/{fid}", session.IsAuthenticated(http.HandlerFunc(handler.GetMessages)))
	mux.Handle("/user/{id}/updatelogin", session.IsAuthenticated(http.HandlerFunc(handler.UpdateLoginDetails)))
	mux.Handle("/ws", session.IsAuthenticated(http.HandlerFunc(socketHandler.Socket)))

	// Home Page EndPoint
	mux.Handle("/matches/user/{id}", session.IsAuthenticated(http.HandlerFunc(HomeApiHandler.GetUsersMatched)))
	mux.Handle("/matches/sendrequest", session.IsAuthenticated(http.HandlerFunc(HomeApiHandler.SendRequest)))
	mux.Handle("/matches/acceptrequest", session.IsAuthenticated(http.HandlerFunc(HomeApiHandler.AcceptRequest)))
	log.Println("Server Started Listening .....")

	//////////////////////////////////////Bini//////////////////////////////////////////
	UserrepoInstance := login_repo.NewUserRepo(dbConne)
	userserviceInstance := login_serv.NewUserServe(UserrepoInstance)
	UserhandlerInstance := mainHand.NewUserHandler(userserviceInstance, tmpl)
	NotifrepoInstance := notif_repo.NewUserRepo(dbConne)
	NotifServiceInstance := notif_serv.NewNotifServe(NotifrepoInstance)
	NotifhandlerInstance := mainHand.NewMainHandler(NotifServiceInstance, tmpl)

	//Login server
	fmt.Println("connection established!")
	mux.HandleFunc("/login", UserhandlerInstance.Login)
	mux.HandleFunc("/signup", UserhandlerInstance.Signup)
	mux.HandleFunc("/signup/validate", UserhandlerInstance.ValidateSignup)
	mux.HandleFunc("/validate", UserhandlerInstance.Validatelogin)
	mux.HandleFunc("/notification", NotifhandlerInstance.SeeNotification)
	mux.HandleFunc("/accept", NotifhandlerInstance.AcceptNotification)
	mux.HandleFunc("/reject", NotifhandlerInstance.RejectNotification)

	serv := http.Server{
		Addr:    ":8081",
		Handler: mux,
	}
	log.Println("litsening on :8081")
	serv.ListenAndServe()

}
