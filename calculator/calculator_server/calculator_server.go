package main

import (
	"context"
	"log"
	"net"
	"fmt"
	calculatorpb "../calculatorpb"
	"google.golang.org/grpc"
)

type server struct {}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	//this is how we get information from request
	x := req.GetSum().GetX()
	y := req.GetSum().GetY()

	result := x + y

	res := &calculatorpb.SumResponse{
		Result : result,
	}

	return res, nil
}

func main() {
	fmt.Println("Calculator")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
		
}