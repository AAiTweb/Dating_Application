package entity

type Questionnarie struct {
	QuestionId   int    `json:"questionId"`
	UserQuestion string `json : "userQuestion"`
	WishQuestion string `json:"wishQuestion"`
}
type Answer struct {
	AnswerId   int    `json:"answerId" `
	QuestionId int    `json:"questionId"`
	Answer     string `json:"answer"`
}

type JsonData struct {
	QuestionId   int      `json:"questionId"`
	UserQuestion string   `json : "userQuestion"`
	WishQuestion string   `json:"wishQuestion"`
	Answr        []Answer `json:"answers"`
}
