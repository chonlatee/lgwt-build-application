package main

import (
	"log"
	"net/http"
)

func main() {
	store := NewInMemoryPlayerStore()
	server := &PlayerServer{store}
	log.Fatal(http.ListenAndServe(":5000", server))
}
