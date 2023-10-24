package world

import (
	"fmt"

	"github.com/arminm/fleetsim/pkg/common"
)

type Graph struct {
	Vertices map[string]*Vertex
}

func NewGraph() *Graph {
	return &Graph{
		Vertices: map[string]*Vertex{},
	}
}

type Vertex struct {
	ID                string
	Position          common.Position
	LanesToVertices   []*Lane
	LanesFromVertices []*Lane
}

type Lane struct {
	ID     string
	Start  *Vertex
	End    *Vertex
	Length float64
	// SpeedLimit float64
}

func (graph *Graph) NewVertex(pos common.Position) *Vertex {
	vertex := &Vertex{
		Position: pos,
	}
	graph.AddVertex(vertex)
	return vertex
}

func (graph *Graph) AddVertex(vertex *Vertex) {
	graph.Vertices[vertexKey(vertex)] = vertex
}

// adds a 2-way link
func (vertex *Vertex) AddLaneToVertex(toVertex *Vertex, length float64) {
	// way to
	laneTo := &Lane{Start: vertex, End: toVertex, Length: length}
	vertex.LanesToVertices = append(vertex.LanesToVertices, laneTo)
	toVertex.LanesFromVertices = append(toVertex.LanesFromVertices, laneTo)

	// way back
	laneBack := &Lane{Start: toVertex, End: vertex, Length: length}
	vertex.LanesFromVertices = append(vertex.LanesFromVertices, laneBack)
	toVertex.LanesToVertices = append(toVertex.LanesToVertices, laneBack)
}

func (vertex *Vertex) RemoveLane(lane *Lane) {
	vertex.LanesToVertices = common.RemoveItemByReference(vertex.LanesToVertices, lane)
	vertex.LanesFromVertices = common.RemoveItemByReference(vertex.LanesFromVertices, lane)
}

func vertexKey(vertex *Vertex) string {
	return fmt.Sprintf("%f,%f", vertex.Position.Latitude, vertex.Position.Longitude)
}
