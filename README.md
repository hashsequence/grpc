grpc
------------------
--overview
-grpc framework deveoped by google
-at high level it allows you to define request and response
for rpc (remote procedure calls) and handles rest for you

-built on http2/2, low latency, supports streaming, language independent, super easy to plug in authentication, load balancing, logging, and monitoring

-what is an rpc?
--an a rpc is a remote procedure call
--for example it looks like you are calling
the function directly on the server
--this is different from rest api
--corba also has his put grpc implement this very cleanly

-why protocol buffers?
--protocol buffers are language agnostics
--code can be generated for pretty much any language
--data is binary and efficiently serialized(small payloads)
--very convinient for transporting lots of data
--protocol buffers allows for easy api evolution using rules


protocol buffers & language interpolability
------------------------------------------
example protocol buffer
```proto
syntax = "proto3";

message Greeting{
  string first_name = 1;
}

message GreetRequest {
  Greeting greeting = 1;
}

message GreetResponse {
  string result = 1;
}

service GreetService {
  rpc Greet(GreetR

    equest) returns (GreetResponse) {};
}
```

-protocol buffers is used to define the :
--messages(data, request, response)
--service(service name with endpoints)

-why protocol buffers over json?

grpc uses protocol buffers for communications.
lets measure the payload size vs json:

-we save bandwith with protocol buffers
-json is cpu intensive, because json is human readable
-parsing protocol buffers (binary format) is less cpu intensive
because its closer to how a machine represents data

-grpc can be used by any language
--because code can be generated for any language, makes it
simple to create micro-services in any lang that interacts with each
other

-summary
--easy to write message definition
--boiler plate code generated from simple .proto
--.payload is binary so efficient to recieve/send
--protocol buffers defines rules to make an api evolved

-what is http/2?

--faster than http1
--http1 released in 1997,
--http1 opens a new tcp connection to server at each request,
and does not compress headers(plain text)
--it only works with request/response mechanism (no server push)

--how http2 is different
-released in 2015
-http2 supports multiplexing
--client & server can push messages in parallel over same tcp connection
--greatly reduces latency
--http2 supports server push
---meaning can push multiple streams/messages for one request from client,
saving roundtrips
--http2 supports header compression, saving network bandwith
--http/2 is binary
--http1 is text, so http2 protocol buffer is a binary protocol making great
for http2
--http/2 is (Secure) ssl is reccommended by default

4 types of api in grpc
------------------------------------------

unary - classic client server response, you send something you recieve something

server streaming - as server gets new data, server opens a stream to client, and keeps sending packets of mesages

client streaming - client is streaming data to server, and server only responds after a couple of messages from client stream

bidirectional streaming - client send a couple of messages, and server responds with a couple of messages, basically asynchronous responses




unary is what traditional api look like (http rest)

```proto
service GreetService {
  //unary
  rpc Greet(GreetRequest) returns (GreetResponse) {};

  //streaming server
  rpc GreetManyTimes(GreetManyTimesRequest) returns (stream GreetManyTimesResponse) {};

  //streaming Client
  rpc longGreet(stream LongGreetRequest) returns (LongGreetResponse) {};

  //Bi directional streaming

  rpc GreetEveryone(stream GreetEveryoneRequest) responds (stream GreetEveryoneResponse) {};

}

```
scalability in grpc
--------------------------------------------------
grpc servers are asynchronous by default
this means they do not block threads on request
therefore each  grpc server can serve millions of request in parallel

grpc clients can be synchronous*blocking) or async
the client decides which model works best for performance needs
grpc lients an perform client side load balancing

security in grpc
-----------------------------------------------------
by default grpc strongly advocates to use ssl (encyption over the wire) in your api

this means grpc has security as first class citizen

each language will provide api to load grpc with required certificates and provide encryption
capability out of the box

additioally using interceptors we can provide authentication

rest api example
---------------------------------------
post, put, get, patch, delete request

post request to api, returns json with headers and everything

GRPC vs rest
--------------
-grpc uses protocol buffers , smaller and faster
-rest uses jsons, slower and bigger
-grpc uses http/2 lower latency
-rest uses http1 higher latency
-grpc is bidirectional and async
-rest is client -> server request only
-grpc has stream support
-rest is only request and response report
-grpc is api oriented "what" no constraint
-rest is crud oriented
-grpc has code generation through protocol buffers in any language - 1st class citizen
-rest has code generation through openapi / swagger 9add on) 2nd class citizen
-grpc is rpc based - client can just call the function
-rest is http verbs based ex. POST http://www.example.com/customers
sample go code for post:
```go
func doPut(url string) {
	client := &http.Client{}
	request, err := http.NewRequest("PUT", url, strings.NewReader("<golang>really</golang>"))
	request.SetBasicAuth("admin", "admin")
	request.ContentLength = 23
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("The calculated length is:", len(string(contents)), "for the url:", url)
		fmt.Println("   ", response.StatusCode)
		hdr := response.Header
		for key, value := range hdr {
			fmt.Println("   ", key, ":", value)
		}
		fmt.Println(contents)
	}
}
```

16 Code Generation Test
------------------------------------

$protoc greet.proto  --go_out=plugin=grp:.

we get greet.pb.go

```go
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greet.proto

package greetpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("greet.proto", fileDescriptor_32c0044392f32579) }

var fileDescriptor_32c0044392f32579 = []byte{
	// 67 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2f, 0x4a, 0x4d,
	0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x8c, 0xf8, 0xb8, 0x78, 0xdc,
	0x41, 0x8c, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0x27, 0xce, 0x28, 0x76, 0xb0, 0x44, 0x41,
	0x52, 0x12, 0x1b, 0x58, 0xa1, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x26, 0x07, 0x30, 0x9e, 0x37,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
}

// UnimplementedGreetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "greet.proto",
}


```

installing protoc on ubuntu:

```bash
# Make sure you grab the latest version
curl -OL https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip
# Unzip
unzip protoc-3.5.1-linux-x86_64.zip -d protoc3
# Move protoc to /usr/local/bin/
sudo mv protoc3/bin/* /usr/local/bin/
# Move protoc3/include to /usr/local/include/
sudo mv protoc3/include/* /usr/local/include/
# Optional: change owner
sudo chown [user] /usr/local/bin/protoc
sudo chown -R [user] /usr/local/include/google
```

ch.17 boilerplate server
------------------------------------------

```go
package main

import (
	"log"
	"net"
	"fmt"
	"../greetpb"
	"google.golang.org/grpc"
)

type server struct {}

func main() {
	fmt.Println("Hello World")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})
	
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
		
}
```

ch.18 boilerplate client
-----------------------------------------

```go
package main

import (
	"fmt"
	"log"
	"google.golang.org/grpc"
	"../greetpb"
)

func main() {
	fmt.Println("Hello I'm the client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()
	
	c:= greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Created client %f", c)
}
```

ch.19 unary api
-------------------------------------------

basically the type of apis that everyone is familar with 

client will send one message and then server will respond 

in grpc unary calls are defined using protocol buffers

for each rpc call we have to define a request message and a response message

new proto file

```proto
syntax = "proto3";

package greet;
option go_package="greetpb";


message Greeting {
    string first_name = 1;
    string last_name = 2;
}

message GreetRequest {
    Greeting greeting = 1;
}

message GreetResponse {
    string result =1;
}

service GreetService{
    //unary
    rpc Greet(GreetRequest) returns (GreetResponse) {};
}


```

ch.21 unary api server implementation
------------------------------------------------


```go

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	//this is how we get information from request
	firstname := req.GetGreeting().GetFirstName()

	result := "Hello " + firstname

	res := &greetpb.GreetResponse{
		Result : result,
	}

	return res, nil
}

```

 ch.22 unary api client implementation
 -----------------------------------------

 ```go
\\in main
	c:= greetpb.NewGreetServiceClient(cc)


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

 ```
