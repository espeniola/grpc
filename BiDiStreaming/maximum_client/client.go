package main

import (
	"context"
	"fmt"
	"go-grpc/BiDiStreaming/maximumpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main(){
	fmt.Printf("hello from client")

	cc, err := grpc.Dial("localhost:50051",grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect %v",err)
	}

	defer cc.Close()

	c := maximumpb.NewMaximumServiceClient(cc)

	doBiDiStreaming(c)

}

func doBiDiStreaming(c maximumpb.MaximumServiceClient)  {
	fmt.Printf("Starting BiDi Streaming RPC....\n")

	stream,err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream %v",err)
	}

	requests := []*maximumpb.MaximumRequest{&maximumpb.MaximumRequest{
		Maximum: &maximumpb.Maximum{
			Number: 4,
		},
	},
		&maximumpb.MaximumRequest{
			Maximum: &maximumpb.Maximum{
				Number: 5,
			},
		},
		&maximumpb.MaximumRequest{
			Maximum: &maximumpb.Maximum{
				Number: 2,
			},
		},
		&maximumpb.MaximumRequest{
			Maximum: &maximumpb.Maximum{
				Number: 12,
			},
		}}

	waitChannel := make(chan struct{})

	go func() {
		for _,req := range requests{
			fmt.Printf("Sending request: %v \n",req)
			stream.Send(req)
			time.Sleep(3000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for  {
			res,err := stream.Recv()
			if err == io.EOF{
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving %v",err)
				break
			}
			fmt.Printf("Received %v \n",res)
		}
		close(waitChannel)
	}()

	<-waitChannel

}