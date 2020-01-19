package questionnarie

import (
	"github.com/biniyam112/TheDatingApp/Dating_Application/entity"
)

type QuestionnarieService interface {
	Questions() ([]entity.Questionnarie, []entity.Answer, error)
	PostAnswers(userCoice *entity.UserChoice) (*entity.UserChoice, error)
}
