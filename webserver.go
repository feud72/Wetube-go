package main

import (
	"log"
	"net/http"
	"webserver/app"
)

func main() {
	addr := ":3000"
	g := app.HandleListening()
	log.Fatal(http.ListenAndServe(addr, g))
}