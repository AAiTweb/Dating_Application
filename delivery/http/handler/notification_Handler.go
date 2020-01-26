package handler

import (
	"fmt"
	"github.com/AAiTweb/Dating_Application/entity"
	"github.com/AAiTweb/Dating_Application/notification/service"
	"html/template"
	"net/http"
)

type NotifHandler struct {
	Templ     template.Template
	Notifserv service.NotifService
}

func NewMainHandler(servc *service.NotifService, temp *template.Template) *NotifHandler {
	return &NotifHandler{
		Templ:     *temp,
		Notifserv: *servc,
	}
}


func (mh NotifHandler) SeeNotification(w http.ResponseWriter, r *http.Request) {
	fmt.Print("in notification page")
	notifInstance, err := mh.Notifserv.AddNotification()
	if err != nil {
		mh.Templ.ExecuteTemplate(w,"notificationPage.html", "You have no notification")
	}
	mh.Templ.ExecuteTemplate(w,"notificationPage.html", notifInstance)
}
func (mh NotifHandler) AcceptNotification(w http.ResponseWriter, r *http.Request) {
	sender_name := r.FormValue("sender_name")
	relation := entity.Relationship{
		RelationshipId:     0,
		SenderId:           mh.Notifserv.NotifInstance.GetId(sender_name),
		RecieverId:         0,
		RelationShipStatus: 0,
	}
	err := mh.Notifserv.AcceptNotification(relation)
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
	err := mh.Notifserv.RejectNotification(relation)
	fmt.Println("reject: " + sender_name)
	if err != nil {
		mh.Templ.ExecuteTemplate(w,"notificationPage.html", "You have no notification")
	}
	mh.SeeNotification(w, r)
}
