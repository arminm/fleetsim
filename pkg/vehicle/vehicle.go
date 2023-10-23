package vehicle

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/arminm/fleetsim/pkg/common"
	"github.com/arminm/fleetsim/pkg/visualizer"
	"github.com/arminm/fleetsim/pkg/world"
	"github.com/google/uuid"
)

type Config struct {
	ID       string
	Position common.Position
	Speed    int
	Size     float64
	Color    *color.RGBA
}

type Vehicle struct {
	config *Config

	ID           string
	Position     common.Position
	prevPosition common.Position
	// Speed is in Meters per second
	Speed int
	Size  float64

	Route []*world.Vertex

	Color color.Color
}

const (
	ColorAlpha   = 255
	DefaultSize  = 5  // m
	DefaultSpeed = 10 // m/s
)

var (
	red = &color.RGBA{R: 255, A: ColorAlpha}
)

func NewVehicle(config *Config) *Vehicle {
	veh := &Vehicle{
		config: config,
	}

	// set default ID to a UUID if necessary
	if config.ID == "" {
		veh.ID = uuid.New().String()
	} else {
		veh.ID = config.ID
	}

	// default position is 0,0
	veh.Position = config.Position

	// set default Speed if necessary
	if config.Speed == 0 {
		veh.Speed = DefaultSpeed
	} else {
		veh.Speed = config.Speed
	}

	// set default Size if necessary
	if config.Size == 0 {
		veh.Size = DefaultSize
	} else {
		veh.Size = config.Size
	}

	// set default Color if necessary
	if config.Color == nil {
		veh.Color = red
	} else {
		veh.Color = config.Color
	}

	return veh
}

func (veh *Vehicle) TentativeMove() {
	veh.prevPosition = veh.Position
	if len(veh.Route) > 0 {
		x1, y1 := veh.Position.Latitude, veh.Position.Longitude
		dest := veh.Route[0].Position
		x2, y2 := dest.Latitude, dest.Longitude
		dx, dy := (x2 - x1), (y2 - y1)
		if math.Abs(dx) > float64(veh.Speed) {
			dx = float64(veh.Speed) * dx / math.Abs(dx)
		}
		if math.Abs(dy) > float64(veh.Speed) {
			dy = float64(veh.Speed) * dy / math.Abs(dy)
		}
		veh.Position.Latitude = veh.Position.Latitude + dx
		veh.Position.Longitude = veh.Position.Longitude + dy
	}
}

func (veh *Vehicle) GoBack() {
	veh.Position = veh.prevPosition
}

func (veh *Vehicle) CommitMove() {
	if len(veh.Route) > 0 {
		veh.updateRoute()
	}
	veh.updateHeading()
}

func (veh *Vehicle) ChangeColor(col color.Color) {
	if col == nil {
		veh.Color = &color.RGBA{
			R: uint8(rand.Intn(256)),
			G: uint8(rand.Intn(256)),
			B: uint8(rand.Intn(256)),
			A: ColorAlpha,
		}
	} else {
		veh.Color = col
	}
}

func (veh *Vehicle) Draw(vis visualizer.Visualizer) {
	// route
	routeColor := color.RGBA{
		G: 255,
		A: 255,
	}
	if len(veh.Route) > 0 {
		x1, y1 := veh.Position.Latitude, veh.Position.Longitude
		for i := 0; i < len(veh.Route); i++ {
			nextVertex := veh.Route[i]
			x2, y2 := nextVertex.Position.Latitude, nextVertex.Position.Longitude
			vis.DrawLine(x1, y1, x2, y2, routeColor)
			x1, y1 = x2, y2
		}
	}

	// vehicle
	x, y, heading := veh.Position.Latitude, veh.Position.Longitude, veh.Position.Heading
	vis.DrawTriangle(x, y, veh.Size, veh.Size, heading, veh.Color)

	// ID label
	labelSize := veh.Size / 2
	vis.DrawText(veh.ID, x, y, labelSize)
}

func (veh *Vehicle) updateHeading() {
	if common.PositionsWithinDistance(veh.prevPosition, veh.Position, 1) {
		return // no need to update
	}

	x1, y1 := veh.prevPosition.Latitude, veh.prevPosition.Longitude
	x2, y2 := veh.Position.Latitude, veh.Position.Longitude

	veh.Position.Heading = math.Atan2(y1-y2, x2-x1) + math.Pi/2
}

func (veh *Vehicle) updateRoute() {
	maxDistanceForArrival := 1. // units
	if common.PositionsWithinDistance(veh.Position, veh.Route[0].Position, maxDistanceForArrival) {
		if len(veh.Route) > 1 {
			veh.Route = veh.Route[1:]
		} else {
			veh.Route = []*world.Vertex{}
		}
	}
}
