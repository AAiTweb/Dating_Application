package Handler

import (
	"encoding/json"
	"github.com/AAiTweb/Dating_Application/ChatApi/repository"
	service2 "github.com/AAiTweb/Dating_Application/ChatApi/service"
	"github.com/AAiTweb/Dating_Application/MokeDatabase"
	repository1 "github.com/AAiTweb/Dating_Application/message/repository"
	"github.com/AAiTweb/Dating_Application/message/service"
	"github.com/julienschmidt/httprouter"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFriends(t *testing.T) {

	ChatApirepo := repository.NewFakeApiRepository(&MokeDatabase.FriendList,&MokeDatabase.LoginDetails)
	messageRepo  := repository1.NewFakeMessageRepository(&MokeDatabase.Messages)
	msgService := service.NewMessageService(messageRepo)
	apiservice := service2.NewApiService(ChatApirepo)
	handler := NewApiHandler(msgService,apiservice)
	req,_:=  http.NewRequest("GET","/user/friends/1",nil)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router := httprouter.New()
	router.GET("/user/friends/:id",handler.GetFriends)
	router.ServeHTTP(rr,req)

	if status := rr.Code; status!=http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected,_ := json.Marshal(MokeDatabase.FriendList[1])
	if string(expected)!=rr.Body.String(){
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expected))
	}

}
