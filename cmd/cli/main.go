package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/arminm/fleetsim/pkg/common"
	pb "github.com/arminm/fleetsim/pkg/server/protos"
	"github.com/manifoldco/promptui"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	cmdHello = "Say Hello!"
	cmdExit  = "Exit"

	addr = flag.String("addr", "localhost:8080", "the address to connect to")
)

func main() {
	common.LoadConfig()
	// init the CLI tool
	prompt := promptui.Select{
		Label: "Choose!",
		Items: []string{
			cmdHello,
			cmdExit,
		},
	}

	exit := false
	for {
		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch result {
		case cmdHello:
			sayHello()
		case cmdExit:
			fallthrough
		default:
			exit = true
		}

		if exit {
			break
		}
	}

	fmt.Println("bye...")
}

func sayHello() {
	fmt.Printf("What's your name? ")
	name := []byte{}
	fmt.Scanln(&name)

	conn := getGrpcClient()
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: string(name)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

func getGrpcClient() *grpc.ClientConn {
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}
