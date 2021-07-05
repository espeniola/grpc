package main

import (
	"context"
	"fmt"
	"go-grpc/ServerStreaming/decompositionpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main(){
	fmt.Println("hello from client")
	cc,err := grpc.Dial("localhost:50051",grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect %v",err)
	}

	defer cc.Close()

	c:= decompositionpb.NewDecompositionServiceClient(cc)

	doServerStreamin(c)
}

func doServerStreamin(c decompositionpb.DecompositionServiceClient){
	fmt.Println("Starting to do Server Streaming RPC...")
	req:= &decompositionpb.DecompositionRequest{
		Decomposition: &decompositionpb.Decomposition{
			Number: 120,
		},
	}

	res,err := c.Decomposition(context.Background(),req)
	if err != nil{
		log.Fatalf("Error while calling Greet RPC: %v",err)
	}

	for  {
		msg,err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil{
			log.Fatalf("Error while reading stream %v",err)
		}
		log.Printf("Response from server: %v",msg.GetResult())
	}

}
