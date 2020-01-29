package repository

import (
	"github.com/AAiTweb/Dating_Application/entity"
	"math"
)

var weight = map[int]int{1:20,2:10,3:20,4:10,5:10,6:30}
func Match(MyWishAnswer []entity.UserChoice, OtherAnswer []entity.UserChoice)int{
	sum := 0
	for index:= range MyWishAnswer{
		_MyWishAnswer := MyWishAnswer[index]
		_OtherAnswer := OtherAnswer[index]
		if _MyWishAnswer.QuestionId!=5 {

			difference := int(math.Abs(float64((_MyWishAnswer.WishAnswerId - _OtherAnswer.OwnAnswerId))))
			gap := weight[_MyWishAnswer.QuestionId]/3
			switch difference {
			case 0:
				sum += weight[_MyWishAnswer.QuestionId]
			case 1:
				sum += weight[_MyWishAnswer.QuestionId]-(gap*1)
			case 2:
				sum += weight[_MyWishAnswer.QuestionId]-(gap*2)
			}
		}else{
			difference := int(math.Abs(float64((_MyWishAnswer.WishAnswerId - _OtherAnswer.OwnAnswerId))))
			gap := weight[_MyWishAnswer.QuestionId]/5
			switch difference {
			case 0:
				sum += weight[_MyWishAnswer.QuestionId]
			case 1:
				sum += weight[_MyWishAnswer.QuestionId]-(gap*1)
			case 2:
				sum += weight[_MyWishAnswer.QuestionId]-(gap*2)
			case 3:
				sum += weight[_MyWishAnswer.QuestionId]-(gap*3)
			case 4:
				sum += weight[_MyWishAnswer.QuestionId]-(gap*4)
			}
		}

	}
	return sum
}


