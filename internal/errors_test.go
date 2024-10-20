package internal

import (
	"net/http"
	"testing"
)

func TestError(t *testing.T) {
	testcases := []struct {
		name        string
		err         ErrorResponse
		expectedErr string
	}{
		{
			name: "test1",
			err: ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: "internal server error",
			},
			expectedErr: "status: 500, error: internal server error",
		},
		{
			name: "test2",
			err: ErrorResponse{
				Status:  0,
				Message: "invalid json request",
			},
			expectedErr: "status: 500, error: internal server error",
		},
		{
			name: "test3",
			err: ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "",
			},
			expectedErr: "status: 500, error: internal server error",
		},
	}
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := tt.err.Error()
			if gotErr != tt.expectedErr {
				t.Errorf("expected error: %v but got: %v", tt.expectedErr, gotErr)
			}
		})
	}
}
