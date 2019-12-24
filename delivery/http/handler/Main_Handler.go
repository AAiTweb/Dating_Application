package handler

import (
	"github.com/Eyosi-G/Dating_Application/user"
	"html/template"
)

type MainHandler struct {
	Templ *template.Template
	Uservice user.UserService

}

