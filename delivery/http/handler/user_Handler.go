package handler

import (
	"fmt"
	"github.com/AAiTweb/Dating_Application/entity"
	"github.com/AAiTweb/Dating_Application/form"
	"github.com/AAiTweb/Dating_Application/session"
	"github.com/AAiTweb/Dating_Application/user/service"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

type UserHandler struct {
	Templ    template.Template
	Uservice service.UserService
}

func NewUserHandler(servc *service.UserService, temp *template.Template) *UserHandler {
	return &UserHandler{
		Templ:    *temp,
		Uservice: *servc,
	}
}


func (mh UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	mh.Templ.ExecuteTemplate(w, "logincopy.html", nil)
}

func (mh UserHandler) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	mh.Templ.ExecuteTemplate(w, "confirm_email.html", nil)
}
func (mh UserHandler) Passwordreset(w http.ResponseWriter, r *http.Request) {
	mh.Templ.ExecuteTemplate(w, "resetpassword.html", nil)
}
func (mh UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		Fusername := r.FormValue("f_username")
		Femail := r.FormValue("f_email")
		Fpassword := r.FormValue("f_password")
		Repassword := r.FormValue("confirm_f_password")
		time := time.Now().String()
		confToken := time + Fusername

		if Fpassword != Repassword {
			mh.Templ.ExecuteTemplate(w, "logincopy.html", "Passwords doesn't match, re-enter password")
			return
		}
		if form.ValidatePassword(Fpassword) != "" {
			mh.Templ.ExecuteTemplate(w, "logincopy.html", form.ValidatePassword(Fpassword))
			return
		}

		if len(Fusername) < 5 {
			println("short username")
			mh.Templ.ExecuteTemplate(w, "logincopy.html", "Username must be longer than 5 characters")
			return
		}
		fmt.Println(Fusername)
		fmt.Println(Femail, Fpassword)
		HashedPassword,err := form.HashPassword(Fpassword)
		if err != nil{
			log.Print("Hashing failed")
		}
		user := entity.User{
			Id:                0,
			UserName:          Fusername,
			Password:          HashedPassword,
			Email:             Femail,
			ConfirmationToken: confToken,
		}
		handlerserivce := mh.Uservice
		err = handlerserivce.RegisterUser(user)
		if err != nil {
			fmt.Println("the error is ", err)
			mh.Templ.ExecuteTemplate(w, "logincopy.html", "Unstable conn or already used username or email")
			return
		} else {
			mh.Templ.ExecuteTemplate(w, "CheckEmail.html", nil)
			return
		}
	}
}

func (mh UserHandler) Validatelogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("Lusername")
		password := r.FormValue("Lpassword")
		fmt.Println(username, password)
		user := entity.User{
			Id:                0,
			UserName:          username,
			Password:          password,
			Email:             "",
			ConfirmationToken: "",
		}
		QuestionerFilled,err := mh.Uservice.QueFilled(user)
		if err != nil{
			fmt.Println("Cant go there")
			mh.Templ.ExecuteTemplate(w,"logincopy.html","Login failed, try again")
			return
		}
		if !QuestionerFilled {
			http.Redirect(w,r,"/user/questionnarie", http.StatusSeeOther)
			return
		}
		userId, UserName, profilePic, err := mh.Uservice.CheckLogin(user)
		if err != nil {
			mh.Templ.ExecuteTemplate(w,"logincopy.html","incorrect username or password")
			fmt.Println("user info not filled inside")
			return
		} else {
			fmt.Println("input: " + password + " saved: " + password)
			println("Correct username and password")

			tokenString, err := session.Generate(userId, UserName, profilePic)
			if err != nil {
				return
			}
			http.SetCookie(w, &http.Cookie{
				Name:  "token",
				Value: tokenString,
			})
			path := "/preload?id="+strconv.Itoa(userId)
			http.Redirect(w,r,path,http.StatusSeeOther)
		}

	}
}
func (mh UserHandler) ValidateSignup(w http.ResponseWriter, r *http.Request) {
	Vkey := r.URL.Query().Get("conftok")
	if !(len(Vkey) > 0) {
		log.Println("verification key not found")
		return
	}
	user := entity.User{
		Id:                0,
		UserName:          "",
		Password:          "",
		Email:             "",
		ConfirmationToken: Vkey,
	}
	err := mh.Uservice.ValidateToken(user)
	if err != nil {
		_ = mh.Templ.ExecuteTemplate(w, "logincopy.html", "Login failed, try again")
	}
	http.Redirect(w,r,"/user/questionnarie",http.StatusSeeOther)
}

func (mh UserHandler) ConfirmEmail(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodPost{
		email := r.FormValue("reset_email")
		err := mh.Uservice.Checkemail(email)
		fmt.Println(email)
		if err != nil{
			if err.Error() == "1"{
				mh.Templ.ExecuteTemplate(w,"confirm_email.html","Incorrect email address try again")
				return
			}
			if err.Error() == "2"{
				mh.Templ.ExecuteTemplate(w,"confirm_email.html","Unstable connection try again")
				return
			} else{
				mh.Templ.ExecuteTemplate(w,"confirm_email.html","Unstable connection try again")
				return
			}
		}
		mh.Templ.ExecuteTemplate(w,"CheckEmail.html",nil)
	}
}

func (mh UserHandler) ConfirmReset(w http.ResponseWriter,r *http.Request) {
	vKey := r.URL.Query().Get("conftok")
	if !(len(vKey) > 0) {
		log.Println("verification key not found")
		return
	}
	fmt.Println(vKey)
	err := mh.Uservice.ConfirmReset(vKey)
	if err != nil {
		fmt.Println(err)
		_ = mh.Templ.ExecuteTemplate(w, "logincopy.html", "Password reset failed")
		return
	}
	username, password, err := mh.Uservice.GetUser(vKey)
	if err != nil {
		fmt.Println(err)
		_ = mh.Templ.ExecuteTemplate(w, "logincopy.html", "Password reset failed")
		return
	}
	user := entity.User{
		Id:                0,
		UserName:          username,
		Password:          password,
		Email:             "",
		ConfirmationToken: "",
	}
	fmt.Println(username, " ", password)
	userId, UserName, profilePic, err1 := mh.Uservice.CheckReset(user)
	if err1 != nil{
		_ = mh.Templ.ExecuteTemplate(w, "logincopy.html", "User doesn't exist")
		return
	}
	fmt.Println(userId, UserName, profilePic, err)
	tokenString, err2 := session.Generate(userId, UserName, profilePic)
	if err2 != nil {
		log.Println(err2)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: tokenString,
	})
	println("executing template...")
	http.Redirect(w,r,"/resetpassword",http.StatusPermanentRedirect)
}

func (mh UserHandler) ResetPassword(w http.ResponseWriter,r *http.Request){
	if r.Method == http.MethodPost{
		newpassword := r.FormValue("new_password")
		confPassword := r.FormValue("new_password_conf")

		if newpassword != confPassword {
			mh.Templ.ExecuteTemplate(w, "logincopy.html", "Passwords doesn't match, re-enter password")
			return
		}
		if form.ValidatePassword(newpassword) != "" {
			mh.Templ.ExecuteTemplate(w, "resetpassword.html", form.ValidatePassword(newpassword))
			return
		}
		data := session.GetSessionData(w,r)
		HashedPassword,err := form.HashPassword(newpassword)
		if err != nil {
			mh.Templ.ExecuteTemplate(w,"logincopy.html","password reset failed")
		}
		log.Println(HashedPassword)
		err = mh.Uservice.ResetPassword(data.Id,HashedPassword)
		if err != nil {
			mh.Templ.ExecuteTemplate(w,"logincopy.html","password reset failed")
		}
		http.Redirect(w,r,"/login",http.StatusPermanentRedirect)
	}
}

func (mh UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	session.RemoveSession(w)
	http.Redirect(w,r,"/login",http.StatusSeeOther)

}
