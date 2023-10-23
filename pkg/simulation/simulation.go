package simulation

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/arminm/fleetsim/pkg/vehicle"
	"github.com/arminm/fleetsim/pkg/visualizer"
	"github.com/arminm/fleetsim/pkg/world"
)

type Config struct {
	VehicleCount int
	VehicleSize  float64
	VehicleSpeed int
	Size         int
	SpeedHz      int
}

type Simulation struct {
	config   *Config
	World    *world.World
	Vehicles []*vehicle.Vehicle
}

func CreateSim(config *Config) *Simulation {
	sim := &Simulation{config: config}

	// create world
	sim.World = world.NewSimpleWorld(config.Size)

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

	randVertex := sim.randomVertex()
	veh := vehicle.NewVehicle(&vehicle.Config{
		ID:       id,
		Position: randVertex.Position,
		Speed:    vehicleSpeed,
		Size:     vehicleSize,
	})

	if len(randVertex.LanesToVertices) > 0 {
		route := []*world.Vertex{}
		link := randVertex.LanesToVertices[0]
		for link != nil {
			route = append(route, link.End)
			if len(link.End.LanesToVertices) > 0 {
				link = link.End.LanesToVertices[0]
			} else {
				link = nil
			}
		}
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

func (sim *Simulation) randomVertex() *world.Vertex {
	vertices := sim.World.RoadNetwork.Vertices
	return vertices[rand.Intn(len(vertices))]
}

func (sim *Simulation) Draw(vis visualizer.Visualizer) {
	sim.World.Draw(vis)
	for _, veh := range sim.Vehicles {
		veh.Draw(vis)
	}
}
