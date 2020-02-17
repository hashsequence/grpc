package main

import (
	"context"
	"fmt"
	"log"
	"google.golang.org/grpc"
	calculatorpb "../calculatorpb"
)

func main() {
	fmt.Println("Hello I'm the client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()
	
	c:= calculatorpb.NewCalculatorServiceClient(cc)
	doUnary(c)
	
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting unary rpc adding 7 and -34")
	req := &calculatorpb.SumRequest{
		Sum : &calculatorpb.Sum {
			X : 7,
			Y : -34,
		},
	}
	//fmt.Printf("Created client %f", c)

	res, err := c.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling calculatorpb's Sum RPC: %v", err)
	}
	log.Printf("Response from Sum %d", res.Result)
}