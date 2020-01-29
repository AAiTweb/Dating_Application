package handler

import (
	"fmt"
	"github.com/AAiTweb/Dating_Application/entity"
	"github.com/AAiTweb/Dating_Application/notification/service"
	"github.com/AAiTweb/Dating_Application/session"
	"html/template"
	"net/http"
)

type NotifHandler struct {
	Templ     template.Template
	Notifserv service.NotifService
}
type HeadData struct {
	Name string
	ProfilePicture string
}

func NewMainHandler(servc *service.NotifService, temp *template.Template) *NotifHandler {
	return &NotifHandler{
		Templ:     *temp,
		Notifserv: *servc,
	}
}


func (mh NotifHandler) SeeNotification(w http.ResponseWriter, r *http.Request) {
	fmt.Print("in notification page\n")
	userData := session.GetSessionData(w,r)
	if userData == nil{
		print("Session can't be generated")
		_ = mh.Templ.ExecuteTemplate(w, "logincopy.html", "User doesn't exits")
		return
	}
	notifInstance, err := mh.Notifserv.AddNotification(userData.Id)
	username := session.GetSessionData(w,r).Username
	profilepicture := session.GetSessionData(w,r).ProfilePicture
	headData := HeadData{Name:username,ProfilePicture:profilepicture}
	tempData := map[string]interface{}{
		"Header": headData,
		"NotificationData" : notifInstance,
	}
	if err != nil {
		mh.Templ.ExecuteTemplate(w,"notificationPage.html", "You have no notification")
	}
	mh.Templ.ExecuteTemplate(w,"notificationPage.html", tempData)
}
func (mh NotifHandler) AcceptNotification(w http.ResponseWriter, r *http.Request) {
	sender_name := r.FormValue("sender_name")
	relation := entity.Relationship{
		RelationshipId:     0,
		SenderId:           mh.Notifserv.NotifInstance.GetId(sender_name),
		RecieverId:         0,
		RelationShipStatus: 0,
	}
	userData := session.GetSessionData(w,r)
	err := mh.Notifserv.AcceptNotification(relation,userData.Id)
	fmt.Println("accept: " + sender_name)
	if err != nil {
		mh.Templ.ExecuteTemplate(w, "notificationPage.html","You have no notification")
	}
	mh.SeeNotification(w, r)
}
func (mh NotifHandler) RejectNotification(w http.ResponseWriter, r *http.Request) {
	sender_name := r.FormValue("sender_name")
	relation := entity.Relationship{
		RelationshipId:     0,
		SenderId:           mh.Notifserv.NotifInstance.GetId(sender_name),
		RecieverId:         0,
		RelationShipStatus: 0,
	}
	userData := session.GetSessionData(w,r)
	err := mh.Notifserv.RejectNotification(relation,userData.Id)
	fmt.Println("reject: " + sender_name)
	if err != nil {
		mh.Templ.ExecuteTemplate(w,"notificationPage.html", "You have no notification")
	}
	mh.SeeNotification(w, r)
}