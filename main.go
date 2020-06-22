package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/games", gameHandlers)

	err := http.ListenAndServe(":3000", nil)

	log.Fatal(err)
}
