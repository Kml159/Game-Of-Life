package window

import (
	"GOL/matrix"
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

const (
	width  = 500
	height = 500
)

func StartWindow(mat *matrix.Matrix) *app.Window {

	window := new(app.Window)
	window.Option(app.Title("Game of Life"))
	window.Option(app.Size(unit.Dp(width), unit.Dp(height)))
	go func() {
		err := run(window, mat)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	return window
}

func run(window *app.Window, mat *matrix.Matrix) error {
	var ops op.Ops
	cellSize := height / float32(len(*mat))
	cellSize = cellSize * 2

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			rows := len(*mat)
			if rows == 0 {
				break
			}
			cols := len((*mat)[0])

			for i := 0; i < rows; i++ {
				for j := 0; j < cols; j++ {
					cell := (*mat)[i][j]
					var cellColor color.NRGBA
					if cell != 0 {
						cellColor = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
					} else {
						cellColor = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
					}
					x := float32(j) * cellSize
					y := float32(i) * cellSize

					paint.FillShape(gtx.Ops, cellColor, clip.Rect{
						Min: image.Pt(int(x), int(y)),
						Max: image.Pt(int(x+cellSize), int(y+cellSize)),
					}.Op())
				}
			}

			e.Frame(gtx.Ops)
		}
	}
}
