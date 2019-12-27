package questionnarie

import (
	"github.com/betse/Dating_Application-master/entity"
)

type QuestionnarieRespository interface {
	Questions() ([]entity.Questionnarie, error)
}
