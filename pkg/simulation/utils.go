package simulation

import (
	"math"

	"github.com/arminm/fleetsim/pkg/vehicle"
)

func areColliding(veh *vehicle.Vehicle, veh2 *vehicle.Vehicle) bool {
	latDiff := math.Abs(veh2.Position.Latitude - veh.Position.Latitude)
	lonDiff := math.Abs(veh2.Position.Longitude - veh.Position.Longitude)
	return latDiff < veh.Size && lonDiff < veh.Size
}
