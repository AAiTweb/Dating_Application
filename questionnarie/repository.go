package questionnarie

import (
	"github.com/betse/Dating_Application-master/entity"
)

type QuestionnarieRespository interface {
	Questions() ([]entity.Questionnarie, []entity.Answer, error)
	PostAnswers(userCoice *entity.UserChoice) (*entity.UserChoice, error)
}
