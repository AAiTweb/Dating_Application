package Socket

import (
	"github.com/Eyosi-G/Dating_Application/message"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type SocketHandler struct {
	Upgrader websocket.Upgrader
	Conncetions map[int]*websocket.Conn // id and conn
	MService message.MessageService
}

//type of requests --- delete,add,messages


func (sh *SocketHandler) Handler(w http.ResponseWriter, r *http.Request){
	conn,err := sh.Upgrader.Upgrade(w,r,nil)
	if err!=nil{
		log.Println(err)
		return
	}
	sh.Conncetions[1] = conn
	//log.Print(sh.Conncetions)
	sh.readMessage(conn)
}
func (sh *SocketHandler) readMessage(c *websocket.Conn){
	for{
		message := &UserRequest{}

		err := c.ReadJSON(message)
		if err != nil{
			return
		}
		//store it in database
		switch message.Type {
		case "ADD":
			err = sh.MService.SaveMessage(message.Message)
			if err!=nil{
				return
			}
			err1  := sh.Conncetions[message.Message.ToId].WriteJSON(message.Message)
			if err1 != nil{
				sh.Conncetions[message.Message.FromId].WriteJSON(nil)
				return
			}
			err2 := sh.Conncetions[message.Message.FromId].WriteJSON(message.Message)
			if err2 != nil{
				return
			}
		//case "DELETE":
		//	err = sh.MService.DeleteMessage(message.Message)
		//	if err==nil{
		//		messages = UserResponse{1,}
		//	}
		case "MESSAGES":
			msgs := sh.MService.Messages(message.Message.FromId,message.Message.ToId)
			err  = sh.Conncetions[message.Message.FromId].WriteJSON(msgs)
			if err != nil{
				return
			}
		}


	}
}