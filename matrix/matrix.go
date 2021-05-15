package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix is a matrix of int[row][col]
type Matrix [][]int

// New use a matrix as a string input and returns either a
// matrix or an error
func New(s string) (Matrix, error) {
	rows := strings.Split(s, "\n")
	nCols := len(strings.Fields(rows[0]))
	matrix := make([][]int, 0, len(rows))
	for _, row := range rows {
		cols := strings.Fields(row)
		if len(cols) != nCols {
			return nil, errors.New("Matrix need to be the same lenght of rows")
		}
		line := make([]int, 0, nCols)
		for _, col := range cols {
			n, err := strconv.Atoi(col)
			if err != nil {
				return nil, err
			}
			line = append(line, n)
		}
		matrix = append(matrix, line)
	}

	return matrix, nil
}

// Rows returns the matrix by rows
func (m Matrix) Rows() [][]int {
	matrix := make([][]int, 0, len(m))
	for i := 0; i < len(m); i++ {
		line := make([]int, len(m[i]))
		copy(line, m[i])
		matrix = append(matrix, line)
	}

	return matrix
}

// Cols returns the matrix by columns
func (m Matrix) Cols() [][]int {
	matrix := make([][]int, 0, len(m[0]))
	for i := 0; i < len(m[0]); i++ {
		line := make([]int, len(m))
		for j := 0; j < len(m); j++ {
			line[j] = m[j][i]
		}
		matrix = append(matrix, line)
	}

	return matrix
}

// Set tries to set the matrix at position (r, c) with value and return ok-ness
func (m Matrix) Set(r int, c int, val int) bool {
	if r < 0 || c < 0 || r >= len(m) || c >= len(m[0]) {
		return false
	}
	m[r][c] = val
	return true
}
