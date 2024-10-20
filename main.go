package main

import (
	"log"
	"net/http"

	"github.com/atharva-shinde/matrices/handler"
)

func main() {
	http.HandleFunc("GET /", handler.SubmitHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
