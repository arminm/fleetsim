package world

import (
	"github.com/arminm/fleetsim/pkg/common"
)

const (
	defaultLaneLength = 1
)

func NewSimpleWorld(size int) *World {
	sw := NewWorld(&Config{})
	nodeGrid := [][]*Vertex{}
	scaleFactor := 100
	gridSize := size / scaleFactor
	padding := scaleFactor / 2
	for i := 0; i < gridSize; i++ {
		row := []*Vertex{}
		nodeGrid = append(nodeGrid, row)
		for j := 0; j < gridSize; j++ {
			v := sw.RoadNetwork.NewVertex(common.Position{
				Latitude:  float64(padding + i*scaleFactor),
				Longitude: float64(padding + j*scaleFactor),
				Heading:   0,
			})
			nodeGrid[i] = append(nodeGrid[i], v)
			if i > 0 {
				v.AddLaneToVertex(nodeGrid[i-1][j], defaultLaneLength)
			}
			if j > 0 {
				v.AddLaneToVertex(nodeGrid[i][j-1], defaultLaneLength)
			}
		}
	}
	return sw
}
