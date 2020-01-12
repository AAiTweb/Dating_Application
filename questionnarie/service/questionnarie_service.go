package service

import (
	"github.com/betse/Dating_Application-master/entity"
	"github.com/betse/Dating_Application-master/questionnarie"
)

type QuestionnaireServiceImpl struct {
	QuestionRepo questionnarie.QuestionnarieRespository
}

func NewQuestionnaireServiceImpl(QuestionRepo questionnarie.QuestionnarieRespository) *QuestionnaireServiceImpl {
	return &QuestionnaireServiceImpl{QuestionRepo: QuestionRepo}
}

func (qsi *QuestionnaireServiceImpl) Questions() ([]entity.Questionnarie, []entity.Answer, error) {
	questions, answers, err := qsi.QuestionRepo.Questions()

	if err != nil {
		return nil, nil, err
	}
	// log.Println("questions[0].Question")
	return questions, answers, nil
}

func (qsi *QuestionnaireServiceImpl) PostAnswers(userChoice *entity.UserChoice) (*entity.UserChoice, error) {
	usrChoice, err := qsi.QuestionRepo.PostAnswers(userChoice)
	if err != nil {
		return nil, err
	}
	return usrChoice, nil
}
