package matrix

import (
	"fmt"
)

type Matrix [][]int

func NewMatrix(xSize int, ySize int) Matrix {
	m := make(Matrix, ySize)
	for i := range m {
		m[i] = make([]int, xSize)
	}
	return m
}

func NewMatrixFrom(rows ...[]int) Matrix {
	m := NewMatrix(len(rows[0]), len(rows))
	for i, row := range rows {
		copy(m[i], row)
	}
	return m
}

func Mult(m1 Matrix, m2 Matrix) (Matrix, error) {
	y1 := len(m1)
	x2 := len(m2[0])
	x1 := len(m1[0])

	if x2 != y1 {
		return nil, fmt.Errorf("cannot multiply due to non-compatible dimensions: %d - %d", x2, y1)
	}

	product := NewMatrix(y1, x2)

	for i := range y1 {
		for j := range x2 {
			product[i][j] = 0

			for k := range x1 {
				product[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}

	return product, nil
}

func MultFast(m1 *Matrix, m2 *Matrix) (Matrix, error) {
	y1 := len(*m1)
	x2 := len((*m2)[0])
	x1 := len((*m1)[0])

	if x2 != y1 {
		return nil, fmt.Errorf("cannot multiply due to non-compatible dimensions: %d - %d", x2, y1)
	}

	product := NewMatrix(y1, x2)

	for i := range y1 {
		for j := range x2 {
			product[i][j] = 0

			for k := range x1 {
				product[i][j] += (*m1)[i][k] * (*m2)[k][j]
			}
		}
	}

	return product, nil
}

func (m1 Matrix) Print() {
	for i := range len(m1) {
		for j := range len(m1[0]) {
			fmt.Print(m1[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

