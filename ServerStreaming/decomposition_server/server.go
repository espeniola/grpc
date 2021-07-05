package main

import (
	"fmt"
	"go-grpc/ServerStreaming/decompositionpb"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct{}

func (*server) Decomposition(req *decompositionpb.DecompositionRequest,
	stream decompositionpb.DecompositionService_DecompositionServer) error{
	fmt.Println("Decomposition func was invoked %v",req)
	number := req.GetDecomposition().GetNumber()
	var x int32= 2
	for number > 1 {

		if number % x == 0{
			res := &decompositionpb.DecompositionResponse{
				Result: x,
			}
			stream.Send(res)
			number = number / x
			time.Sleep(1000 * time.Millisecond)
		}else {
			x++
		}
	}
	return nil
}

func main(){
	fmt.Println("hello from decomposition server")

	lis,err := net.Listen("tcp","0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	decompositionpb.RegisterDecompositionServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
