package runner

import "github.com/arminm/fleetsim/pkg/sim"

func Run() {
	s := sim.CreateSim()
	s.Start()
}
