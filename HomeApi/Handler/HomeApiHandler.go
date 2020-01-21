package Handler

import (
	"encoding/json"
	"github.com/AAiTweb/Dating_Application/HomeApi"
	"github.com/AAiTweb/Dating_Application/relationship"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type HomeApiHandeler struct {
	repository HomeApi.HomeApiRepository
	service HomeApi.HomeApiService
	relationshp relationship.RelationshipService
}

func NewHomeApiHandler(r HomeApi.HomeApiRepository ,s HomeApi.HomeApiService, relation relationship.RelationshipService)HomeApiHandeler{
	return HomeApiHandeler{repository:r,service:s,relationshp:relation}
}

func (H *HomeApiHandeler)GetUsersMatched(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		id,_ := strconv.Atoi(mux.Vars(r)["id"])
		matches,err := H.service.GetMatches(id)
		if err!=nil{
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(matches)
	}

}

func (H *HomeApiHandeler)SendRequest(w http.ResponseWriter, r *http.Request){
	// /sendrequest
	if r.Method==http.MethodPost{
		payload := struct {
			senderId int
			recieverId int
		}{}
		json.NewDecoder(r.Body).Decode(payload)
		H.relationshp.SendRequest(payload.senderId,payload.recieverId)
	}
}

func (H *HomeApiHandeler)AcceptRequest(w http.ResponseWriter, r *http.Request){
	// /acceptrequest
	if r.Method==http.MethodPost{
		payload := struct {
			senderId int
			recieverId int
		}{}
		json.NewDecoder(r.Body).Decode(payload)
		H.relationshp.AcceptRequest(payload.senderId,payload.recieverId)
	}
}
