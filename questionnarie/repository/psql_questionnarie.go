package repository

import (
	"database/sql"
	"errors"

	"github.com/AAiTweb/Dating_Application/entity"
)

type QuestionnarieRespositoryImpl struct {
	Conn *sql.DB
}

func NewQuestionnarieRepoImpl(Conn *sql.DB) *QuestionnarieRespositoryImpl {
	return &QuestionnarieRespositoryImpl{Conn: Conn}
}
func (qri *QuestionnarieRespositoryImpl) getAnswers(entity.Answer) {

}
func (qri *QuestionnarieRespositoryImpl) PostAnswers(userChoice *entity.UserChoice) (*entity.UserChoice, error) {
	// usrChoice := userChoice

	_, err := qri.Conn.Exec("INSERT INTO user_own_answer(user_own_id,own_question_id,own_choice_answer_id) values($1,$2,$3) ", userChoice.UserId, userChoice.QuestionId, userChoice.OwnAnswerId)
	_, err = qri.Conn.Exec("INSERT INTO user_wish_answer(user_wish_id,wish_question_id,wish_choice_answer_id) values($1,$2,$3) ", userChoice.UserId, userChoice.QuestionId, userChoice.WishAnswerId)
	if err != nil {
		return nil, err
	}
	_,err =qri.Conn.Exec("update users set quefilled=1 where user_id = $1",userChoice.UserId)

	if err != nil {
		return nil, err
	}

	return userChoice, nil
}

func (qri *QuestionnarieRespositoryImpl) Questions() ([]entity.Questionnarie, []entity.Answer, error) {
	query := `SELECT
	questionnaires.questionnaire_id,
	questionnaires.user_own_questions,
	questionnaires.user_wish_questions,
	questionnaire_choices.choice_id,
	questionnaire_choices.choice_questionnaire_id,
	questionnaire_choices.choice

	
	
			FROM
			questionnaires
			INNER JOIN 	questionnaire_choices ON questionnaires.questionnaire_id=questionnaire_choices.choice_questionnaire_id;
		`
	// query2 := "SELECT * FROM dating_app.questionnaires; "

	rows, err := qri.Conn.Query(query)
	if err != nil {
		return nil, nil, errors.New("Could not query the database")
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
			return nil, nil, err
		}
		questions = append(questions, question)
		answers = append(answers, answer)

	}

	// jsonDataS := []entity.JsonData{}
	// current := 1
	// for _, value := range questions {
	// 	jData := entity.JsonData{}
	// 	JAnswers := []entity.Answer{}
	// 	jData.QuestionId = value.QuestionId
	// 	jData.UserQuestion = value.UserQuestion
	// 	jData.WishQuestion = value.WishQuestion

	// 	if current == value.QuestionId {

	// 		for _, valueA := range answers {
	// 			if valueA.QuestionId == current {
	// 				JAnswers = append(JAnswers, valueA)
	// 			}

	// 		}
	// 		jData.Answr = JAnswers
	// 		jsonDataS = append(jsonDataS, jData)
	// 		current++

	// 	}

	// }

	// jsonFile, _ := json.MarshalIndent(jsonDataS, " ", "\t")

	// err = ioutil.WriteFile("../../entity/questions.json", jsonFile, 0644)

	// if err != nil {
	// 	log.Fatal(err)
	// 	return nil, nil, err
	// }
	// for _, value := range questions {

	// 	}

	// log.Println(len(questions), "length")
	// for value, _ := range questions {

	// 	log.Println(value)

	// 	println(" ", index)
	// }
	return questions, answers, nil
}
