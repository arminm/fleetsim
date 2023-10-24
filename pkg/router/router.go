package router

import (
	"github.com/arminm/fleetsim/pkg/common"
	"github.com/arminm/fleetsim/pkg/visualizer"
)

type Route struct {
	Path []common.Position
}

type Router interface {
	GetRoute(from, to common.Position) (*Route, error)
	Draw(vis visualizer.Visualizer)
}
