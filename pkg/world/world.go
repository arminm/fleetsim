package world

import (
	"image/color"

	"github.com/arminm/fleetsim/pkg/common"
	"github.com/arminm/fleetsim/pkg/router"
	"github.com/arminm/fleetsim/pkg/visualizer"
)

type Config struct {
}

type World struct {
	config      *Config
	RoadNetwork *Graph
}

func NewWorld(config *Config) *World {
	w := &World{
		config:      config,
		RoadNetwork: NewGraph(),
	}

	return w
}

func (sw *World) GetRoute(from, to common.Position) (*router.Route, error) {
	fromVertex := sw.findClosestVertex(from)
	toVertex := sw.findClosestVertex(to)
	path, err := sw.findPath(fromVertex, toVertex)
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

func (sw *World) findClosestVertex(pos common.Position) *Vertex {
	var minDistance float64
	var closestVertex *Vertex
	for _, ver := range sw.RoadNetwork.Vertices {
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

func (sw *World) findPath(start, end *Vertex) ([]*Vertex, error) {
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

func (sw *World) Draw(vis visualizer.Visualizer) {
	laneColor := &color.RGBA{
		B: 255, A: 255,
	}
	vertexColor := laneColor
	for _, ver := range sw.RoadNetwork.Vertices {
		x1, y1 := ver.Position.Latitude, ver.Position.Longitude
		vis.DrawRectangle(x1, y1, 5, 5, 0, vertexColor)
		for _, lane := range ver.LanesToVertices {
			x2, y2 := lane.End.Position.Latitude, lane.End.Position.Longitude
			vis.DrawLine(x1, y1, x2, y2, laneColor)
		}
	}
}
