package matrix

import (
	"reflect"
	"testing"
)

func TestMultiply(t *testing.T) {
	testcases := []struct {
		name     string
		matrices Matrices
		result   Matrix
	}{
		{
			name: "3x3 matrices",
			matrices: Matrices{
				[2]Matrix{
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
			result: Matrix{
				Rows:    3,
				Columns: 3,
				Data: [][]int{
					{6, 6, 6},
					{12, 12, 12},
					{18, 18, 18},
				},
			},
		},
		{
			name: "3x3 large matrices",
			matrices: Matrices{
				[2]Matrix{
					{
						Rows:    3,
						Columns: 3,
						Data: [][]int{
							{123, 123, 123},
							{213, 213, 213},
							{312, 312, 312},
						},
					},
					{
						Rows:    3,
						Columns: 3,
						Data: [][]int{
							{123, 123, 123},
							{213, 213, 213},
							{312, 312, 312},
						},
					},
				},
			},
			result: Matrix{
				Rows:    3,
				Columns: 3,
				Data: [][]int{
					{79704, 79704, 79704},
					{138024, 138024, 138024},
					{202176, 202176, 202176},
				},
			},
		},
		{
			name: "2x2 matrices",
			matrices: Matrices{
				[2]Matrix{
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
			result: Matrix{
				Rows:    2,
				Columns: 2,
				Data: [][]int{
					{3, 3},
					{6, 6},
				},
			},
		},
		{
			name: "2x2 matrices",
			matrices: Matrices{
				[2]Matrix{
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
			result: Matrix{
				Rows:    2,
				Columns: 2,
				Data: [][]int{
					{3, 3},
					{6, 6},
				},
			},
		},
		{
			name: "negative 2x2 matrices",
			matrices: Matrices{
				[2]Matrix{
					{
						Rows:    2,
						Columns: 2,
						Data: [][]int{
							{1, -1},
							{2, 2},
						},
					},
					{
						Rows:    2,
						Columns: 2,
						Data: [][]int{
							{1, 1},
							{2, -2},
						},
					},
				},
			},
			result: Matrix{
				Rows:    2,
				Columns: 2,
				Data: [][]int{
					{-1, 3},
					{6, -2},
				},
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.matrices.Multiply()
			if !reflect.DeepEqual(got, tt.result) {
				t.Errorf("expected: %v but got: %v", tt.result, got)
			}
		})
	}
}

func BenchmarkMultipy(b *testing.B) {
	matrices := Matrices{
		Matrices: [2]Matrix{
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
	}
	for i := 0; i < b.N; i++ {
		matrices.Multiply()
	}
}
