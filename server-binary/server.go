package main 

import (
	"net/http"
	"fmt"
	"log"
	"calendar/server"
)

func defaulthandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "No such page")
}

func main() {
	http.HandleFunc("/foo", defaulthandle)
	http.HandleFunc("/add_entry", server.AddEntry)
	http.HandleFunc("/", server.Mainpage)
	http.HandleFunc("/date.js", server.DateJavascript)
	http.HandleFunc("/main.js", server.Javascript)
	log.Fatal(http.ListenAndServe(":80", nil))
}
