package handler

import (
	"html/template"
	"net/http"

	"github.com/betse/Dating_Application-master/questionnarie"
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
func (userQ *UserQuestionnarieHandler) Questionnaire(w http.ResponseWriter, r *http.Request) {
	questions, err := userQ.questionnaireServ.Questions()
	if err != nil {
		panic(err)
	}
	if questions == nil {
		println("empty")
	} else {
		println("not empty")

	}
}
