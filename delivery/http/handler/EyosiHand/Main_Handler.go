package EyosiHand

import (
	"bytes"
	"encoding/json"
	"github.com/AAiTweb/Dating_Application/HomeApi/Models"
	"github.com/AAiTweb/Dating_Application/session"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type MainHandler struct {
	Templ *template.Template

}

type HeadData struct {
Name string
ProfilePicture string
}

func NewMainHandler(t *template.Template) MainHandler {
	return MainHandler{t}
}

func (m MainHandler) Index(w http.ResponseWriter, r *http.Request) {
	username := session.GetSessionData(w,r).Username
	profilepicture := session.GetSessionData(w,r).ProfilePicture
	headData := HeadData{Name:username,ProfilePicture:profilepicture}
	tempData := map[string]interface{}{
		"Header": headData,
	}
	m.Templ.ExecuteTemplate(w, "chatpage.html", tempData)

}
func (m MainHandler) Search(w http.ResponseWriter, r*http.Request){
	searchWord := r.URL.Query().Get("username")

	path := "http://:8081/Home/Search/"+searchWord
	req,err := http.NewRequest(http.MethodGet,path,nil)
	tk, _ := r.Cookie("token")
	req.AddCookie(tk)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err!=nil{
		log.Fatal(".....Search1")
	}else{
			jsondata,_ := ioutil.ReadAll(resp.Body)
			data := make([]Models.UserMatch,0)
			json.Unmarshal(jsondata,&data)

			var templates = make([]string,0)
			for _,val := range data{
				var tpl bytes.Buffer
				m.Templ.ExecuteTemplate(&tpl, "card", val)
				templates = append(templates,tpl.String())
			}

			temp,_ := json.Marshal(templates)

			w.Write(temp)
	}
	//response, err := http.Get(path)
	//log.Println(err)

	//if err!=nil{
	//	log.Fatal(".....Search1")
	//}else {
	//	jsondata,_ := ioutil.ReadAll(response.Body)
	//	data := make([]Models.UserMatch,0)
	//	json.Unmarshal(jsondata,&data)
	//	var templates []string
	//	for _,val := range data{
	//		var tpl bytes.Buffer
	//		m.Templ.ExecuteTemplate(&tpl, "card", val)
	//		templates = append(templates,tpl.String())
	//	}
	//	temp,_ := json.Marshal(templates)
	//	w.Write(temp)
	//}
	//var templates []string
	//
	//data := Models.UserMatch{5,29, "Mexico", "Mexico", "p3ic5.jpg" ,"Abebe", 90}
	//



}
func (m MainHandler) Home(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet{
		id := session.GetSessionData(w,r).Id
		path := "http://:8081/matches/user/"+strconv.Itoa(id)
		response, err := http.Get(path)
		if err!=nil{
			log.Println("errpr....")
		}else{
			jsondata,_ := ioutil.ReadAll(response.Body)
			data := make([]Models.UserMatch,0)
			json.Unmarshal(jsondata,&data)

			username := session.GetSessionData(w,r).Username
			profilepicture := session.GetSessionData(w,r).ProfilePicture
			headData := HeadData{Name:username,ProfilePicture:profilepicture}
			tempData := map[string]interface{}{
				"Header": headData,
				"Board":data,
			}
			log.Println(data)
			m.Templ.ExecuteTemplate(w, "dashboard.html", tempData)
		}
	}


}

func (m MainHandler) Preload(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet{
		id := request.URL.Query().Get("id")
		log.Println(".............preload............")
		log.Println(id)
		log.Println(".............preload............")
		m.Templ.ExecuteTemplate(writer,"preload.html",id)
	}

}