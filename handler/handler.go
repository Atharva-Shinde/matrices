// handler package contains all the http request handling logic
package handler

import (
	"log"
	"net/http"

	"github.com/atharva-shinde/matrices/internal"
)

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	matrices, err := internal.ParseRequest(r)
	if err != nil {
		log.Printf("error: %v", err.(internal.ErrorResponse).Message)
		w.WriteHeader(err.(internal.ErrorResponse).Status)
		w.Write([]byte(err.(internal.ErrorResponse).Message))
		return
	}
	err = internal.Validate(matrices)
	if err != nil {
		log.Printf("error: %v", err.(internal.ErrorResponse).Message)
		w.WriteHeader(err.(internal.ErrorResponse).Status)
		w.Write([]byte(err.(internal.ErrorResponse).Message))
		return
	}
}
