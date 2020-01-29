package handler

import (
	//"bytes"
	"encoding/json"
	"fmt"
	"github.com/AAiTweb/Dating_Application/form"
	"github.com/AAiTweb/Dating_Application/session"
	"io"
	"math"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	_ "github.com/gorilla/mux"
	"log"

	"html/template"

	"github.com/AAiTweb/Dating_Application/entity"
	"github.com/AAiTweb/Dating_Application/user_profile"
)

type UserProfileHandler struct {
	tmpl        *template.Template
	userService user_profile.ProfileService
}

func NewUserProfileHandler(T *template.Template, usrService user_profile.ProfileService) *UserProfileHandler {
	return &UserProfileHandler{tmpl: T, userService: usrService}
}

func (uph *UserProfileHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	//log.Println(r.URL.Query().Get("id"))
	//
	//id,err:=strconv.Atoi(r.URL.Query().Get("id"))
	//if err !=nil{
	//	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	//	log.Println(err)
	//	return
	//}
	//log.Println(id)
	id := session.GetSessionData(w, r).Id
	//id := 1
	log.Println(id, "eyosis session")
	user, err := uph.userService.UserProfile(uint(id))
	//log.Println("handler")
	//log.Println(user)

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
	//headData:=entity.HeadData{
	//	Name:"user name",
	//	ProfilePicture:"new.img",
	//}
	headData := entity.HeadData{
		Name:           session.GetSessionData(w, r).Username,
		ProfilePicture: session.GetSessionData(w, r).ProfilePicture,
	}

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

	tempData := map[string]interface{}{
		"HeadData":    headData,
		"profileData": anonymousUser,
		"profilePic":  session.GetSessionData(w, r).ProfilePicture,
	}
	//log.Println(tempData)

	// fmt.Println(anonymousUser)
	//log.Println(user.FirstName, "first name")

	uph.tmpl.ExecuteTemplate(w, "user_profile.layout", tempData)
	// w.Write(userProfile)
	// return

}
func (uph *UserProfileHandler) GetFriendProfile(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	user, err := uph.userService.UserProfile(uint(id))
	//log.Println("handler")
	//log.Println(user)

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

	date := user.Dob
	now := time.Now()

	birthday := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	age := math.Floor(today.Sub(birthday).Hours() / 24 / 365)

	headData := entity.HeadData{
		Name:           session.GetSessionData(w, r).Username,
		ProfilePicture: session.GetSessionData(w, r).ProfilePicture,
	}

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

	tempData := map[string]interface{}{
		"HeadData":    headData,
		"profileData": anonymousUser,
		"profilePic":  anonymousUser.ProfPicPath[len(anonymousUser.ProfPicPath)-1],
	}
	//log.Println(tempData)

	// fmt.Println(anonymousUser)
	//log.Println(user.FirstName, "first name")

	uph.tmpl.ExecuteTemplate(w, "friend_profile.layout", tempData)
	//friend_profile.layout

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
	if r.Method != "POST" {
		log.Println("not post")
		w.Header().Set("Content-type", "application/json")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if r.Method == http.MethodPost {

		err := r.ParseForm()
		if err != nil {
			//log.Println("not post2")
			log.Fatal(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Println(r.PostForm, "POST FORM")

		profileForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		profileForm.Required("fName", "lName", "country", "city", "dob", "sex", "bio")

		//log.Println(r.FormValue("fName"),"first name")

		//data:=
		tempData := map[string]interface{}{
			"Error": "you can not leave empty field",
		}

		if !profileForm.Valid() {
			log.Println(profileForm)
			uph.tmpl.ExecuteTemplate(w, "user_form", tempData)

			return
		}
	}

	log.Println("////////////////////.........................")

	layoutISO := "2006-01-02"
	user := &entity.UserPro{}
	id, _ := strconv.Atoi(r.FormValue("hiddenId"))
	log.Println(id, "user id")
	user.UserId = uint64(id)

	//log.Println(r.FormValue("fName"),"first name")
	user.FirstName = r.FormValue("fName")
	user.LastName = r.FormValue("lName")
	user.Country = r.FormValue("country")
	user.City = r.FormValue("city")
	user.Dob, _ = time.Parse(layoutISO, r.FormValue("dob"))
	user.Sex = r.FormValue("sex")
	//log.Println(user.Sex,"user sex")
	user.Bio = r.FormValue("bio")

	user, err := uph.userService.AddProfile(user)
	if err != nil {

		log.Println("not done")
		// log.Fatal(err)
		w.Header().Set("Content-type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return

	}
	log.Println("posting user")
	// http.Redirect(w, r, "/user/questionnarie/questions", 301)
	w.WriteHeader(http.StatusCreated)
	uph.tmpl.ExecuteTemplate(w, "questionnarie", user.UserId)
	// http.Redirect(w, r, "#questionnarie-modal", 302)

}
func (uph *UserProfileHandler) PutUser(w http.ResponseWriter, r *http.Request) {
	log.Println("PUT USER")
	if r.Method != "POST" {
		log.Println("POST ERROR")
		w.Header().Set("Content-type", "application/json")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	log.Println("log passed")

	if r.Method == http.MethodPost {
		log.Println("post form")
		err := r.ParseMultipartForm((1024 * 1024 * 16))
		//err:=r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Println(r.PostForm, "POST FORM")

		profileForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		profileForm.Required("fName", "lName", "country", "city", "dob", "sex", "bio")

		log.Println(r.FormValue("fName"), "first name")

		if !profileForm.Valid() {
			log.Println(profileForm)
			http.Redirect(w, r, "/user/profile", http.StatusSeeOther)

			return
		}
	}

	////r.Header.Add("Content-Type", w.FormDataContentType())
	//bodyBuf := bytes.NewBufferString("")
	//bodyWriter := multipart.NewWriter(bodyBuf)
	//contentType := bodyWriter.FormDataContentType()
	//r.Header.Set("Content-Type", contentType)
	layoutISO := "2006-01-02"
	user := entity.UserPro{}
	user.UserId = uint64(session.GetSessionData(w, r).Id)
	//log.Println(r.FormValue("fName"),"first name")
	//user.UserId=1
	user.FirstName = r.FormValue("fName")
	user.LastName = r.FormValue("lName")
	user.Country = r.FormValue("country")
	user.City = r.FormValue("city")
	user.Dob, _ = time.Parse(layoutISO, r.FormValue("dob"))
	user.Sex = r.FormValue("sex")
	//log.Println(user.Sex,"user sex")
	user.Bio = r.FormValue("bio")
	// log.Println("not working 1")
	// log.Println(r.FormFile("profileImg"))

	mf, fh, err := r.FormFile("profileImg")
	log.Println(fh.Filename)

	// log.Println("not working 1.5")
	if err != nil {
		log.Println("not working 1.7")
		log.Fatal(err)
		return
	}
	if mf != nil {
		log.Println(fh.Filename)

		user.ProfPicPath = append(user.ProfPicPath, fh.Filename)

		if err != nil {
			log.Println(err)
			log.Println("not working 1.9")
		}

		defer mf.Close()

		writeFile(&mf, user.ProfPicPath[0])
		// log.Println("not working 2.9")

		fmt.Println(user.ProfPicPath[0])
	} else {
		user.ProfPicPath[0] = r.FormValue("image")
	}
	//log.Println(user.UserId,"user id in handler")
	// log.Println("not working 2")

	usr, err := uph.userService.UpdateProfile(&user)
	log.Println(usr, "user data")

	if err != nil {
		log.Fatal(err)
	}
	index := len(usr.ProfPicPath) - 1
	profilePic := user.ProfPicPath[index]
	claim := session.GetSessionData(w, r)
	if claim != nil {
		session.RenewSession(claim.Username, claim.Id, profilePic, w)
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
func (uph *UserProfileHandler) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	//id:=uint(session.GetSessionData(w,r).Id)
	id, _ := strconv.Atoi(r.FormValue("hiddenId"))
	id2 := uint(id)
	id2, err := uph.userService.DeleteProfile(id2)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	uph.tmpl.ExecuteTemplate(w, "user_form", nil)

}
