package main

import (
	"GOL/matrix"
	"GOL/operators"
	"GOL/window"
	"gioui.org/app"
	"time"
)

const (
	frameIntervalInMilisec = 100
	populationDensity = 0.1
)

func main() {

	mat := matrix.NewMatrix(50, 50) // xSize = ySize
	operators.Initialize(&mat, populationDensity)

	w := window.StartWindow(&mat)

	go func() {
		for {
			operators.NextMatrix(&mat)
			w.Invalidate()
			time.Sleep(frameIntervalInMilisec * time.Millisecond)
		}
	}()

	app.Main()
}
