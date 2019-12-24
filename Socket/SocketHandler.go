package Socket

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"time"
)

type SocketHandler struct {
	Upgrader websocket.Upgrader
	Conncetions map[int]*websocket.Conn // id and conn
	//MService service.MessageService
}

//type of requests --- delete,add,messages
func (sh *SocketHandler)Socket(w http.ResponseWriter, r *http.Request)  {
	sh.Upgrader.CheckOrigin  = func(r *http.Request) bool {
		return true;
	}
	path := mux.Vars(r)
	id,_  := strconv.ParseInt(path["id"],0,0)

	conn,err := sh.Upgrader.Upgrade(w,r,nil)
	if err!=nil{
		return
	}
	sh.Conncetions[int(id)]= conn;
	//sh.readMessage(conn)
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
			message.Message.SendTime = time.Now()
		//	err = sh.MService.SaveMessage(message.Message)
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
			fmt.Println(message.Limit)
			//msgs,err := sh.MService.Messages(message.Message.FromId,message.Message.ToId)
			if err!=nil{
				return
			}
			//err  = sh.Conncetions[message.Message.FromId].WriteJSON(msgs)
			if err != nil{
				return
			}
		}


	}
}