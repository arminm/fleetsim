package world

import "github.com/arminm/fleetsim/pkg/common"

type Graph struct {
	Vertices []*Vertex
}

func NewGraph() *Graph {
	return &Graph{
		Vertices: []*Vertex{},
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
	graph.Vertices = append(graph.Vertices, vertex)
}

func (vertex *Vertex) AddLaneToVertex(toVertex *Vertex, length float64) *Lane {
	lane := &Lane{Start: vertex, End: toVertex, Length: length}
	vertex.LanesToVertices = append(vertex.LanesToVertices, lane)
	toVertex.LanesFromVertices = append(vertex.LanesFromVertices, lane)
	return lane
}

func (vertex *Vertex) RemoveLane(lane *Lane) {
	vertex.LanesToVertices = common.RemoveItemByReference(vertex.LanesToVertices, lane)
	vertex.LanesFromVertices = common.RemoveItemByReference(vertex.LanesFromVertices, lane)
}
