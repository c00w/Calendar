// Package to provide functions which serve files
package server

import (
	"net/http"
	"log"
	"html/template"
)

func AddEntry(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if (err != nil) {
		log.Print(err)
		return
	}
	event := make(map[string]string, 3)
	event["date"] = r.Form["date"][0]
	event["time"] = r.Form["time"][0]
	event["item"] = r.Form["item"][0]
	StoreEvent(event)
}

func DateJavascript(w http.ResponseWriter, r *http.Request) {
	w.Header()["Content-Type"] = []string{"text/javascript"}
	html, err := template.ParseFiles("/home/ubuntu/work/src/calendar/templates/date-en-US.js")
	if err != nil {
		log.Fatal(err)
	}
	html.ExecuteTemplate(w, "frontpage.js", nil) 

}

func Javascript(w http.ResponseWriter, r *http.Request) {
	w.Header()["Content-Type"] = []string{"text/javascript"}
	html, err := template.ParseFiles("/home/ubuntu/work/src/calendar/templates/frontpage.js")
	if err != nil {
		log.Fatal(err)
	}
	html.ExecuteTemplate(w, "frontpage.js", nil) 

}

func Mainpage(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("/home/ubuntu/work/src/calendar/templates/frontpage.html")
	if err != nil {
		log.Fatal(err)
	}
	html.ExecuteTemplate(w, "frontpage.html", nil) 
} 


