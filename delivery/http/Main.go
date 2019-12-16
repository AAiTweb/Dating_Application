package main

import (
	"database/sql"
	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
	"net/http"
)

var upgrader websocket.Upgrader
var users  = make(map[int]*websocket.Conn)

func main() {
	//templ := template.Must(template.ParseGlob("../../ui/templates/*.html"))
	db,err := sql.Open("postgres","postgres://postgres:password@localhost/project?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		panic(err)
	}



	upgrader.CheckOrigin = func(r *http.Request) bool{
		return true
	}

	//sampleMessage := entity.Message{-1,1,4,"hello",time.Now(),2}

	//repositoryMessage := repository.RepositoryMessage{db}
	//serviceMessage := service.MessageService{repositoryMessage}
	//socketHandler := Socket.SocketHandler{upgrader,users,serviceMessage}
	//

	//api
	//Handler := APIHandler{db}
	//http.HandleFunc("/friends",Handler.friends)
	//http.ListenAndServe("localhost:8081",nil)
	//


	//fmt.Println(serviceMessage.Messages(1,5))


	//userrepoInstance := repository.UserRepositoryInstance{db}
	//userserviceInstance:=service.UserServiceInstance{RepositoryInstance:userrepoInstance}
	//handlerInstance := handler.MainHandler{
	//	Templ:    templ,
	//	Uservice: userserviceInstance,
	//}
	// := http.NewServeMux()
	//fs := http.FileServer(http.Dir("../../ui/assets"))
	//mux.Handle("/assets/",http.StripPrefix("/assets/",fs))
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




