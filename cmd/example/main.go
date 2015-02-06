package main

import (
	"net/http"
	"os"

	"github.com/freeformz/googlegoauth"
)

func main() {

	g := &googlegoauth.Handler{
		RequireDomain: os.Getenv("REQUIRE_DOMAIN"),
		Key:           os.Getenv("KEY"),
		ClientID:      os.Getenv("CLIENT_ID"),
		ClientSecret:  os.Getenv("CLIENT_SECRET"),
		Handler:       http.FileServer(http.Dir(".")),
	}

	http.ListenAndServe(":"+os.Getenv("PORT"), g)
}
