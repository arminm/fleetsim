package p5drawer

import (
	"image/color"

	"github.com/arminm/fleetsim/pkg/visualizer"
	"github.com/go-p5/p5"
)

type p5Drawer struct {
	size        int
	setupFunc   func()
	drawFunc    func()
	strokeWidth float64
}

func NewP5Drawer() visualizer.Visualizer {
	return &p5Drawer{
		strokeWidth: 2,
	}
}
func (pd *p5Drawer) Setup(config visualizer.Config) {
	pd.size = config.Size
	pd.setupFunc = func() {
		p5.Canvas(pd.size, pd.size)
		p5.Background(color.Gray{Y: 220})
	}
	pd.drawFunc = func() {
		config.DrawLoop(pd)
	}
}

func (pd *p5Drawer) DrawLine(x1, y1, x2, y2 float64, col color.Color) {
	p5.Push()
	p5.StrokeWidth(pd.strokeWidth)
	p5.Stroke(col)
	p5.Line(x1, y1, x2, y2)
	p5.Pop()
}

func (pd *p5Drawer) DrawRectangle(x, y, w, h float64, heading float64, col color.Color) {
	p5.Push()
	p5.StrokeWidth(pd.strokeWidth)
	p5.Fill(col)
	p5.Translate(-w/2, -h/2)
	p5.Rect(x, y, w, h)
	p5.Translate(w/2, h/2)
	p5.Pop()
}

func (pd *p5Drawer) DrawTriangle(x, y, w, h float64, heading float64, col color.Color) {
	p5.Push()
	p5.StrokeWidth(pd.strokeWidth)
	p5.Fill(col)
	p5.Translate(x, y)
	p5.Rotate(heading)
	p5.Triangle(-w/2, -h/2, w/2, -h/2, 0, h/2)
	p5.Translate(-x, -y)
	p5.Pop()
}

func (pd *p5Drawer) DrawText(text string, x, y, size float64) {
	charWidth := 10
	p5.Push()
	p5.TextSize(size)
	x = x - float64(len(text)*charWidth)/2
	y = y + size/2
	p5.Text(text, x, y)
	p5.Pop()
}

func (pd *p5Drawer) Run() {
	p5.Run(pd.setupFunc, pd.drawFunc)
}
