package main

import (
	"context"
	"fmt"
	"go-grpc/Unary/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main(){
	fmt.Println("hello im client")
	cc,err := grpc.Dial("localhost:50051",grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect %v",err)
	}

	defer cc.Close()

	c:= greetpb.NewGreetServiceClient(cc)

	doUnary(c)

}

func doUnary(c greetpb.GreetServiceClient){
	fmt.Println("Starting to do a Unary RPC...")
	req:= &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "giannis",
			LastName: "antetokounmpo",
		},
	}

	res, err := c.Greet(context.Background(),req)
	if err != nil{
		log.Fatalf("Error while calling Greet RPC: %v",err)
	}
	log.Printf("Response from Greet: %v",res)
}