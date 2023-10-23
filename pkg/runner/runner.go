package runner

import (
	"github.com/arminm/fleetsim/pkg/simulation"
	"github.com/arminm/fleetsim/pkg/visualizer"
)

func Run() {
	size := 1000 // pixels
	s := simulation.CreateSim(&simulation.Config{
		VehicleCount: 20,
		VehicleSize:  40,
		VehicleSpeed: 10,
		Size:         size,
	})
	go s.Start()
	visualizer.Start(size, s.Vehicles)
}
