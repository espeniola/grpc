package main

import (
	"context"
	"fmt"
	"go-grpc/Unary/calculator/calculatorpb"
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

	c:= calculatorpb.NewCalculatorServiceClient(cc)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient){
	fmt.Println("Starting to do Unary RPC...")
	req:= &calculatorpb.CalculatorRequest{
		Calculator: &calculatorpb.Calculator{
			FirstNumber: 3,
			SecondNumber: 10,
		},
	}
	res,err := c.Calculate(context.Background(),req)
	if err != nil{
		log.Fatalf("Error while calling Greet RPC: %v",err)
	}
	log.Printf("Response from Greet: %v",res)
}
