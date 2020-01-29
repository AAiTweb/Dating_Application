package MatchHandler

import (
	"github.com/AAiTweb/Dating_Application/Matches"
	"log"
	"net/http"
	"strconv"
)

type MatchHandler struct {
	MatchService Matches.MatchService
}

func NewMatchHandler(matchService Matches.MatchService)MatchHandler{
	return MatchHandler{MatchService:matchService}
}
func (m *MatchHandler) DoMatching(w http.ResponseWriter , r * http.Request){
	if r.Method == http.MethodGet{
		id , _ := strconv.Atoi(r.URL.Query().Get("id"))
		log.Println(id)

		result := m.MatchService.DoMatch(id)
		if result!=nil{
			//error
			return
			log.Fatal("............In Handler")
		} else{
			// success
			http.Redirect(w,r,"/home",http.StatusSeeOther)
		}
	}

}