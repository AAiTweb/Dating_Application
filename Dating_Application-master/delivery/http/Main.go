package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/betse/Dating_Application-master/delivery/http/handler"
	// "github.com/betse/Dating_Application-master/entity"
	"github.com/betse/Dating_Application-master/questionnarie/repository"
	"github.com/betse/Dating_Application-master/questionnarie/service"
)

func main() {
	dbConne, err := sql.Open("postgres", "postgres://betse:26300001@localhost/dating_app")
	if err != nil {
		panic(err)
	}
	defer dbConne.Close()

	if err := dbConne.Ping(); err != nil {
		panic(err)
	}
	// query2 := "SELECT * FROM dating_app.questionnaires; "
	// rows, err := dbConne.Query(query2)
	// defer rows.Close()

	// questions := []test{}
	// for rows.Next() {
	// 	question := test{}
	// 	err := rows.Scan(&question.Id, &question.Question1, &question.Question2)

	// 	if err != nil {
	// 		log.Println("err")
	// 	}
	// 	questions = append(questions, question)
	// }
	// log.Println(len(questions), "LENGTH")

	log.Println("database connected")

	tmpl := template.Must(template.ParseGlob("../../ui/template/*"))
	//Questionnaire
	questionRespo := repository.NewQuestionnarieRepoImpl(dbConne)
	questionService := service.NewQuestionnaireServiceImpl(questionRespo)
	questionnarieHandler := handler.NewUserQuestionnarieHandler(tmpl, questionService)

	// questionnarieHandler := handler.NewUserQuestionnarieHandler(tmpl)

	fs := http.FileServer(http.Dir("../../ui/assets"))
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/user/questionnarie", questionnarieHandler.MainQuestionnarie)
	mux.HandleFunc("/user/questionnarie/questions", questionnarieHandler.Questionnaire)
	serv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Println("litsening on :8080")
	serv.ListenAndServe()

}
