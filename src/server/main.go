package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"bytes"
	"encoding/json"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/hello", hello)

	log.Println("Listening...")
	http.ListenAndServe(":8080", r)
}

func StructToJSON(data interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := json.NewEncoder(buf).Encode(data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func hello(w http.ResponseWriter, _ *http.Request) {

	var data = struct {
		Title string `json:"title"`
	}{
		Title: "Golang + Angular Starter Kit",
	}

	jsonBytes, err := StructToJSON(data)
	if err != nil {
		fmt.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
	return
}
