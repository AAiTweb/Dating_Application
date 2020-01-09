package handler

import (
	"encoding/json"
	"net/http"

	"time"

	"log"

	"html/template"

	"github.com/betse/Dating_Application-master/entity"
	"github.com/betse/Dating_Application-master/user_profile"
)

type UserProfileHandler struct {
	tmpl        *template.Template
	userService user_profile.ProfileService
}

func NewUserProfileHandler(T *template.Template, usrService user_profile.ProfileService) *UserProfileHandler {
	return &UserProfileHandler{tmpl: T, userService: usrService}
}
func (uph *UserProfileHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	uph.tmpl.ExecuteTemplate(w, "user_profile.layout", nil)

}
func (uph *UserProfileHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := uph.userService.UsersProfile()
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return

	}
	usersProfile, err := json.MarshalIndent(users, "", "\t\t")
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return

	}
	w.Header().Set("Content-type", "application/json")
	w.Write(usersProfile)
	return

}

func (uph *UserProfileHandler) PostUser(w http.ResponseWriter, r *http.Request) {

	// if r.Method != "POST" {
	// 	w.Header().Set("Content-type", "application/json")
	// 	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	// 	return
	// }
	// fName := r.FormValue("fName")
	// lName := r.FormValue("lName")
	// country := r.FormValue("country")
	// city := r.FormValue("city")
	dob := time.Now()
	// sex := r.FormValue("sex")
	// bio := r.FormValue("bio")

	// if fName == "" || lName == "" || country == "" || city == "" || sex == "" || bio == "" {
	// 	// log.Println("not done")

	// 	w.Header().Set("Content-type", "application/json")
	// 	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	// 	return

	// }
	user := &entity.User{}
	length := r.ContentLength
	body := make([]byte, length)
	r.Body.Read(body)
	user.UserId = 1
	user.ProfPic = 1
	user.Dob = dob
	if r.Header.Get("Content-type") != " " {

		err := json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	}

	// user := &entity.User{1, 1, fName, lName, country, city, bio, sex, dob}

	user, err := uph.userService.AddProfile(user)
	if err != nil {

		log.Println("not done")
		// log.Fatal(err)
		w.Header().Set("Content-type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return

	}
	// http.Redirect(w, r, "/user/questionnarie/questions", 301)
	w.WriteHeader(http.StatusCreated)
	// http.Redirect(w, r, "#questionnarie-modal", 302)

}
