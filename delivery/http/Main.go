package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
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
	"github.com/AAiTweb/Dating_Application/Matches/MatchHandler"
	matchRepo "github.com/AAiTweb/Dating_Application/Matches/repository"
	matchServ "github.com/AAiTweb/Dating_Application/Matches/service"
	Socket "github.com/AAiTweb/Dating_Application/Socket"
	handler2 "github.com/AAiTweb/Dating_Application/delivery/http/handler/EyosiHand"
	MeseRep "github.com/AAiTweb/Dating_Application/message/repository"
	MesServe "github.com/AAiTweb/Dating_Application/message/service"
	repository2 "github.com/AAiTweb/Dating_Application/relationship/repository"
	service2 "github.com/AAiTweb/Dating_Application/relationship/service"
	//////////////////////////////////////Bini//////////////////////////////////////////
	//newHand "github.com/biniyam112/TheDatingApp/Dating_Application/delivery/http/handler"
	notif_repo "github.com/AAiTweb/Dating_Application/notification/repositry"
	notif_serv "github.com/AAiTweb/Dating_Application/notification/service"
	login_repo "github.com/AAiTweb/Dating_Application/user/repository"
	login_serv "github.com/AAiTweb/Dating_Application/user/service"
)

var tmpl = template.Must(template.ParseGlob("../../ui/template/*.html"))

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

	//tmpl := template.Must(template.ParseGlob("github.com/AAiTweb/Dating_Application/ui/template/*.html"))

	//fs := http.FileServer(http.Dir("../../ui/assets"))
	//mux := mux.NewRouter()
	mux := httprouter.New()
	//mux.Han
	//mux.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", fs))
	//http.Handle("/assets/",http.StripPrefix("/assets/",fs))
	mux.ServeFiles("/assets/*filepath", http.Dir("../../ui/assets"))
	//Questionnaire
	questionRespo := QueRep.NewQuestionnarieRepoImpl(dbConne)
	questionService := QueSer.NewQuestionnaireServiceImpl(questionRespo)
	questionnarieHandler := mainHand.NewUserQuestionnarieHandler(tmpl, questionService)

	userProfileRepo := usrProRepo.NewUserProfileRepoImpl(dbConne)
	userProfileServ := usrProServ.NewUserProfileServiceImpl(userProfileRepo)
	userProfHandler := mainHand.NewUserProfileHandler(tmpl, userProfileServ)

	//mux.Han
	mux.POST("/user/questionnarie/answers/:user_id/:index", questionnarieHandler.PostAnswers)
	mux.HandlerFunc("GET", "/user/profile", userProfHandler.GetUser)
	mux.HandlerFunc("GET", "/user/questionnarie", questionnarieHandler.MainQuestionnarie)
	mux.HandlerFunc("GET", "/user/questionnarie/questions", questionnarieHandler.Questionnaire)
	mux.HandlerFunc("POST", "/user/addUser", userProfHandler.PostUser)
	mux.HandlerFunc("POST", "/user/update", userProfHandler.PutUser)
	mux.HandlerFunc("GET", "/user/FriendProfile", userProfHandler.GetFriendProfile)
	mux.HandlerFunc("GET", "/user/delete", userProfHandler.DeleteProfile)
	//////////////////////////////////////betse//////////////////////////////////////////

	//fs = http.FileServer(http.Dir("../../ui/assets"))
	matchRepo := matchRepo.NewMatchRepository(dbConne)
	matchServ := matchServ.NewMatchService(matchRepo)
	matchHander := MatchHandler.NewMatchHandler(matchServ)
	mux.HandlerFunc("GET", "/match", matchHander.DoMatching)

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
	mainHandler := handler2.NewMainHandler(tmpl)

	//mux.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))
	mux.HandlerFunc("GET", "/chat", mainHandler.Index)
	mux.HandlerFunc("GET", "/home", mainHandler.Home)
	mux.HandlerFunc("GET", "/search/user/", mainHandler.Search)
	mux.HandlerFunc("GET", "/preload", mainHandler.Preload)

	//Chat Page EndPoints
	mux.GET("/user/friends/:id", handler.GetFriends)
	mux.GET("/chats/user/:uid/friends/:fid", handler.GetMessages)
	mux.GET("/user/updatelogin/:id", handler.UpdateLoginDetails)
	mux.HandlerFunc("GET", "/ws", socketHandler.Socket)

	// Home Page EndPoint
	mux.GET("/matches/user/:id", HomeApiHandler.GetUsersMatched)
	mux.HandlerFunc("GET", "/matches/sendrequest", HomeApiHandler.SendRequest)
	mux.HandlerFunc("POST", "/matches/acceptrequest", HomeApiHandler.AcceptRequest)
	mux.GET("/Home/Search/:uname", HomeApiHandler.Search)
	// match

	//////////////////////////////////////Bini//////////////////////////////////////////
	UserrepoInstance := login_repo.NewUserRepo(dbConne)
	userserviceInstance := login_serv.NewUserServe(UserrepoInstance)
	UserhandlerInstance := mainHand.NewUserHandler(userserviceInstance, tmpl)
	NotifrepoInstance := notif_repo.NewUserRepo(dbConne)
	NotifServiceInstance := notif_serv.NewNotifServe(NotifrepoInstance)
	NotifhandlerInstance := mainHand.NewMainHandler(NotifServiceInstance, tmpl)

	//Login server
	fmt.Println("connection established!")
	mux.HandlerFunc("GET", "/", UserhandlerInstance.Login)
	mux.HandlerFunc("GET", "/login", UserhandlerInstance.Login)
	mux.HandlerFunc("POST", "/signup", UserhandlerInstance.Signup)
	mux.HandlerFunc("GET", "/signup/validate", UserhandlerInstance.ValidateSignup)
	mux.HandlerFunc("POST", "/validate", UserhandlerInstance.Validatelogin)
	mux.HandlerFunc("GET", "/notification", NotifhandlerInstance.SeeNotification)
	mux.HandlerFunc("GET", "/accept", NotifhandlerInstance.AcceptNotification)
	mux.HandlerFunc("GET", "/reject", NotifhandlerInstance.RejectNotification)
	mux.HandlerFunc("GET", "/Logout", UserhandlerInstance.Logout)

	mux.HandlerFunc("GET", "/confirmemailpage", UserhandlerInstance.ForgotPassword)
	mux.HandlerFunc("POST", "/forgotpassword", UserhandlerInstance.ConfirmEmail)
	mux.HandlerFunc("GET", "/confirmreset", UserhandlerInstance.ConfirmReset)
	mux.HandlerFunc("GET", "/resetpassword", UserhandlerInstance.Passwordreset)
	mux.HandlerFunc("POST", "/passwordreset", UserhandlerInstance.ResetPassword)

	serv := http.Server{
		Addr:    "localhost:8081",
		Handler: mux,
	}
	log.Println("litsening on :8081")
	serv.ListenAndServe()

}
