package main

import (
	"net/http"
	"os"

	"github.com/freeformz/googlegoauth"
)

func main() {
	g := &googlegoauth.Handler{
		RequireDomain: "heroku.com",
		Key:           os.Getenv("KEY"),
		ClientID:      os.Getenv("CLIENT_ID"),
		ClientSecret:  os.Getenv("CLIENT_SECRET"),
	}

	http.ListenAndServe(":"+os.Getenv("PORT"), g)
}
