package operators

import (
	"GOL/matrix"
	"GOL/util"
)

type Index struct {
	i int
	j int
}

const (
	upperLimit = 100000
)

func Initialize(mat *matrix.Matrix, threshold float64) {
	for i := range len(*mat) {
		for j := range len((*mat)[0]) {
			if util.RandomFloatBetween(0, 1) <= float64(threshold) {
				(*mat)[i][j] = util.RandomIntBetween(1, upperLimit)
			}
		}
	}
}

func Neighbors(mat *matrix.Matrix, i int, j int) []Index {
	neighbors := make([]Index, 0, 8)
	rows := len(*mat)
	cols := len((*mat)[0])

	dirs := []Index{
		{i: -1, j: -1}, {i: -1, j: 0}, {i: -1, j: 1},
		{i: 0, j: -1}, {i: 0, j: 1},
		{i: 1, j: -1}, {i: 1, j: 0}, {i: 1, j: 1},
	}

	for _, d := range dirs {
		ni, nj := i+d.i, j+d.j
		if ni >= 0 && ni < rows && nj >= 0 && nj < cols {
			neighbors = append(neighbors, Index{i: ni, j: nj})
		}
	}

	return neighbors
}

func LiveNeighbors(mat *matrix.Matrix, i int, j int) []Index {
	neighbors := Neighbors(mat, i, j)
	liveNeighbors := make([]Index, 0, 8)
	for _, n := range neighbors {
		if (*mat)[n.i][n.j] == 1 {
			liveNeighbors = append(liveNeighbors, n)
		}
	}
	return liveNeighbors
}

func LiveNeighborsAmount(mat *matrix.Matrix, i int, j int) int {
	neighbors := Neighbors(mat, i, j)
	liveNeighbors := 0
	for _, n := range neighbors {
		if (*mat)[n.i][n.j] != 0 {
			liveNeighbors++
		}
	}
	return liveNeighbors
}

func Next(mat *matrix.Matrix, i int, j int) int {

	liveNeighbors := LiveNeighborsAmount(mat, i, j)
	if (*mat)[i][j] != 0 { // alive

		if liveNeighbors < 2 { // Any live cell with fewer than two live neighbors dies, as if by underpopulation.
			return 0
		} else if liveNeighbors == 2 || liveNeighbors == 3 { // Any live cell with two or three live neighbors lives on to the next generation.
			return (*mat)[i][j]
		} else if liveNeighbors > 3 { // Any live cell with more than three live neighbors dies, as if by overpopulation.
			return 0
		}

	} else { // dead

		// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
		if liveNeighbors == 3 {
			return util.RandomIntBetween(1, upperLimit)
		}
	}
	return (*mat)[i][j]
}

func NextMatrix(mat *matrix.Matrix) {
	nextMat := (*mat)
	for i := range len(*mat) {
		for j := range len((*mat)[0]) {
			nextMat[i][j] = Next(mat, i, j)
		}
	}
	mat.CopyFrom(&nextMat)
}
