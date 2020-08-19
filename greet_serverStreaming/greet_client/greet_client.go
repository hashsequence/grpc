package main

import (
	"context"
	"fmt"
	"log"
	"google.golang.org/grpc"
	greetpb "../greetpb"
	"io"
)

func main() {
	fmt.Println("Hello I'm the client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()
	
	c:= greetpb.NewGreetServiceClient(cc)
	//doUnary(c)
	doServerStreaming(c)
	
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting unary rpc")
	req := &greetpb.GreetRequest{
		Greeting : &greetpb.Greeting {
			FirstName : "Stephanie",
			LastName : "Wong",
		},
	}
	//fmt.Printf("Created client %f", c)

	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting server streaming rpc")
	req := &greetpb.GreetManytimesRequest {
		Greeting: &greetpb.Greeting {
			FirstName : "Avery",
			LastName : "Wong",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			//we reach end of stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}
}