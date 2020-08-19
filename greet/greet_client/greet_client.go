package main

import (
	"context"
	"fmt"
	"log"
	"google.golang.org/grpc"
	greetpb "../greetpb"
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
	
	c:= greetpb.NewGreetServiceClient(cc)
	//doUnary(c)
	//doServerStreaming(c)
	//doClientStreaming(c)
	doBiDiStreaming(c)
	
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

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting client streaming rpc")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest {
			Greeting: &greetpb.Greeting {
				FirstName : "Avery",
				LastName : "Wong",
			},
		},
		&greetpb.LongGreetRequest {
			Greeting: &greetpb.Greeting {
				FirstName : "dan",
				LastName : "Wan",
			},
		},
		&greetpb.LongGreetRequest {
			Greeting: &greetpb.Greeting {
				FirstName : "max",
				LastName : "Lee",
			},
		},
		&greetpb.LongGreetRequest {
			Greeting: &greetpb.Greeting {
				FirstName : "Bane",
				LastName : "Anderson",
			},
		},
	}
	stream, err := c.LongGreet(context.Background()) 
	if err != nil {
		log.Fatalf("error while calling LongGreet RPC: %v", err)
	}

	for _, req := range requests {
		fmt.Println("Sending req %v", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while recieving response from LongGreet &v", err)
	}
	fmt.Printf("LongGreet Response: %v\n",res) 
}

func doBiDiStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("starting BiDi Streaming RPC")
	stream, err := c.GreetEveryone(context.Background()) 
	if err != nil {
		log.Fatalf("error while calling GreetEveryone RPC: %v", err)
		return
	}

	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest {
			Greeting: &greetpb.Greeting {
				FirstName : "Avery",
				LastName : "Wong",
			},
		},
		&greetpb.GreetEveryoneRequest {
			Greeting: &greetpb.Greeting {
				FirstName : "dan",
				LastName : "Wan",
			},
		},
		&greetpb.GreetEveryoneRequest {
			Greeting: &greetpb.Greeting {
				FirstName : "max",
				LastName : "Lee",
			},
		},
		&greetpb.GreetEveryoneRequest {
			Greeting: &greetpb.Greeting {
				FirstName : "Bane",
				LastName : "Anderson",
			},
		},
	}

	waitc := make(chan struct{})
	go func() {
		//send a bunch of messages
		for _, req := range requests {
		fmt.Println("Sending req %v", req)
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
			fmt.Println("Recieved " + res.GetResult())
		}
		close(waitc)
	}()

	<-waitc

}