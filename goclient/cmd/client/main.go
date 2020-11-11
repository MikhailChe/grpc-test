package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/mikhailche/grpc-test/src/main/proto"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:8080"
	defaultName = "world"
)

func main() {
	fmt.Println("Going to dial a grpc address")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(fmt.Sprintf("did not connect %v", err))
	}

	defer conn.Close()
	fmt.Println("Creating service client")
	c := pb.NewGreetingServiceClient(conn)

	fmt.Println("Creating context")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println("Calling reomte procedure")
	r, err := c.Greeting(ctx, &pb.HelloRequest{Firstname: "Mikhail", Lastname: "Chernoskutov", Age: 28})

	if err != nil {
		panic(fmt.Sprintf("did not connect %v", err))
	}
	fmt.Printf("Greetings: %v", r.GetGreeting())
}
