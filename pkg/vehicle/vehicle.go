package vehicle

import (
	"image/color"
	"math/rand"

	"github.com/arminm/fleetsim/pkg/common"
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

func (veh *Vehicle) Move(speed int) {
	veh.prevPosition = veh.Position
	veh.Position.Latitude += (rand.Float64()*2 - 1) * float64(speed)
	veh.Position.Longitude += (rand.Float64()*2 - 1) * float64(speed)
}

func (veh *Vehicle) GoBack() {
	veh.Position = veh.prevPosition
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
