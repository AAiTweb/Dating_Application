package repository

import (
	"database/sql"
	"github.com/AAiTweb/Dating_Application/entity"
	"github.com/AAiTweb/Dating_Application/questionnarie"
)

type MockQuesRepository struct {
	conn *sql.DB
}
func NewMockQuesRepo(db *sql.DB )questionnarie.QuestionnarieRespository{
	return &MockQuesRepository{conn:db}
}
func(mockQ *MockQuesRepository) Questions()([]entity.Questionnarie, []entity.Answer, error){
	return entity.Questions,entity.Answers,nil
}

func (mockQ *MockQuesRepository)PostAnswers(userChoice *entity.UserChoice)(*entity.UserChoice, error ) {
	userCh:=userChoice
	return userCh,nil


}