package questionnarie

import (
	"github.com/AAiTweb/Dating_Application/entity"
)

type QuestionnarieRespository interface {
	Questions() ([]entity.Questionnarie, []entity.Answer, error)
	PostAnswers(userCoice *entity.UserChoice) (*entity.UserChoice, error)
}
