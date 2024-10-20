package handler

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/atharva-shinde/matrices/internal"
)

func TestSubmitHandler(t *testing.T) {
	testcases := []struct {
		name        string
		requestBody []byte
		expected    string
		expectedErr error
	}{
		{
			name: "test1",
			requestBody: []byte(`{
				"matrices": [
					{
						"rows": 3,
						"columns": 3,
						"data": [
							[
								1,
								1,
								1
							],
							[
								2,
								2,
								2
							],
							[
								3,
								3,
								3
							]
						]
					},
					{
						"rows": 3,
						"columns": 3,
						"data": [
							[
								1,
								1,
								1
							],
							[
								2,
								2,
								2
							],
							[
								3,
								3,
								3
							]
						]
					}
				]
			}`),
			expected:    `{"rows":3,"columns":3,"data":[[6,6,6],[12,12,12],[18,18,18]]}` + "\n",
			expectedErr: nil,
		},
		{
			name: "test2",
			requestBody: []byte(`{
				"matrices": [
					{
						"rows": 2,
						"columns": 3,
						"data": [
							[
								1,
								1,
								1
							],
							[
								2,
								2,
								2
							],
							[
								3,
								3,
								3
							]
						]
					},
					{
						"rows": 3,
						"columns": 3,
						"data": [
							[
								1,
								1,
								1
							],
							[
								2,
								2,
								2
							],
							[
								3,
								3,
								3
							]
						]
					}
				]
			}`),
			expected:    "",
			expectedErr: errors.New("invalid request, must provide valid row length"),
		},
		{
			name: "test3",
			requestBody: []byte(`{
				"matrices": [
					{
						"rows": 3,dskjlsd
						"columns": 3,
						"data": [
							[
								1,
								1,
								1
							],
							[
								2,
								2,
								2
							],
							[
								3,
								3,
								3
							]
						]
					},
					{
						"rows": 3,
						"columns": 3,
						"data": [
							[
								1,
								1,
								1
							],
							[
								2,
								2,
								2
							],
							[
								3,
								3,
								3
							]
						]
					}
				]
			}`),
			expected:    "",
			expectedErr: errors.New("invalid json request"),
		},
	}
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			mockRequest := http.Request{
				Header: http.Header{"Content-Type": []string{"application/json"}},
				Body:   io.NopCloser(bytes.NewReader(tt.requestBody)),
			}
			w := httptest.NewRecorder()
			SubmitHandler(w, &mockRequest)
			response := w.Result()
			responseBody, _ := io.ReadAll(response.Body)
			if tt.expectedErr != nil {
				if string(responseBody) != tt.expectedErr.Error() {
					t.Errorf("expected error: %v but got: %v", tt.expectedErr.(internal.ErrorResponse).Message, string(responseBody))
				}
			} else {
				if string(responseBody) != tt.expected {
					t.Errorf("expected: %v but got: %v", tt.expected, string(responseBody))
				}
			}

		})
	}
}
