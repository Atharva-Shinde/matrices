package main

import (
	"log"
	"net/http"

	"github.com/atharva-shinde/matrices/handler"
)

func main() {
	http.HandleFunc("POST /", handler.SubmitHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
