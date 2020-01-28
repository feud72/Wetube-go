package main

import (
	"log"
	"net/http"

	v1 "github.com/feud72/Wetube-go/v1"
)

func main() {
	addr := ":3000"
	app := v1.Engine()
	log.Fatal(http.ListenAndServe(addr, app))
}
