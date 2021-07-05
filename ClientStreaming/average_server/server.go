package main

import (
	"fmt"
	"go-grpc/ClientStreaming/averagepb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type server struct{}

func (*server)CalculateAverage(stream averagepb.AverageService_CalculateAverageServer) error{
	fmt.Println("CalculateAverage func was invoked")
	var counter int32 = 0
	var total int32 = 0
	var average int32 = 0
	for  {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&averagepb.AverageResponse{
				Average: average,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream %v",err)

		}
		counter++
		number := req.GetAverage().GetNumber()
		total = total + number
		average = total / counter
	}
}

func main(){
	fmt.Println("hello from average service")

	lis,err := net.Listen("tcp","0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	averagepb.RegisterAverageServiceServer(s,&server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}

}
