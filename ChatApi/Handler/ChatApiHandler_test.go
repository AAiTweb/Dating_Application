package Handler

import (
	"github.com/AAiTweb/Dating_Application/ChatApi/repository"
	service2 "github.com/AAiTweb/Dating_Application/ChatApi/service"
	repository1 "github.com/AAiTweb/Dating_Application/message/repository"
	"github.com/AAiTweb/Dating_Application/message/service"

	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFriends(t *testing.T) {

	ChatApirepo := repository.NewFakeApiRepository(&repository.FriendList, &repository.LoginDetails)
	messageRepo := repository1.NewFakeMessageRepository(&repository.Messages)
	msgService := service.NewMessageService(messageRepo)
	apiservice := service2.NewApiService(ChatApirepo)
	handler := NewApiHandler(msgService, apiservice)
	req, _ := http.NewRequest("GET", "/user/1/friends", nil)
	rr := httptest.NewRecorder()
	handler_ := http.HandlerFunc(handler.GetFriends)
	handler_.ServeHTTP(rr, req)
	log.Println(rr.Body.String())

}
