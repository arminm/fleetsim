package runner

import (
	"github.com/arminm/fleetsim/pkg/simulation"
	"github.com/arminm/fleetsim/pkg/visualizer"
	"github.com/arminm/fleetsim/pkg/visualizer/p5drawer"
	"github.com/arminm/fleetsim/pkg/world"
)

func Run() {
	size := 1000
	sim := simulation.CreateSim(&simulation.Config{
		Router:       world.NewSimpleWorld(size),
		VehicleCount: 5,
		VehicleSize:  40,
		VehicleSpeed: 10,
		Size:         size,
		SpeedHz:      10, //tick per second
	})
	go sim.Start()

	// visualize the simulation
	vis := p5drawer.NewP5Drawer()
	vis.Setup(visualizer.Config{
		Size:     size,
		DrawLoop: sim.Draw,
	})
	vis.Run()
}
