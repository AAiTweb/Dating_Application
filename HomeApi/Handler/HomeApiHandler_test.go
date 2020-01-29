package Handler

import (
	HomeRepo "github.com/AAiTweb/Dating_Application/HomeApi/Repository"
	HomeService "github.com/AAiTweb/Dating_Application/HomeApi/Service"
	"github.com/AAiTweb/Dating_Application/relationship/repository"
	"github.com/AAiTweb/Dating_Application/relationship/service"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsersMatched(t *testing.T){

	// relationship
	//_relationship := MokeDatabase.Relationships
	//_userPro := MokeDatabase.UserProfile
	//_gallery := MokeDatabase.Gallery
	//_matches := MokeDatabase.Matches
	//_users := MokeDatabase.Users

	//relationshipRepo :=repository.NewFakeRelationshipRepo(_relationship)
	//relationshipService :=service.NewRelationshipService(relationshipRepo)
	//HomeApiRepository := HomeRepo.NewFakeHomeApiRepo()
	//HomeApiService := HomeService.NewHomeApiService(HomeApiRepository)
	//HomeApiHandler := NewHomeApiHandler(HomeApiRepository, HomeApiService, nil)
	//
	//req,_:=  http.NewRequest("GET","/matches/user/1",nil)
	//req.Header.Set("Content-Type", "application/json")
	//
	//rr := httptest.NewRecorder()
	//
	//router := mux.NewRouter()
	//router.HandleFunc("/matches/user/{id}",HomeApiHandler.GetUsersMatched)
	//router.ServeHTTP(rr,req)
	//
	//if status := rr.Code; status!=http.StatusOK {
	//	t.Errorf("handler returned wrong status code: got %v want %v",
	//		status, http.StatusOK)
	//}
	//expected,_ := json.Marshal(MokeDatabase.UsersMatches[1])
	//if string(expected)!=rr.Body.String(){
	//	t.Errorf("handler returned unexpected body: got %v want %v",
	//		rr.Body.String(), string(expected))
	//}
}

func TestSendRequest(t *testing.T){
	//_relationship := MokeDatabase.Relationships
	//_userPro := MokeDatabase.UserProfile
	//_gallery := MokeDatabase.Gallery
	//_matches := MokeDatabase.Matches
	//_users := MokeDatabase.Users

	relationshipRepo :=repository.NewFakeRelationshipRepo()
	relationshipService :=service.NewRelationshipService(relationshipRepo)
	HomeApiRepository := HomeRepo.NewFakeHomeApiRepo()
	HomeApiService := HomeService.NewHomeApiService(HomeApiRepository)
	HomeApiHandler := NewHomeApiHandler(HomeApiRepository, HomeApiService, relationshipRepo)

	req,_:=  http.NewRequest("GET","/matches/sendrequest?id=1",nil)
	//req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/matches/sendrequest",HomeApiHandler.SendRequest)
	router.ServeHTTP(rr,req)

	if status := rr.Code; status!=http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	//expected,_ := json.Marshal("User Updated")
	//if string(expected)!=rr.Body.String(){
	//	t.Errorf("handler returned unexpected body: got %v want %v",
	//		rr.Body.String(), string(expected))
	//}
}

func TestSearch(t *testing.T){


}