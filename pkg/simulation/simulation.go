package simulation

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/arminm/fleetsim/pkg/common"
	"github.com/arminm/fleetsim/pkg/vehicle"
)

type Config struct {
	VehicleCount int
	VehicleSize  float64
	VehicleSpeed int
	Size         int
}

type Simulation struct {
	config   *Config
	Vehicles []*vehicle.Vehicle
}

func CreateSim(config *Config) *Simulation {
	sim := &Simulation{config: config}
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
	return vehicle.NewVehicle(&vehicle.Config{
		ID:       id,
		Position: sim.randomVehiclePosition(vehicleSize / 2),
		Speed:    vehicleSpeed,
		Size:     vehicleSize,
	})
}

func (sim *Simulation) Start() {
	println("Start sim!")
	for {
		sim.tick()
		time.Sleep(time.Millisecond * 100)
	}
}

func (sim *Simulation) tick() {
	for i := 0; i < len(sim.Vehicles); i++ {
		veh := sim.Vehicles[i]
		veh.Move(10)
		if sim.hasCollision(veh) {
			veh.ChangeColor(nil)
			veh.GoBack()
		}
		if sim.isOutOfBounds(veh) {
			veh.GoBack()
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
	outOfLatitudeBounds := veh.Position.Latitude < veh.Size/2 || veh.Position.Latitude > float64(sim.config.Size)-veh.Size/2
	outOfLongitudeBounds := veh.Position.Longitude < veh.Size/2 || veh.Position.Longitude > float64(sim.config.Size)-veh.Size/2
	return outOfLatitudeBounds || outOfLongitudeBounds
}

func (sim *Simulation) randomVehiclePosition(padding float64) common.Position {
	pos := common.Position{
		Latitude:  padding + float64(rand.Intn(sim.config.Size-int(padding))),
		Longitude: padding + float64(rand.Intn(sim.config.Size-int(padding)*2)),
	}
	return pos
}
