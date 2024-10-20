// matrix package contains the logic and configurations required for matrices computation
package matrix

// Matrix represents a single matrix
type Matrix struct {
	Rows    int     `json:"rows,omitempty"`
	Columns int     `json:"columns,omitempty"`
	Data    [][]int `json:"data"`
}

// Matrices is an array of two matrices
type Matrices struct {
	Matrices [2]Matrix `json:"matrices"`
}

// NewMatrix initialises a pointer to Matrix
func (m Matrices) NewMatrix() *Matrix {
	rows := m.Matrices[0].Rows
	columns := m.Matrices[1].Columns
	data := make([][]int, rows)
	for i := range data {
		data[i] = make([]int, columns)
	}
	return &Matrix{
		Rows:    rows,
		Columns: columns,
		Data:    data, // make([][]int{},rows) doesn't work here, it only initiates the rows len and cap leaving columns still  nil; creating an array doesn't work either as it doesn't support dynamic variable assignment
	}
}
