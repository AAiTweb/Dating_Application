package handler

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/AAiTweb/Dating_Application/entity"
	"github.com/AAiTweb/Dating_Application/questionnarie"
	"github.com/gorilla/mux"
)

type UserQuestionnarieHandler struct {
	tmpl              *template.Template
	questionnaireServ questionnarie.QuestionnarieService
}

func NewUserQuestionnarieHandler(T *template.Template, qs questionnarie.QuestionnarieRespository) *UserQuestionnarieHandler {
	return &UserQuestionnarieHandler{tmpl: T, questionnaireServ: qs}
}

func (userQ *UserQuestionnarieHandler) MainQuestionnarie(w http.ResponseWriter, r *http.Request) {

	userQ.tmpl.ExecuteTemplate(w, "main_questionnarie.layout", nil)
}

func (userQ *UserQuestionnarieHandler) PostAnswers(w http.ResponseWriter, r *http.Request) {
	// userChoice := &entity.UserChoice{}

	user := &entity.UserChoice{}
	length := r.ContentLength
	body := make([]byte, length)
	r.Body.Read(body)

	user.UserId, _ = strconv.Atoi(mux.Vars(r)["user_id"])
	if r.Header.Get("Content-type") != " " {

		err := json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	}
	user, err := userQ.questionnaireServ.PostAnswers(user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	log.Println(user)

	// w.Header().Set("content-type", "application/json")
	// log.Println("hello from ajax")
	// log.Printf("question %d", user.QuestionId)
	// log.Println("wish %d", user.QuestionId)
	// log.Println("ownAnswerId %d", user.OwnAnswerId)

	// log.Println("user id %d", userChoice.UserId)
}

func (userQ *UserQuestionnarieHandler) Questionnaire(w http.ResponseWriter, r *http.Request) {
	questions, answers, err := userQ.questionnaireServ.Questions()
	if err != nil {
		w.Header().Set("content-type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	jsonDataS := []entity.JsonData{}
	current := 1
	for _, value := range questions {
		jData := entity.JsonData{}
		JAnswers := []entity.Answer{}
		jData.QuestionId = value.QuestionId
		jData.UserQuestion = value.UserQuestion
		jData.WishQuestion = value.WishQuestion

		if current == value.QuestionId {

			for _, valueA := range answers {
				if valueA.QuestionId == current {
					JAnswers = append(JAnswers, valueA)
				}

			}
			jData.Answr = JAnswers
			jsonDataS = append(jsonDataS, jData)
			current++

		}

	}

	jsonFile, err := json.MarshalIndent(jsonDataS, " ", "\t")

	if err != nil {
		w.Header().Set("content-type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(jsonFile)
	return

}
