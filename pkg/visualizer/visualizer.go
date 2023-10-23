package visualizer

import "image/color"

type Config struct {
	Size     int
	DrawLoop func(Visualizer)
}

type Visualizer interface {
	Setup(config Config)
	Run()

	DrawLine(x1, y1, x2, y2 float64, color color.Color)
	DrawRectangle(x, y, w, h float64, heading float64, color color.Color)
	DrawTriangle(x, y, w, h float64, heading float64, color color.Color)
	DrawText(text string, x, y, size float64)
}
