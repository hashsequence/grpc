package main

import (
	"context"
	"fmt"
	"log"
	"google.golang.org/grpc"
	findMax "../pb"
	"io"
	"time"
)

func main() {
	fmt.Println("Hello I'm the client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()
	
	c:= findMax.NewFindMaxServiceClient(cc)

	doBiDiStreaming(c)
	
}



func doBiDiStreaming(c findMax.FindMaxServiceClient) {
	fmt.Println("starting BiDi Streaming RPC")
	stream, err := c.ComputeMax(context.Background()) 
	if err != nil {
		log.Fatalf("error while calling GreetEveryone RPC: %v", err)
		return
	}

	requests := []*findMax.FindMaxRequest{
		&findMax.FindMaxRequest {
			Val : 1,
		},
		&findMax.FindMaxRequest {
			Val : 5,
		},
		&findMax.FindMaxRequest {
			Val : 3,
		},
		&findMax.FindMaxRequest {
			Val : 6,
		},
		&findMax.FindMaxRequest {
			Val : 2,
		},
		&findMax.FindMaxRequest {
			Val : 20,
		},
	}

	waitc := make(chan struct{})
	go func() {
		//send a bunch of messages
		for _, req := range requests {
			fmt.Printf("Sending req %v\n", req)
			stream.Send(req)
		time.Sleep(100 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break;
			}
			if err != nil {
				log.Fatalf("error while recieving %v", err)
				break;
			}
			fmt.Printf("Recieved %v \n", res.GetMax())
		}
		close(waitc)
	}()

	<-waitc

}