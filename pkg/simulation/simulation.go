package simulation

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/arminm/fleetsim/pkg/common"
	"github.com/arminm/fleetsim/pkg/router"
	"github.com/arminm/fleetsim/pkg/vehicle"
	"github.com/arminm/fleetsim/pkg/visualizer"
)

type Config struct {
	Router       router.Router
	VehicleCount int
	VehicleSize  float64
	VehicleSpeed int
	Size         int
	SpeedHz      int
}

type Simulation struct {
	config   *Config
	Router   router.Router
	Vehicles []*vehicle.Vehicle
}

func CreateSim(config *Config) *Simulation {
	sim := &Simulation{config: config}

	// assign router
	sim.Router = config.Router

	// add vehicles
	for i := 0; i < config.VehicleCount; i++ {
		id := fmt.Sprintf("%d", i)
		veh := sim.placeNewVehicle(id)
		for sim.hasCollision(veh) {
			veh = sim.placeNewVehicle(id)
		}
		sim.Vehicles = append(sim.Vehicles, veh)
	}

	return sim
}

func (sim *Simulation) placeNewVehicle(id string) *vehicle.Vehicle {
	vehicleSize := sim.config.VehicleSize
	if vehicleSize == 0 {
		vehicleSize = vehicle.DefaultSize
	}
	vehicleSpeed := sim.config.VehicleSpeed
	if vehicleSpeed == 0 {
		vehicleSpeed = vehicle.DefaultSpeed
	}

	randPos := sim.randomPosition()
	veh := vehicle.NewVehicle(&vehicle.Config{
		ID:       id,
		Position: randPos,
		Speed:    vehicleSpeed,
		Size:     vehicleSize,
	})

	route, err := sim.Router.GetRoute(randPos, sim.randomPosition())
	if route == nil || err != nil {
		fmt.Println("failed to generate route")
	} else {
		veh.Route = route
	}

	return veh
}

func (sim *Simulation) Start() {
	println("Start sim!")
	for {
		sim.tick()
		time.Sleep(time.Second / time.Duration(sim.config.SpeedHz))
	}
}

func (sim *Simulation) tick() {
	for i := 0; i < len(sim.Vehicles); i++ {
		veh := sim.Vehicles[i]
		if len(veh.Route.Path) == 0 {
			route, err := sim.Router.GetRoute(veh.Position, sim.randomPosition())
			if route == nil || err != nil {
				fmt.Println("failed to generate route")
			} else {
				veh.Route = route
			}
		}
		veh.TentativeMove()
		if sim.isOutOfBounds(veh) || sim.hasCollision(veh) {
			veh.ChangeColor(nil)
			veh.GoBack()
		} else {
			veh.CommitMove()
		}
	}
}

func (sim *Simulation) hasCollision(veh *vehicle.Vehicle) bool {
	for i := 0; i < len(sim.Vehicles); i++ {
		veh2 := sim.Vehicles[i]
		if veh.ID == veh2.ID {
			continue
		}

		if areColliding(veh, veh2) {
			return true
		}
	}
	return false
}

func (sim *Simulation) isOutOfBounds(veh *vehicle.Vehicle) bool {
	outOfLatitudeBounds := veh.Position.Latitude < veh.Size || veh.Position.Latitude > float64(sim.config.Size)-veh.Size
	outOfLongitudeBounds := veh.Position.Longitude < veh.Size || veh.Position.Longitude > float64(sim.config.Size)-veh.Size
	return outOfLatitudeBounds || outOfLongitudeBounds
}

func (sim *Simulation) randomPosition() common.Position {
	lat := float64(sim.config.Size/4) + rand.Float64()*float64(sim.config.Size/2)
	lon := float64(sim.config.Size/4) + rand.Float64()*float64(sim.config.Size/2)
	return common.Position{Latitude: lat, Longitude: lon}
}

func (sim *Simulation) Draw(vis visualizer.Visualizer) {
	sim.Router.Draw(vis)
	for _, veh := range sim.Vehicles {
		veh.Draw(vis)
	}
}
