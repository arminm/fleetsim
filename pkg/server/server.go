package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/arminm/fleetsim/pkg/server/protos"
	"google.golang.org/grpc"
)

// server is used to implement server.GreeterServer.
type simServer struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements server.GreeterServer
func (s *simServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func Run(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &simServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
