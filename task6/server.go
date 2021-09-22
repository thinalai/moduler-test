package main

import (
	"fmt"
	"net/http"
)

type PlayerServer struct{}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello wiredcraft")
}
