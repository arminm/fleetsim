package main

import (
	"flag"

	"github.com/arminm/fleetsim/pkg/server"
)

var (
	port = flag.Int("port", 8080, "The server port")
)

func main() {
	flag.Parse()
	server.Run(*port)
}
