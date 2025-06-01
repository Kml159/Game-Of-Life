package main

import (
	"GOL/matrix"
	"GOL/operators"
	"time"
)

const (
	frameInterval = 100
)

func main() {

	mat := matrix.NewMatrix(10, 10)
	operators.Initialize(&mat, 0.6)

	for i := 0; i < 100; i++ {
		operators.NextMatrix(&mat)
		mat.Print()
		time.Sleep(frameInterval * time.Millisecond)
	}

}
