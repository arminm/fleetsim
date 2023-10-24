package simplerouter

import (
	"image/color"

	"github.com/arminm/fleetsim/pkg/common"
	"github.com/arminm/fleetsim/pkg/router"
	"github.com/arminm/fleetsim/pkg/visualizer"
)

type Config struct {
	Size int
}

type SimplerRouter struct {
	config      *Config
	RoadNetwork *Graph
}

const (
	defaultLaneLength = 1
)

func NewSimpleRouter(config *Config) *SimplerRouter {
	sr := &SimplerRouter{
		config:      &Config{},
		RoadNetwork: NewGraph(),
	}
	nodeGrid := [][]*Vertex{}
	scaleFactor := 100
	gridSize := config.Size / scaleFactor
	padding := scaleFactor / 2
	for i := 0; i < gridSize; i++ {
		row := []*Vertex{}
		nodeGrid = append(nodeGrid, row)
		for j := 0; j < gridSize; j++ {
			v := sr.RoadNetwork.NewVertex(common.Position{
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
	return sr
}

func (sr *SimplerRouter) GetRoute(from, to common.Position) (*router.Route, error) {
	fromVertex := sr.findClosestVertex(from)
	toVertex := sr.findClosestVertex(to)
	path, err := sr.findPath(fromVertex, toVertex)
	if err != nil {
		return nil, err
	}

	pathPositions := []common.Position{}
	for _, ver := range path {
		pathPositions = append(pathPositions, ver.Position)
	}
	return &router.Route{
		Path: pathPositions,
	}, nil
}

func (sr *SimplerRouter) findClosestVertex(pos common.Position) *Vertex {
	var minDistance float64
	var closestVertex *Vertex
	for _, ver := range sr.RoadNetwork.Vertices {
		if closestVertex == nil {
			closestVertex = ver
			minDistance = common.Distance(pos, ver.Position)
			continue
		}

		d := common.Distance(ver.Position, pos)
		if d < minDistance {
			minDistance = d
			closestVertex = ver
		}
	}

	return closestVertex
}

func (sr *SimplerRouter) findPath(start, end *Vertex) ([]*Vertex, error) {
	visited := map[*Vertex]bool{}
	queue := [][]*Vertex{{start}}
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		vertex := path[len(path)-1]
		if visited[vertex] {
			continue
		}
		for _, lane := range vertex.LanesToVertices {
			newPath := make([]*Vertex, len(path))
			copy(newPath, path)
			newPath = append(newPath, lane.End)
			if lane.End == end {
				return newPath, nil
			}
			queue = append(queue, newPath)
		}
		visited[vertex] = true
	}
	return nil, nil

}

func (sr *SimplerRouter) Draw(vis visualizer.Visualizer) {
	laneColor := &color.RGBA{
		B: 255, A: 255,
	}
	vertexColor := laneColor
	for _, ver := range sr.RoadNetwork.Vertices {
		x1, y1 := ver.Position.Latitude, ver.Position.Longitude
		vis.DrawRectangle(x1, y1, 5, 5, 0, vertexColor)
		for _, lane := range ver.LanesToVertices {
			x2, y2 := lane.End.Position.Latitude, lane.End.Position.Longitude
			vis.DrawLine(x1, y1, x2, y2, laneColor)
		}
	}
}
