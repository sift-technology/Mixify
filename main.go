package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/hello", func(resp http.ResponseWriter, _ *http.Request) {

		fmt.Fprint(resp, "Hello there!")
	})

	log.Println("Listening...")
	http.ListenAndServe(":8080", r)
}
