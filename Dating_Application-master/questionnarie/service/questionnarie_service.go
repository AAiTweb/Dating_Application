package service

import (
	"log"

	"github.com/betse/Dating_Application-master/entity"
	"github.com/betse/Dating_Application-master/questionnarie"
)

type QuestionnaireServiceImpl struct {
	QuestionRepo questionnarie.QuestionnarieRespository
}

func NewQuestionnaireServiceImpl(QuestionRepo questionnarie.QuestionnarieRespository) *QuestionnaireServiceImpl {
	return &QuestionnaireServiceImpl{QuestionRepo: QuestionRepo}
}

func (qsi *QuestionnaireServiceImpl) Questions() ([]entity.Questionnarie, error) {
	questions, err := qsi.QuestionRepo.Questions()
	if err != nil {
		return nil, err
	}
	log.Println("questions[0].Question")
	return questions, nil
}
