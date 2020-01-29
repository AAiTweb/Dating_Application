package main

import (
	"bytes"
	"fmt"
	"github.com/AAiTweb/Dating_Application/delivery/http/handler"
	"github.com/AAiTweb/Dating_Application/entity"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strings"

	"github.com/AAiTweb/Dating_Application/user_profile/service"
	"io/ioutil"
	"net/http/httptest"

	 "github.com/AAiTweb/Dating_Application/user_profile/repository"
	"net/http"
	"testing"
	"html/template"
	//"net/url"
)

func TestUserProfile(t *testing.T){
	tmpl := template.Must(template.ParseGlob("../../ui/template/*.html"))
	userProfileRepo:=repository.NewMockUserProfileRepo(nil)
	userProfileService:=service.NewUserProfileServiceImpl(userProfileRepo)
	userProfileHandler:=handler.NewUserProfileHandler(tmpl,userProfileService)


	mux:=http.NewServeMux()
	mux.HandleFunc("/user/profile",userProfileHandler.GetUser)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/user/profile")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}


	if !bytes.Contains(body, []byte("mock user name")) {
		t.Errorf("want body to contain %q", body)
	}

}
func TestUsersProfile(t *testing.T){
	tmpl := template.Must(template.ParseGlob("../../ui/template/*.html"))
	userProfileRepo:=repository.NewMockUserProfileRepo(nil)
	userProfileService:=service.NewUserProfileServiceImpl(userProfileRepo)
	userProfileHandler:=handler.NewUserProfileHandler(tmpl,userProfileService)


	mux:=http.NewServeMux()
	mux.HandleFunc("/users/profile",userProfileHandler.GetUsers)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/users/profile")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("mock user name")) {
		t.Errorf("want body to contain %q", body)
	}
}
func TestUserProfilesUpdate(t *testing.T){
	tmpl := template.Must(template.ParseGlob("../../ui/template/*.html"))
	userProfileRepo:=repository.NewMockUserProfileRepo(nil)
	userProfileService:=service.NewUserProfileServiceImpl(userProfileRepo)
	userProfileHandler:=handler.NewUserProfileHandler(tmpl,userProfileService)


	mux:=http.NewServeMux()
	mux.HandleFunc("/user/update",userProfileHandler.PutUser)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	surl := ts.URL
	//form := url.Values{}

	//"fName","lName","country","city","dob","sex","bio"

	//form.Add("fName",entity.UserProfileMock.FirstName)
	//form.Add("lName",entity.UserProfileMock.LastName)
	//form.Add("country",entity.UserProfileMock.Country)
	//form.Add("city",entity.UserProfileMock.City)
	//form.Add("dob",entity.UserProfileMock.Dob.String())
	//form.Add("sex",entity.UserProfileMock.Sex)
	//form.Add("bio",entity.UserProfileMock.Bio)
	//form.Add("profileImg","/home/betse/go/src/github.com/AAiTweb/Dating_Application/ui/assets/chat/images/default.jpg")
	values := map[string]io.Reader{
		 // lets assume its this file
		"fName":strings.NewReader(entity.UserProfileMock.FirstName),
		"lName":strings.NewReader(entity.UserProfileMock.LastName),
		"country":strings.NewReader(entity.UserProfileMock.Country),
		"city":strings.NewReader(entity.UserProfileMock.City),
		"dob":strings.NewReader(entity.UserProfileMock.Dob.String()),
		"sex":strings.NewReader(entity.UserProfileMock.Sex),
		"bio":strings.NewReader(entity.UserProfileMock.Bio),

	}
	var b bytes.Buffer
	var fw io.Writer
	var err error
	w := multipart.NewWriter(&b)
	for key, r := range values {
		//var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		w, err := w.CreateFormField(key)
		if  err != nil {
			log.Fatal(err)
		// Add an image file
		}
		if wr, err := io.Copy(w, r); err != nil {
			log.Fatal(err,wr)
		}

	}


	//file := mustOpen(fileName)
	file, err := os.Open("p1ic7.jpg")
	if err != nil {
		pwd, _ := os.Getwd()
		fmt.Println("PWD: ", pwd)
		panic(err)
	}

	if fw, err = w.CreateFormFile("profileImg", file.Name()); err != nil {
		t.Errorf("Error creating writer: %v", err)
	}
	if _, err = io.Copy(fw, file); err != nil {
		t.Errorf("Error with io.Copy: %v", err)
	}
	w.Close()


	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", surl+"/user/update", &b)
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp,err:=tc.Do(req)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	//body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	//if !bytes.Contains(body, []byte("mock user name")) {
	//	t.Errorf("want body to contain %q", body)
	//}


	if err != nil {
		t.Fatal(err)
	}

}
