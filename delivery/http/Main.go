package main

import (
	"database/sql"
	Handler3 "github.com/Eyosi-G/Dating_Application/ChatApi/Handler"
	repository3 "github.com/Eyosi-G/Dating_Application/ChatApi/repository"
	service3 "github.com/Eyosi-G/Dating_Application/ChatApi/service"
	Handler2 "github.com/Eyosi-G/Dating_Application/HomeApi/Handler"
	"github.com/Eyosi-G/Dating_Application/HomeApi/Repository"
	"github.com/Eyosi-G/Dating_Application/HomeApi/Service"
	"github.com/Eyosi-G/Dating_Application/Socket"
	handler2 "github.com/Eyosi-G/Dating_Application/delivery/http/handler"
	"github.com/Eyosi-G/Dating_Application/message/repository"
	"github.com/Eyosi-G/Dating_Application/message/service"
	repository2 "github.com/Eyosi-G/Dating_Application/relationship/repository"
	service2 "github.com/Eyosi-G/Dating_Application/relationship/service"
	"github.com/Eyosi-G/Dating_Application/session"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
)

var upgrader websocket.Upgrader
var users  = make(map[int]*websocket.Conn)
var templ = template.Must(template.ParseGlob("../../ui/templates/*.html"))

var db,err = sql.Open("postgres","postgres://postgres:password@localhost/dating_app?sslmode=disable")




func main() {

	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("../../ui/assets"))

	//Chat Api
	MessageRepository :=repository.NewRepositoryMessage(db)
	MessageService := service.NewMessageService(MessageRepository)
	ChatApiRepo := repository3.NewApiRepository(db)
	ChatApiService := service3.NewApiService(ChatApiRepo)
	handler := Handler3.NewApiHandler(MessageService,ChatApiService)
	socketHandler := Socket.NewSocketHandler(upgrader,users,MessageService)


	//Home Api
	relationshipRepo := repository2.NewRelationshipRepository(db)
	relationshipService := service2.NewRelationshipService(relationshipRepo)
	HomeApiRepository := Repository.NewHomeApiRepository(db)
	HomeApiService := Service.NewHomeApiService(HomeApiRepository)
	HomeApiHandler :=Handler2.NewHomeApiHandler(HomeApiRepository,HomeApiService,relationshipService)

	//Main Handler
	mainHandler := handler2.NewMainHandler(templ,db)

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/",fs))
	router.Handle("/chat",session.IsAuthenticated(http.HandlerFunc(mainHandler.Index)))
	router.HandleFunc("/login",mainHandler.Login)
	router.Handle("/home",session.IsAuthenticated(http.HandlerFunc(mainHandler.Home)))

	//Chat Page EndPoints
	router.HandleFunc("/user/{id}/friends",handler.GetFriends)
	router.Handle("/chats/user/{uid}/friends/{fid}",session.IsAuthenticated(http.HandlerFunc(handler.GetMessages)))
	router.Handle("/user/{id}/updatelogin",session.IsAuthenticated(http.HandlerFunc(handler.UpdateLoginDetails)))
	router.Handle("/ws",session.IsAuthenticated(http.HandlerFunc(socketHandler.Socket)))

    // Home Page EndPoint
	router.Handle("/matches/user/{id}",session.IsAuthenticated(http.HandlerFunc(HomeApiHandler.GetUsersMatched)))
	router.Handle("/matches/sendrequest",session.IsAuthenticated(http.HandlerFunc(HomeApiHandler.SendRequest)))
	router.Handle("/matches/acceptrequest",session.IsAuthenticated(http.HandlerFunc(HomeApiHandler.AcceptRequest)))
	log.Println("Server Started Listening .....")
	http.ListenAndServe("localhost:8081",router)

}
