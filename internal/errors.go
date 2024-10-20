package internal

import (
	"fmt"
	"net/http"
)

// ErrorResponse is a struct for handling custom Error
type ErrorResponse struct {
	Status  int
	Message string
}

func (e ErrorResponse) Error() string {
	if e.Status == 0 || e.Message == "" {
		return fmt.Sprintf("status: %d, error: %v", http.StatusInternalServerError, "internal server error")
	}
	return fmt.Sprintf("status: %d, error: %v", e.Status, e.Message)
}
