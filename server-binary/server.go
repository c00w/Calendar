package main 

import (
	"net/http"
	"fmt"
	"log"
)

func defaulthandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "No such page")
}

func main() {
	http.HandleFunc("/foo", defaulthandle)
	log.Fatal(http.ListenAndServe(":80", nil))
}
