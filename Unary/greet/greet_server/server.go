package main

import (
	"context"
	"fmt"
	"go-grpc/Unary/greet/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error){
	fmt.Printf("Greet function was invoked with %v",req)
	firstName := req.GetGreeting().GetFirstName()

	result := "Hello " + firstName + " from service"
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res,nil
}

func main() {
	fmt.Println("hello from server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}

}
