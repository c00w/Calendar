// Package to provide functions which serve files
package server

import (
	"net/http"
	"log"
	"html/template"
)

func Mainpage(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("/home/ubuntu/work/src/calendar/templates/frontpage.html")
	if err != nil {
		log.Fatal(err)
	}
	html.ExecuteTemplate(w, "frontpage.html", nil) 
} 

