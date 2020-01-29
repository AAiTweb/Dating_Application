package Handler

import (
	"encoding/json"
	"github.com/AAiTweb/Dating_Application/HomeApi"
	"github.com/AAiTweb/Dating_Application/relationship"
	"github.com/AAiTweb/Dating_Application/session"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

type HomeApiHandeler struct {
	repository  HomeApi.HomeApiRepository
	service     HomeApi.HomeApiService
	relationshp relationship.RelationshipService

}

func NewHomeApiHandler(r HomeApi.HomeApiRepository, s HomeApi.HomeApiService, relation relationship.RelationshipService) HomeApiHandeler {
	return HomeApiHandeler{repository: r, service: s, relationshp: relation}
}

func (H *HomeApiHandeler) GetUsersMatched(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.Method == http.MethodGet {
		id, _ := strconv.Atoi(p.ByName("id"))
		matches, err := H.service.GetMatches(id)
		if err != nil {
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(matches)
	}

}

func (H *HomeApiHandeler) SendRequest(w http.ResponseWriter, r *http.Request) {
	// /sendrequest
	if r.Method == http.MethodGet {
		id,_ := strconv.Atoi(r.URL.Query().Get("id"))
		recieverId := id
		senderId:= session.GetSessionData(w,r).Id
		//payload := struct {
		//	senderId   int
		//	recieverId int
		//}{}
		//json.NewDecoder(r.Body).Decode(payload)
		H.relationshp.SendRequest(senderId, recieverId)
		http.Redirect(w,r,"/home",http.StatusSeeOther)
	}
}

func (H *HomeApiHandeler)Search(w http.ResponseWriter, r*http.Request,p httprouter.Params){
	// /Home/Search/{uname}


	if r.Method == http.MethodGet{
		uname := p.ByName("uname")
		claim  := session.GetSessionData(w,r)
		if claim==nil{
			log.Println(claim)
			log.Fatal("...Search....")
		}
		userId:=claim.Id


		matches,_ := H.service.SearchByName(userId,uname)
		jsonMatches,_ := json.Marshal(matches)
		log.Println(jsonMatches)
		w.Write(jsonMatches)
	}

}

func (H *HomeApiHandeler) AcceptRequest(w http.ResponseWriter, r *http.Request) {
	// /acceptrequest
	if r.Method == http.MethodPost {
		payload := struct {
			senderId   int
			recieverId int
		}{}
		json.NewDecoder(r.Body).Decode(payload)
		H.relationshp.AcceptRequest(payload.senderId, payload.recieverId)
	}
}
