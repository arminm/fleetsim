package visualizer

import (
	"image/color"

	"github.com/arminm/fleetsim/pkg/vehicle"
	"github.com/go-p5/p5"
)

type visualizer struct {
	size     int
	vehicles []*vehicle.Vehicle
}

func Start(size int, vehList []*vehicle.Vehicle) {
	vis := &visualizer{
		vehicles: vehList,
		size:     size,
	}
	p5.Run(vis.setup, vis.draw)
}

func (vis *visualizer) setup() {
	p5.Canvas(vis.size, vis.size)
	p5.Background(color.Gray{Y: 220})
}

func (vis *visualizer) draw() {
	for i := 0; i < len(vis.vehicles); i++ {
		veh := vis.vehicles[i]

		// vehicle
		p5.StrokeWidth(2)
		p5.Fill(veh.Color)
		x, y := veh.Position.Latitude, veh.Position.Longitude
		p5.Ellipse(x, y, veh.Size, veh.Size)

		// ID label
		labelSize := veh.Size / 2
		// labelX, labelY := x-veh.Size/2, y+veh.Size/2
		labelX, labelY := x-veh.Size/4., y+labelSize/2
		p5.TextSize(labelSize)
		p5.Text(veh.ID, labelX, labelY)
	}

	// p5.Fill(color.RGBA{B: 255, A: 208})
	// p5.Quad(50, 50, 80, 50, 80, 120, 60, 120)

	// p5.Fill(color.RGBA{G: 255, A: 208})
	// p5.Triangle(100, 100, 120, 120, 80, 120)

	// p5.Stroke(color.Black)
	// p5.StrokeWidth(5)
	// p5.Arc(300, 100, 80, 20, 0, 1.5*math.Pi)
}
