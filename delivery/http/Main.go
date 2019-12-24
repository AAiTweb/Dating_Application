package main

import (
	"database/sql"
	"github.com/Eyosi-G/Dating_Application/Api"
	repository2 "github.com/Eyosi-G/Dating_Application/Api/repository"
	service2 "github.com/Eyosi-G/Dating_Application/Api/service"
	"github.com/Eyosi-G/Dating_Application/message/repository"
	"github.com/Eyosi-G/Dating_Application/message/service"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
)

var upgrader websocket.Upgrader
var users  = make(map[int]*websocket.Conn)
var templ = template.Must(template.ParseGlob("../../ui/templates/*.html"))

func main() {

	db,err := sql.Open("postgres","postgres://postgres:password@localhost/dating_app?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		panic(err)
	}




	//sampleMessage := entity.Message{-1,1,4,"hello",time.Now(),2}

	//repositoryMessage := repository.RepositoryMessage{db}
	//serviceMessage := service.MessageService{repositoryMessage}
	//socketHandler := Socket.SocketHandler{upgrader,users}


	//api
	//Handler := Api.APIHandler{Db:db}

	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("../../ui/assets"))
	messageReppo := repository.NewRepositoryMessage(db)
	msgService := service.NewMessageService(messageReppo)
	apirepo := repository2.NewApiRepository(db)
	apiservice := service2.NewApiService(apirepo)
	handler := Api.NewApiHandler(msgService,apiservice)

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/",fs))
	router.HandleFunc("/",index)
	router.HandleFunc("/user/{id}/friends",handler.GetFriends)
	router.HandleFunc("/chats/user/{uid}/friends/{fid}",handler.GetMessages)
	router.HandleFunc("/ws",Socket)
	http.ListenAndServe("localhost:8081",router)

	//fmt.Println(time.Now())


	//fmt.Println(serviceMessage.Messages(1,5))


	//userrepoInstance := repository.UserRepositoryInstance{db}
	//userserviceInstance:=service.UserServiceInstance{RepositoryInstance:userrepoInstance}
	//handlerInstance := handler.MainHandler{
	//	Templ:    templ,
	//	Uservice: userserviceInstance,
	//}
	//:= http.NewServeMux()

	//mux.HandleFunc("/login",handler.Login)
	//mux.HandleFunc("/signup",handler.Signup)
	//mux.HandleFunc("/notification",handler.Notification)
	//mux.HandleFunc("/message",handler.Message)
	//mux.HandleFunc("/profile",handler.Profile)
	//mux.HandleFunc("/home",handlerInstance.Home)
	//mux.HandleFunc("/questionnaire",handler.Questionnaire)
	//mux.HandleFunc("/",handlerInstance.Home)
	//
	//http.ListenAndServe("localhost:8082",mux)
	//

}




func index(w http.ResponseWriter, r *http.Request){
	templ.ExecuteTemplate(w,"index.html",nil)
}
func Socket(w http.ResponseWriter, r *http.Request)  {
	upgrader.CheckOrigin  = func(r *http.Request) bool {
		return true;
	}
	conn,err := upgrader.Upgrade(w,r,nil)
	if err!=nil{
		return
	}
	for{
		messageType, message,_ := conn.ReadMessage()
		conn.WriteMessage(messageType,message)
	}


	//sh.Conncetions[int(id)]= conn;
	//sh.readMessage(conn)
}
