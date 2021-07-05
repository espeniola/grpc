package main

import (
	"fmt"
	"go-grpc/BiDiStreaming/maximumpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"time"
)

type server struct{}


func GetMax(slice []int32) (int32) {
	var max int32 = slice[0]
	for _, value := range slice {
		if max < value {
			max = value
		}
	}
	return max
}

func (s *server) FindMaximum(stream maximumpb.MaximumService_FindMaximumServer) error{
	fmt.Printf("FindMaximum was invoked")
	var slice []int32
	for  {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream %v",err)
		}
		number := req.GetMaximum().GetNumber()
		fmt.Printf("Recived request %v\n",number)
		slice = append(slice,number)
		max := GetMax(slice)
		sErr := stream.Send(&maximumpb.MaximumResponse{
			Max: max,
		})
		time.Sleep(5000 * time.Millisecond)

		if sErr != nil {
			log.Fatalf("Error while sending data to client %v",err)
		}

	}
}

func main(){
	fmt.Printf("server started")
	lis , err := net.Listen("tcp","0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	maximumpb.RegisterMaximumServiceServer(s,&server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
