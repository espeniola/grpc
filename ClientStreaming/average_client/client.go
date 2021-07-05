package main

import (
	"context"
	"fmt"
	"go-grpc/ClientStreaming/averagepb"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main(){
	fmt.Println("this is client")
	cc,err := grpc.Dial("localhost:50051",grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect %v",err)
	}

	defer cc.Close()

	c := averagepb.NewAverageServiceClient(cc)

	doClientStreaming(c)

}

func doClientStreaming(c averagepb.AverageServiceClient){
	fmt.Println("Starting to do Client Streaming RPC...")

	requests := []*averagepb.AverageRequest{&averagepb.AverageRequest{
		Average: &averagepb.Average{
			Number: 1,
		},
	},&averagepb.AverageRequest{
		Average: &averagepb.Average{
			Number: 2,
		},
	},&averagepb.AverageRequest{
		Average: &averagepb.Average{
			Number: 3,
		},
	},&averagepb.AverageRequest{
		Average: &averagepb.Average{
			Number: 4,
		},
	},&averagepb.AverageRequest{
		Average: &averagepb.Average{
			Number: 5,
		},
	}}
	stream, err := c.CalculateAverage(context.Background())
	if err != nil {
		log.Fatalf("error while calling calculate average service %v",err)
	}

	for _, req := range requests {
		fmt.Printf("sending request %v\n",req)
		stream.Send(req)
		time.Sleep(5000 * time.Millisecond)
	}

	res , err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("fatal error while receiving resp from service %v",err)
	}
	fmt.Printf("response: %v",res)

}
