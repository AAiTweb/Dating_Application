package questionnarie

import (
	"github.com/betse/Dating_Application-master/entity"
)

type QuestionnarieService interface {
	Questions() ([]entity.Questionnarie, error)
}
