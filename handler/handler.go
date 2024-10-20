// handler package contains all the http request handling logic
package handler

import (
	"net/http"
)

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	w.Write([]byte("hello from the server"))
}
