package repository

import (
	"database/sql"
	"github.com/AAiTweb/Dating_Application/entity"
	"log"
)

type MatchRepo struct {
	db *sql.DB
}

func (m MatchRepo) DoMatch(id int) error {
	query := "select user_id from users where user_id!=$1"
	row,err:= m.db.Query(query,id)
	if err!=nil{
		return err
	}
	for row.Next(){
		var userId int
		err := row.Scan(&userId)
		if err!=nil{

			return err
		}
	userOtherChoice,_ := m.GetAnswerOtherChoice(userId)
	userWishChoice, _ := m.GetUserChoice(userId)
	matchPercentage := Match(userWishChoice,userOtherChoice)
	query1 := "select exists(select * from matches where match_userone_id=$1 and match_usertwo_id=$2)"
	row1 := m.db.QueryRow(query1,id,userId)
	var isExist bool;
	err = row1.Scan(&isExist)
	if err!=nil{

		return err
	}
	if isExist{
		// update the table
		log.Println(matchPercentage)
		queryUpdate := "update matches set match_point=$1 where match_userone_id=$2 and match_usertwo_id=$3"
		_,err := m.db.Exec(queryUpdate,matchPercentage,id,userId)
		if err!=nil{

			log.Fatal(".....Here....")
			return err
		}
	}else{

		queryInsert := "insert into matches(match_userone_id,match_usertwo_id,match_point) values($1,$2,$3);"
		_,err := m.db.Exec(queryInsert,id,userId,matchPercentage)
		if err!=nil{
			log.Fatal(".....Here2")
			return err
		}
	}
	}
	return nil;
}

func NewMatchRepository(_db *sql.DB)MatchRepo{
	return MatchRepo{db:_db}
}

func (m MatchRepo)GetUserChoice(id int) ([]entity.UserChoice, error) {
	query := "select user_wish_id,wish_question_id,wish_choice_answer_id  from user_wish_answer where user_wish_id=$1"
	row, _ := m.db.Query(query,id)
	userWishAnswers := []entity.UserChoice{}
	for row.Next(){
		uChoice := entity.UserChoice{}
		row.Scan(&uChoice.UserId, &uChoice.QuestionId,&uChoice.WishAnswerId)
		userWishAnswers = append(userWishAnswers,uChoice)
	}
	return userWishAnswers,nil

}

func (m MatchRepo) GetAnswerOtherChoice(id int) ([]entity.UserChoice, error) {
	query := "select user_own_id,own_question_id,own_choice_answer_id  from user_own_answer where user_own_id=$1"
	row, _ := m.db.Query(query,id)
	userOwnAnswers := []entity.UserChoice{}
	for row.Next(){
		uChoice := entity.UserChoice{}
		row.Scan(&uChoice.UserId, &uChoice.QuestionId,&uChoice.OwnAnswerId)
		userOwnAnswers = append(userOwnAnswers,uChoice)
	}
	return userOwnAnswers,nil
}



