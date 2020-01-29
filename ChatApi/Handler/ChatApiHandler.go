package Handler

import (
	"encoding/json"
	"github.com/AAiTweb/Dating_Application/ChatApi"
	"github.com/AAiTweb/Dating_Application/message"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"strings"
)

type APIHandler struct {
	msgService message.MessageService
	apiService ChatApi.APIService
}

func NewApiHandler(msgservice message.MessageService, apiService ChatApi.APIService) APIHandler {
	return APIHandler{msgservice, apiService}
}
func (Apihandler *APIHandler) GetFriends(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, _ := strconv.ParseInt(params.ByName("id"), 0, 0)
	friendsInformation, _ := Apihandler.apiService.LoadFriendInformation(int(id))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(friendsInformation)
}
func (Apihandler *APIHandler) GetMessages(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.Method == http.MethodGet {
		userId, _ := strconv.ParseInt(p.ByName("uid"), 0, 0)
		friendId, _ := strconv.ParseInt(p.ByName("fid"), 0, 0)
		messages := Apihandler.msgService.Messages(int(userId), int(friendId))
		type json_message struct {
			MessageId,
			FromId,
			ToId int
			Message  string
			SendTime string
			Status   int
		}
		json_messages := []json_message{}
		for _, msg := range messages {
			jmessage := json_message{}
			jmessage.MessageId = msg.MessageId
			jmessage.FromId = msg.FromId
			jmessage.ToId = msg.ToId
			jmessage.Message = strings.TrimSuffix(strings.TrimPrefix(msg.Message, "'"), "'")
			jmessage.SendTime = ChatApi.MessageSendTimeChanger(msg.SendTime)
			json_messages = append(json_messages, jmessage)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(json_messages)
	}

}
func (Apihandler *APIHandler) UpdateLoginDetails(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.Method == http.MethodGet {
		id, _ := strconv.ParseInt(p.ByName("id"), 0, 0)
		Apihandler.apiService.UpdateLoginInformation(int(id))
	}

}
