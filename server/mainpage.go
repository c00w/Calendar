// Package to provide functions which serve files
package server

import (
	"os"
	"net.http"
	"log"
)

func generate_html(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("response.html")
	if err != nil {
		log.Fatal(err)
	}
} 

