package main

import (
	"github.com/feud72/Wetube-go/v1"
	"log"
	"net/http"
)

func main() {
	addr := ":3000"
	g := v1.BaseEndpoint()
	log.Fatal(http.ListenAndServe(addr, g))
}
