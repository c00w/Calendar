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
	http.HandleFunc("/", server.Mainpage)
	log.Fatal(http.ListenAndServe(":80", nil))
}
