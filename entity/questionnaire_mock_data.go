package entity

var Questions=[]Questionnarie{
	{1,"user question 1","wish question 1"},
	{2,"user question 2","wish question 2"},
	{3,"user question 3","wish question 3"},
	{4,"user question 4","wish question 4"},
	{5,"user question 5","wish question 5"},


}
var Answers=[]Answer{
	{1, 1,"answer 1 1"},
	{2, 1,"answer 2 1"},
	{1, 2,"answer 1 2"},
	{2, 2,"answer 2 2"},
	{1, 3,"answer 1 3"},
	{2, 3,"answer 2 3"},
	{1, 4,"answer 1 4"},
	{2, 4,"answer 2 4"},
	{1, 5,"answer 1 5"},
	{2, 5,"answer 2 5"},

}

//var userChoice=[]UserChoice{
//	{1,1,1,1},
//	{1,1,1,1},
//	{1,2,2,2},
//	{1,1,1,1},
//	{1,2,2,2},
//	{1,1,1,1},
//}
var userChoice=UserChoice{
	UserId:       1,
	QuestionId:   1,
	OwnAnswerId:  1,
	WishAnswerId: 1,
}
