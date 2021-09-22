package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/players/score", &PlayerServer{})
	log.Fatal(http.ListenAndServe(":3000", mux))
}
