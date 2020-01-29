package main

import (
	"bytes"
	"io/ioutil"

	//"bytes"
	"github.com/AAiTweb/Dating_Application/delivery/http/handler"
	"github.com/AAiTweb/Dating_Application/questionnarie/repository"
	"github.com/AAiTweb/Dating_Application/questionnarie/service"
	"html/template"
	//"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestQuestionnaireGet(t *testing.T){

	tmpl := template.Must(template.ParseGlob("../../ui/template/*.html"))
	questionnaireRepo :=repository.NewMockQuesRepo(nil)
	questionnaireService:=service.NewQuestionnaireServiceImpl(questionnaireRepo)
	questionnaireHandler:=handler.NewUserQuestionnarieHandler(tmpl,questionnaireService)
	mux:=http.NewServeMux()
	mux.HandleFunc("/user/questionnarie/questions",questionnaireHandler.Questionnaire)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()
	tc := ts.Client()
	url := ts.URL
	resp, err := tc.Get(url + "/user/questionnarie/questions")
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

	if !bytes.Contains(body, []byte("user question 1")) {
		t.Errorf("want body to contain %q", body)
	}


}
//func TestQuestionnairesPost(t *testing.T){
//
//}