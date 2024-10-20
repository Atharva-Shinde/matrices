package matrix

// Multiply performs multiplication of 2 matrices and returns a resultant Matrix
func (m Matrices) Multiply() Matrix {
	result := m.NewMatrix()
	matrix1 := m.Matrices[0]
	matrix2 := m.Matrices[1]
	ch := make(chan RowResult)
	for i := 0; i < len(matrix1.Data[0]); i++ {
		go helper(i, matrix1.Data[i], matrix2, ch)
		rowResult := <-ch
		for _, entity := range rowResult {
			result.Data[entity.row][entity.column] = entity.num
		}
	}
	defer close(ch)
	return *result
}

func helper(row int, m1Row []int, m2 Matrix, ch chan<- RowResult) {
	var rowResult RowResult // stores result of an entire row of the result matrix
	var result int          // stores result of a slot of the result matrix
	for i := 0; i < len(m1Row); i++ {
		for j := 0; j < m2.Columns; j++ {
			result += m1Row[j] * m2.Data[j][i]
		}
		rowResult = append(rowResult, Entity{
			num:    result,
			column: i,
			row:    row,
		})
		result = 0
	}
	ch <- rowResult
}

// Using pointers
// func (m Matrices) Multiply() ([3][3]int, error) {
// 	// fmt.Fprint(w,"Starting matrix multiplication")
// 	resultMatrix := NewMatrix()
// 	matrix1 := m[0]
// 	matrix2 := m[1]
// 	for i := 0; i < 3; i++ { // i represents current row number
// 		ch := make(chan *Matrix)
// 		go helper(i, matrix1.Data[i], matrix2.Data, ch, resultMatrix)
// 		rowResult := <-ch
// 		fmt.Println(rowResult)
// 	}
// 	// fmt.Fprint(w,"matrix multiplication completed")
// 	return resultMatrix.Data, nil
// }

// func helper(row int, m1Row [3]int, m2 [3][3]int, ch chan<- *Matrix, finalResult *Matrix) {
// 	var rowResult RowResult // stores result of an entire row of the result matrix
// 	var result int          // stores result of a slot of the result matrix
// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 3; j++ {
// 			result += m1Row[j] * m2[j][i]
// 		}
// 		rowResult = append(rowResult, Entity{
// 			num:    result,
// 			column: i,
// 			row:    row,
// 		})
// 		result = 0
// 	}
// 	for _, v := range rowResult {
// 		finalResult.Data[v.row][v.column] = v.num
// 	}
// 	ch <- finalResult
// }
