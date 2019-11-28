package main

import (
	"html/template"
	"net/http"
)

var templs = template.Must(template.ParseFiles("logincopy.html"))

func indexhandler(w http.ResponseWriter, r *http.Request) {
	templs.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", indexhandler)
	http.ListenAndServe(":8080", mux)
}
