package main

import (
	"context"
	"fmt"
	"go-grpc/Unary/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (*server) Calculate(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse,error){
	fmt.Printf("Calculate function was invoked with %v",req)
	firstNumber := req.GetCalculator().GetFirstNumber()
	secondNumber := req.GetCalculator().GetSecondNumber()

	result := firstNumber + secondNumber

	res := &calculatorpb.CalculatorResponse{
		Result: result,
	}
	return res,nil

}

func main(){
	fmt.Println("hello from calculator server")

	lis,err := net.Listen("tcp","0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}

}
