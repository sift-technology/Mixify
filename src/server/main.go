package main

import (
	"net/http"

	"github.com/aarti2626/Mixify/src/server/api"
)

func main() {

	srv := api.NewServer()

	http.ListenAndServe(":8080", srv)

}
