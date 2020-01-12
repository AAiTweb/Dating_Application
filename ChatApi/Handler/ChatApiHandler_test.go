package Handler

import (
	"encoding/json"
	"github.com/Eyosi-G/Dating_Application/ChatApi/repository"
	service2 "github.com/Eyosi-G/Dating_Application/ChatApi/service"
	repository1 "github.com/Eyosi-G/Dating_Application/message/repository"
	"github.com/Eyosi-G/Dating_Application/message/service"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFriends(t *testing.T){

	ChatApirepo := repository.NewFakeApiRepository(&repository.FriendList,&repository.LoginDetails)
	messageRepo  := repository1.NewFakeMessageRepository(&repository.Messages)
	msgService := service.NewMessageService(messageRepo)
	apiservice := service2.NewApiService(ChatApirepo)
	handler := NewApiHandler(msgService,apiservice)
	req,_:=  http.NewRequest("GET","/user/1/friends",nil)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/user/{id}/friends",handler.GetFriends)
	router.ServeHTTP(rr,req)

	if status := rr.Code; status!=http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected,_ := json.Marshal(repository.FriendList[1])
	if string(expected)!=rr.Body.String(){
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expected))
	}
}