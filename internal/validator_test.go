package internal

import (
	"bytes"
	"io"
	"net/http"
	"reflect"
	"testing"

	"github.com/atharva-shinde/matrices/internal/matrix"
)

func TestParseRequest(t *testing.T) {
	testcases := []struct {
		name        string
		requestBody []byte
		expected    matrix.Matrices
		expectedErr error
	}{
		{
			name: "valid json #1",
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
			expected: matrix.Matrices{
				Matrices: [2]matrix.Matrix{
					{
						Rows:    3,
						Columns: 3,
						Data: [][]int{
							{1, 1, 1},
							{2, 2, 2},
							{3, 3, 3},
						},
					},
					{
						Rows:    3,
						Columns: 3,
						Data: [][]int{
							{1, 1, 1},
							{2, 2, 2},
							{3, 3, 3},
						},
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "valid json #2",
			requestBody: []byte(`{
    "matrices": [
        {
            "rows": 2,
            "columns": 2,
            "data": [
                [
                    1,
                    1
                ],
                [
                    2,
                    2
                ]
            ]
        },
        {
            "rows": 2,
            "columns": 2,
            "data": [
                [
                    1,
                    1
                ],
                [
                    2,
                    2
                ]
            ]
        }
    ]
}`),
			expected: matrix.Matrices{
				Matrices: [2]matrix.Matrix{
					{
						Rows:    2,
						Columns: 2,
						Data: [][]int{
							{1, 1},
							{2, 2},
						},
					},
					{
						Rows:    2,
						Columns: 2,
						Data: [][]int{
							{1, 1},
							{2, 2},
						},
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "invalid json",
			requestBody: []byte(`{
"matrices": [
	{
	"rows": 2,
	"columns": 2,
	"data": [
		[1, 1],
		[2, 2],
	]
	},
	{
	"rows: 2,
	"columns": 2,
	"data": [
		[1, 1],tskkk
		[2, 2],
	]
	}
]
}`),
			expected: matrix.Matrices{},
			expectedErr: ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "invalid json request",
			},
		},
	}
	for _, tt := range testcases {
		mockRequest := http.Request{
			Header: http.Header{"Content-Type": []string{"application/json"}},
			Body:   io.NopCloser(bytes.NewReader(tt.requestBody)),
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseRequest(&mockRequest)
			if err != tt.expectedErr {
				t.Errorf("expected error: %v but got: %v", tt.expectedErr, err)
			} else if !reflect.DeepEqual(tt.expected, got) {
				t.Errorf("expected: %v but got: %v", tt.expected, got)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	testcases := []struct {
		name        string
		matrices    matrix.Matrices
		expectedErr error
	}{
		{
			name: "No error",
			matrices: matrix.Matrices{
				Matrices: [2]matrix.Matrix{
					{
						Rows:    3,
						Columns: 3,
						Data: [][]int{
							{1, 1, 1},
							{2, 2, 2},
							{3, 3, 3},
						},
					},
					{
						Rows:    3,
						Columns: 3,
						Data: [][]int{
							{1, 1, 1},
							{2, 2, 2},
							{3, 3, 3},
						},
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "null matrix",
			matrices: matrix.Matrices{
				Matrices: [2]matrix.Matrix{
					{
						Rows:    3,
						Columns: 3,
						Data: [][]int{
							{1, 1, 1},
							{2, 2, 2},
							{3, 3, 3},
						},
					},
					{},
				},
			},
			expectedErr: ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "incomplete request",
			},
		},
		{
			name: "row length is 0",
			matrices: matrix.Matrices{
				Matrices: [2]matrix.Matrix{
					{
						Rows:    0,
						Columns: 3,
						Data: [][]int{
							{1, 1, 1},
							{2, 2, 2},
							{3, 3, 3},
						},
					},
					{
						Rows:    3,
						Columns: 3,
						Data: [][]int{
							{1, 1, 1},
							{2, 2, 2},
							{3, 3, 3},
						},
					},
				},
			},
			expectedErr: ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "incomplete request",
			},
		},
		{
			name: "Inconsistent Data",
			matrices: matrix.Matrices{
				Matrices: [2]matrix.Matrix{
					{
						Rows:    3,
						Columns: 3,
						Data: [][]int{
							{1, 1},
							{2, 2, 2},
							{3, 3, 3},
						},
					},
					{
						Rows:    3,
						Columns: 3,
						Data: [][]int{
							{1, 1, 1},
							{2, 2, 2},
							{3, 3, 3},
						},
					},
				},
			},
			expectedErr: ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "invalid request, must provide valid matrices",
			},
		},
		{
			name: "Rows object is not consistent with Data",
			matrices: matrix.Matrices{
				Matrices: [2]matrix.Matrix{
					{
						Rows:    2,
						Columns: 3,
						Data: [][]int{
							{1, 1, 1},
							{2, 2, 2},
							{3, 3, 3},
						},
					},
					{
						Rows:    3,
						Columns: 3,
						Data: [][]int{
							{1, 1, 1},
							{2, 2, 2},
							{3, 3, 3},
						},
					},
				},
			},
			expectedErr: ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "invalid request, must provide valid row length",
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := Validate(tt.matrices)
			if gotErr != tt.expectedErr {
				t.Errorf("expected error: %v but got: %v", tt.expectedErr, gotErr)
			}
		})
	}
}
