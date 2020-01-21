package Socket

import (
	"encoding/json"
	"github.com/AAiTweb/Dating_Application/ChatApi"
	_ "github.com/AAiTweb/Dating_Application/entity"
	"github.com/AAiTweb/Dating_Application/message"
	"github.com/AAiTweb/Dating_Application/session"
	"github.com/gorilla/websocket"
	"log"
	_ "log"
	"net/http"
	"time"
)

type SocketHandler struct {
	Upgrader websocket.Upgrader
	Conncetions map[int]*websocket.Conn // id and conn
	MessageService message.MessageService
	//MService service.MessageService
}

func NewSocketHandler(upgrader websocket.Upgrader, connections map[int]*websocket.Conn, messageservice message.MessageService)SocketHandler{
	return SocketHandler{upgrader,connections, messageservice}
}

func (s *SocketHandler)Socket(w http.ResponseWriter, r *http.Request)  {
	s.Upgrader.CheckOrigin  = func(r *http.Request) bool {
		return true;
	}
	conn,err := s.Upgrader.Upgrade(w,r,nil)
	if err!=nil{
		return
	}
	claims := session.GetSessionData(w,r)

	log.Println(claims.Id)
	s.Conncetions[claims.Id] = conn
	for{

		messageType, message,_ := conn.ReadMessage()

		jmessage := struct {
			SenderId int
			ReceiverId int
			MessageText string
			SenderPicture string
			Time string
		}{}
		json.Unmarshal(message,&jmessage)
		//msgs := entity.Message{
		//	FromId: jmessage.SenderId,
		//	ToId:jmessage.ReceiverId,
		//	 Message:jmessage.MessageText,
		//	SendTime:time.Now(),
		//	}
		//	s.MessageService.SaveMessage(msgs)


		jmessage.Time = ChatApi.MessageSendTimeChanger(time.Now())
		messageByte,_ := json.Marshal(jmessage)

		//conn.WriteMessage(messageType,messageByte)
		log.Println(s.Conncetions)
		if _,ok := s.Conncetions[jmessage.SenderId]; ok{
			s.Conncetions[jmessage.SenderId].WriteMessage(messageType,messageByte)
		}
		if _,ok := s.Conncetions[jmessage.ReceiverId];ok{
			s.Conncetions[jmessage.ReceiverId].WriteMessage(messageType,messageByte)

		}


	}

}


