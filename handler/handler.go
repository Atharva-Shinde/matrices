// handler package contains all the http request handling logic
package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/atharva-shinde/matrices/internal"
	"github.com/atharva-shinde/matrices/internal/matrix"
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
	result := operate(matrices)
	json.NewEncoder(w).Encode(result)
}

func operate(o matrix.Operator) matrix.Matrix {
	result := o.Multiply()
	return result
}
