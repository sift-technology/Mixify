package main

import (
	"fmt"
	"log"
	"net/http"
	"src/server/utils"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/hello", hello)

	log.Println("Listening...")
	http.ListenAndServe(":8080", r)
}

func hello(resp http.ResponseWriter, _ *http.Request) {

	var data = struct {
		Title string `json:"title"`
	}{
		Title: "Golang + Angular Starter Kit",
	}

	jsonBytes, err := utils.StructToJSON(data)
	if err != nil {
		fmt.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
	return
}
