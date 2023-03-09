package main

import (
	"net/http"

	"github.com/aarti2626/Mixify/src/server/api"
	"github.com/aarti2626/Mixify/src/server/spotify"
)

func main() {
	spotify.Authenticate()
	srv := api.NewServer()

	http.ListenAndServe(":8080", srv)

}
