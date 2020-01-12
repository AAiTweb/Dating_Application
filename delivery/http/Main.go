package main

import (
	"database/sql"
	"errors"
	"github.com/Eyosi-G/Dating_Application/ChatApi/Handler"
	"github.com/Eyosi-G/Dating_Application/session"

	repository2 "github.com/Eyosi-G/Dating_Application/ChatApi/repository"
	service2 "github.com/Eyosi-G/Dating_Application/ChatApi/service"
	Socket2 "github.com/Eyosi-G/Dating_Application/Socket"
	"github.com/Eyosi-G/Dating_Application/message/repository"
	"github.com/Eyosi-G/Dating_Application/message/service"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
)

var upgrader websocket.Upgrader
var users  = make(map[int]*websocket.Conn)
var templ = template.Must(template.ParseGlob("../../ui/templates/*.html"))

var db,err = sql.Open("postgres","postgres://postgres:password@localhost/dating_app?sslmode=disable")




func main() {

	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		panic(err)
	}





	//sampleMessage := entity.Message{-1,1,4,"hello",time.Now(),2}

	//repositoryMessage := repository.RepositoryMessage{db}
	//serviceMessage := service.MessageService{repositoryMessage}



	//api
	//Handler := ChatApi.APIHandler{Db:db}

	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("../../ui/assets"))
	messageReppo := repository.NewRepositoryMessage(db)
	apirepo := repository2.NewApiRepository(db)

	msgService := service.NewMessageService(messageReppo)
	apiservice := service2.NewApiService(apirepo)
	handler := Handler.NewApiHandler(msgService,apiservice)
	socketHandler := Socket2.NewSocketHandler(upgrader,users,msgService)

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/",fs))
	router.Handle("/chat",session.IsAuthenticated(http.HandlerFunc(index)))
	router.HandleFunc("/login",login)

	log.Println("Server Started Listening ")

	router.HandleFunc("/user/{id}/friends",handler.GetFriends)
	router.HandleFunc("/chats/user/{uid}/friends/{fid}",handler.GetMessages)
	router.HandleFunc("/user/{id}/updatelogin",handler.UpdateLoginDetails)
	router.HandleFunc("/ws",socketHandler.Socket)
	http.ListenAndServe("localhost:8081",router)


	//fmt.Println(time.Now())


	//fmt.Println(serviceMessage.Messages(1,5))


	//userrepoInstance := repository.UserRepositoryInstance{db}
	//userserviceInstance:=service.UserServiceInstance{RepositoryInstance:userrepoInstance}
	//handlerInstance := handler.MainHandler{
	//	Templ:    templ,
	//	Uservice: userserviceInstance,
	//}
	//:= http.NewServeMux()

	//mux.HandleFunc("/login",handler.Login)
	//mux.HandleFunc("/signup",handler.Signup)
	//mux.HandleFunc("/notification",handler.Notification)
	//mux.HandleFunc("/message",handler.Message)
	//mux.HandleFunc("/profile",handler.Profile)
	//mux.HandleFunc("/home",handlerInstance.Home)
	//mux.HandleFunc("/questionnaire",handler.Questionnaire)
	//mux.HandleFunc("/",handlerInstance.Home)
	//
	//http.ListenAndServe("localhost:8082",mux)
	//

}

func login(writer http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet{
		templ.ExecuteTemplate(writer,"login.html",nil)
	}else{
		uname := request.FormValue("username")
		pword := request.FormValue("password")
		//
		//log.Println(uname)
		//log.Println(pword)

		//row:= db.QueryRow("select user_id,username from users where username=$1 and password=$2;",uname,pword)
		row := db.QueryRow(`select t1.user_id,t1.username,t3.picture_path
		from (select user_id, username from users where username=$1 and password=$2) as t1
			inner join (select picture_owner_id,picture_path from user_profile inner join gallery on picture_id=profile_picture) t3 on
				t3.picture_owner_id=t1.user_id;`,uname,pword)
		usr := struct {
			Id int
			UserName,
			ProfilePicture string
		}{}
		row.Scan(&usr.Id,&usr.UserName,&usr.ProfilePicture)
		log.Println(usr)


		if err!=nil && err==sql.ErrNoRows {
			errors.New("User Doesn't Exist")
		}else{
			//expirationTime := time.Now().Add(5 * time.Minute)
			//claims := &entity.Claims{
			//	Username: usr.UserName,
			//	Id: usr.Id,
			//	ProfilePicture : usr.ProfilePicture,
			//	StandardClaims: jwt.StandardClaims{
			//		// In JWT, the expiry time is expressed as unix milliseconds
			//		ExpiresAt: expirationTime.Unix(),
			//	},
			//}
			////log.Println(*claims)
			//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			//tokenString, _ := token.SignedString(entity.JwtKey)
			tokenString,err := session.Generate(usr.Id,usr.UserName,usr.ProfilePicture)
			if err!=nil{
				return
			}
			http.SetCookie(writer, &http.Cookie{
				Name:    "token",
				Value:   tokenString,
			})
			http.Redirect(writer,request,"/chat",http.StatusSeeOther)
		}










	}

}


//func isAuthorized(next http.Handler)http.Handler{
//	fn := func (w http.ResponseWriter, r *http.Request){
//		c, _ := r.Cookie("token")
//		tknStr := c.Value
//		claims := &entity.Claims{}
//		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
//			return entity.JwtKey, nil
//		})
//		if err != nil {
//			if err == jwt.ErrSignatureInvalid {
//				w.Write([]byte("access not autorized"))
//				return
//			}
//			w.Write([]byte("access not autorized"))
//			return
//		}
//		if !tkn.Valid {
//			w.Write([]byte("access not autorized"))
//			return
//		}
//		next.ServeHTTP(w,r)
//	}
//	return http.HandlerFunc(fn)
//}



func index(w http.ResponseWriter, r *http.Request){
	//c, _ := r.Cookie("token")
	//tknStr := c.Value
	//claims := &entity.Claims{}
	//tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
	//	return entity.JwtKey, nil
	//})
	//if err != nil {
	//	if err == jwt.ErrSignatureInvalid {
	//		w.Write([]byte("access not autorized"))
	//		return
	//	}
	//	w.Write([]byte("access not autorized"))
	//	return
	//}
	//if !tkn.Valid {
	//	w.Write([]byte("access not autorized"))
	//	return
	//}
	templ.ExecuteTemplate(w,"chatpage.html",nil)
}
//func Socket(w http.ResponseWriter, r *http.Request)  {
//	upgrader.CheckOrigin  = func(r *http.Request) bool {
//		return true;
//	}
//	conn,err := upgrader.Upgrade(w,r,nil)
//	if err!=nil{
//		return
//	}
//	for{
//
//		messageType, message,_ := conn.ReadMessage()
//		jmessage := struct {
//			SenderId int
//			ReceiverId int
//			MessageText string
//			Time string
//		}{}
//		json.Unmarshal(message,&jmessage)
//		log.Println(jmessage)
//
//		//msg := struct {
//		//	Message string
//		//	Time string
//		//}{}
//		//msg.Message = string(message)
//		jmessage.Time = ChatApi.MessageSendTimeChanger(time.Now())
//		messageByte,_ := json.Marshal(jmessage)
//		conn.WriteMessage(messageType,messageByte)
//	}
//
//
//	//sh.Conncetions[int(id)]= conn;
//	//sh.readMessage(conn)
//}
