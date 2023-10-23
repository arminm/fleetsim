package world

import (
	"image/color"

	"github.com/arminm/fleetsim/pkg/visualizer"
)

type Config struct {
}

type World struct {
	config      *Config
	RoadNetwork *Graph
}

func NewWorld(config *Config) *World {
	w := &World{
		config:      config,
		RoadNetwork: NewGraph(),
	}

	return w
}

func (sw *World) Draw(vis visualizer.Visualizer) {
	laneColor := &color.RGBA{
		B: 255, A: 255,
	}
	vertexColor := laneColor
	for _, ver := range sw.RoadNetwork.Vertices {
		x1, y1 := ver.Position.Latitude, ver.Position.Longitude
		vis.DrawRectangle(x1, y1, 5, 5, 0, vertexColor)
		for _, lane := range ver.LanesToVertices {
			x2, y2 := lane.End.Position.Latitude, lane.End.Position.Longitude
			vis.DrawLine(x1, y1, x2, y2, laneColor)
		}
	}
}
