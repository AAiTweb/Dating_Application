package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"time"

	"log"

	"html/template"

	"github.com/biniyam112/TheDatingApp/Dating_Application/entity"
	"github.com/biniyam112/TheDatingApp/Dating_Application/user_profile"
	v2 "github.com/liamylian/jsontime/v2"
)

type UserProfileHandler struct {
	tmpl        *template.Template
	userService user_profile.ProfileService
}

func NewUserProfileHandler(T *template.Template, usrService user_profile.ProfileService) *UserProfileHandler {
	return &UserProfileHandler{tmpl: T, userService: usrService}
}

func (uph *UserProfileHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := uph.userService.UserProfile(1)

	if err != nil {
		w.Header().Set("Content-type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		log.Println(err)
		return
	}
	// userProfile, err := json.MarshalIndent(user, "", "\t\t")
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		log.Println(err)
		return

	}

	// log.Println(user)
	// w.Header().Set("Content-type", "application/json")
	// w.Write(userProfile)
	// log.Println(user[0].Dob)
	date := user.Dob
	now := time.Now()

	birthday := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	// birthday := time.Date(1989, 1, 2, 0, 0, 0, 0, time.UTC)
	// today := time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)
	// today := time.Now()

	age := math.Floor(today.Sub(birthday).Hours() / 24 / 365)
	anonymousUser := struct {
		ProfPicPath []string
		FirstName   string
		UserId      uint64

		LastName string
		Country  string
		City     string
		Bio      string
		Sex      string
		Age      float64
	}{
		ProfPicPath: user.ProfPicPath,
		FirstName:   user.FirstName,
		Country:     user.Country,
		LastName:    user.LastName,
		Bio:         user.Bio,
		Sex:         user.Sex,
		Age:         age,
		UserId:      user.UserId,
		City:        user.City}

	// fmt.Println(anonymousUser)
	log.Println(user.FirstName, "first name")

	uph.tmpl.ExecuteTemplate(w, "user_profile.layout", anonymousUser)
	// w.Write(userProfile)
	// return

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
	log.Println("post")
	// if r.Method != "POST" {
	// 	w.Header().Set("Content-type", "application/json")
	// 	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	// 	return
	// }
	// fName := r.FormValue("fName")
	// lName := r.FormValue("lName")
	// country := r.FormValue("country")
	// city := r.FormValue("city")
	// dob := time.Now()
	// sex := r.FormValue("sex")
	// bio := r.FormValue("bio")

	// if fName == "" || lName == "" || country == "" || city == "" || sex == "" || bio == "" {
	// 	// log.Println("not done")

	// 	w.Header().Set("Content-type", "application/json")
	// 	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	// 	return

	// }
	user := &entity.UserPro{}
	length := r.ContentLength
	body := make([]byte, length)
	r.Body.Read(body)
	user.UserId = 1
	// user.ProfPic = 1
	// user.Dob = dob
	layoutISO := "2006-01-02"
	v2.SetDefaultTimeFormat(layoutISO, time.Local)
	var js = v2.ConfigWithCustomTimeFormat
	if r.Header.Get("Content-type") != " " {

		err := js.Unmarshal(body, &user)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	}

	// user := &entity.UserPro{1, 1, fName, lName, country, city, bio, sex, dob}

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
func (uph *UserProfileHandler) PutUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("POST ERROR")
		w.Header().Set("Content-type", "application/json")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// r.Header.Add("Content-Type", w.FormDataContentType())
	layoutISO := "2006-01-02"
	user := entity.UserPro{}
	user.UserId, _ = strconv.ParseUint(r.FormValue("id"), 10, 64)
	user.FirstName = r.FormValue("fName")
	user.LastName = r.FormValue("lName")
	user.Country = r.FormValue("country")
	user.City = r.FormValue("city")
	user.Dob, _ = time.Parse(layoutISO, r.FormValue("dob"))
	user.Sex = r.FormValue("sex")
	user.Bio = r.FormValue("bio")
	// log.Println("not working 1")
	// log.Println(r.FormFile("profileImg"))

	mf, fh, err := r.FormFile("profileImg")
	// log.Println("not working 1.5")
	if err != nil {
		log.Println("not working 1.7")
		log.Println(err)
		return
	}
	if mf != nil {
		log.Println(fh.Filename)

		user.ProfPicPath = append(user.ProfPicPath, fh.Filename)
		log.Println("not working 1.8")

		if err != nil {
			log.Println(err)
			// log.Println("not working 1.9")
		}

		defer mf.Close()

		writeFile(&mf, user.ProfPicPath[0])
		// log.Println("not working 2.9")

		fmt.Println(user.ProfPicPath[0])
	} else {
		user.ProfPicPath[0] = r.FormValue("image")
	}
	// log.Println("not working 2")
	_, err = uph.userService.UpdateProfile(&user)

	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/user/profile", http.StatusSeeOther)

}
func writeFile(mf *multipart.File, fname string) {

	wd, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}

	path := filepath.Join(wd, "/../../ui", "assets", "images", fname)
	image, err := os.Create(path)

	if err != nil {
		log.Println(err)
	}
	defer image.Close()
	io.Copy(image, *mf)
}
