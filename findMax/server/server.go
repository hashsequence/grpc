package main

import (
	"log"
	"net"
	"fmt"
	findMax "../pb"
	"google.golang.org/grpc"
	"io"
)

type server struct {}

func (*server) ComputeMax(stream findMax.FindMaxService_ComputeMaxServer) error {
	fmt.Printf("ComputeMax function has invoked with a streaming request\n")
	firstCall := true
	max := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v", err)
			return err
		}
		val := req.GetVal() 
		if firstCall == true {
			firstCall = false
			max = val
		} else if val > max {
			max = val
		}
		fmt.Printf("ComputeMax sending %v\n",max)
		senderErr := stream.Send(&findMax.FindMaxResponse {
			Max: max,
		})
		if senderErr != nil {
			log.Fatalf("Error while sending data to cclient: %v\n", err)
			return err
		}
	}
}


func main() {
	fmt.Println("Hello World")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	findMax.RegisterFindMaxServiceServer(s, &server{})
	
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
		
}