package main

import (
	"github.com/feud72/Wetube-go/app"
	"log"
	"net/http"
)

func main() {
	addr := ":3000"
	g := app.HandleListening()
	log.Fatal(http.ListenAndServe(addr, g))
}
