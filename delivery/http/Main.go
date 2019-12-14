package main

import (
	"database/sql"
	"fmt"
	"github.com/Eyosi-G/Dating_Application/message/repository"
	"github.com/Eyosi-G/Dating_Application/message/service"
	_ "github.com/lib/pq"
)

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

	//sampleMessage := entity.Message{-1,1,4,"hello",time.Now(),2}
	repositoryMessage := repository.RepositoryMessage{db}
	serviceMessage := service.MessageService{repositoryMessage}

	fmt.Println(serviceMessage.Messages(1,5))


	//userrepoInstance := repository.UserRepositoryInstance{db}
	//userserviceInstance:=service.UserServiceInstance{RepositoryInstance:userrepoInstance}
	//handlerInstance := handler.MainHandler{
	//	Templ:    templ,
	//	Uservice: userserviceInstance,
	//}
	//mux := http.NewServeMux()
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