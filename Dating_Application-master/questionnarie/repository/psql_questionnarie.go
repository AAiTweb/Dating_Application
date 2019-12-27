package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"github.com/betse/Dating_Application-master/entity"
)

type QuestionnarieRespositoryImpl struct {
	Conn *sql.DB
}

func NewQuestionnarieRepoImpl(Conn *sql.DB) *QuestionnarieRespositoryImpl {
	return &QuestionnarieRespositoryImpl{Conn: Conn}
}
func (qri *QuestionnarieRespositoryImpl) getAnswers(entity.Answer) {

}
func (qri *QuestionnarieRespositoryImpl) Questions() ([]entity.Questionnarie, error) {
	query := `SELECT
	dating_app.questionnaires.questionnaire_id,
	dating_app.questionnaires.user_own_questions,
	dating_app.questionnaires.user_wish_questions,
	dating_app.questionnaire_choices.choice_id,
	dating_app.questionnaire_choices.choice_questionnaire_id,
	dating_app.questionnaire_choices.choice

	
	
			FROM
			dating_app.questionnaires
			INNER JOIN 	dating_app.questionnaire_choices ON dating_app.questionnaires.questionnaire_id=dating_app.questionnaire_choices.choice_questionnaire_id;
		`
	// query2 := "SELECT * FROM dating_app.questionnaires; "

	rows, err := qri.Conn.Query(query)
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()
	// answers := make([]string, 20)
	answers := []entity.Answer{}
	questions := []entity.Questionnarie{}
	for rows.Next() {
		question := entity.Questionnarie{}
		answer := entity.Answer{}
		err := rows.Scan(&question.QuestionId, &question.UserQuestion, &question.WishQuestion, &answer.AnswerId, &answer.QuestionId, &answer.Answer)

		if err != nil {
			return nil, err
		}
		questions = append(questions, question)
		answers = append(answers, answer)
		
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

	

	jsonFile, _ := json.MarshalIndent(jsonDataS, " ", "\t")

	err = ioutil.WriteFile("../../entity/questions.json", jsonFile, 0644)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// for _, value := range questions {

	// 	}

	log.Println(len(questions), "length")
	// for value, _ := range questions {

	// 	log.Println(value)

	// 	println(" ", index)
	// }
	return questions, nil
}
