package internal

import (
	"encoding/json"
	"net/http"

	"github.com/atharva-shinde/matrices/internal/matrix"
)

// checks if the parsed input is valid before the actual computation
func Validate(m matrix.Matrices) error {
	if m.Matrices[0].Rows == 0 || m.Matrices[0].Columns == 0 || m.Matrices[0].Data == nil || m.Matrices[1].Rows == 0 || m.Matrices[1].Columns == 0 || m.Matrices[1].Data == nil {
		return ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "incomplete request",
		}
	}
	// checks whether rows of a matrix are consistent in length
	for _, matrix := range m.Matrices {
		rowLength := len(matrix.Data[0]) // caveat: deliberately not checking for Rows as it may not accurately depict actual length of a row
		for _, row := range matrix.Data {
			if rowLength != len(row) {
				return ErrorResponse{
					Status:  http.StatusBadRequest,
					Message: "invalid request, must provide valid matrices",
				}
			}
		}
	}
	// checks "Rows" integer is consitent with "Data"
	// TODO: create a check for columns
	for _, matrix := range m.Matrices {
		if len(matrix.Data) != matrix.Rows {
			return ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "invalid request, must provide valid row length",
			}
		}
	}
	return nil
}

// Unmarshals the JSON request from the client and returns a Calculation interface
func ParseRequest(r *http.Request) (matrix.Matrices, error) {
	var matrices matrix.Matrices
	err := json.NewDecoder(r.Body).Decode(&matrices)
	if err != nil {
		switch err {
		case err.(*json.SyntaxError): // TODO: handle err.(*json.UnmarshalTypeError)
			return matrix.Matrices{}, ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "invalid json request",
			}
			// TODO: default case: handle internal server error
		}
	}
	return matrices, nil
}
