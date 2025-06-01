package main

import (
	"GOL/matrix"
	"GOL/operators"
	"GOL/window"
	"gioui.org/app"
	"time"
)

const (
	frameInterval = 100
)

func main() {

	mat := matrix.NewMatrix(50, 50) // xSize = ySize
	operators.Initialize(&mat, 0.1)

	w := window.StartWindow(&mat)

	time.Sleep(1000 * time.Millisecond)

	go func() {
		for {
			operators.NextMatrix(&mat)
			w.Invalidate()
			time.Sleep(frameInterval * time.Millisecond)
		}
	}()

	app.Main()
}
