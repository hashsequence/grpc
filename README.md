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

simple to create micro-services in any lang that interacts with each other

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

grpc clients an perform client side load balancing

security in grpc
-----------------------------------------------------
by default grpc strongly advocates to use ssl (encyption over the wire) in your api

this means grpc has security as first class citizen

each language will provide api to load grpc with required certificates and provide encryption capability out of the box

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

17 boilerplate server
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

18 boilerplate client
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

19 unary api
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

21 unary api server implementation
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

 22 unary api client implementation
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

Sum Unary RPC  Api Implemetation
----------------------------------------

```bash
├── calculator
│   ├── calculator_client
│   │   └── calculator_client.go
│   ├── calculatorpb
│   │   ├── calculator.pb.go
│   │   ├── calculator.proto
│   │   └── generate.sh
│   └── calculator_server
│       └── calculator_server.go

```
calculator.proto
```proto
syntax = "proto3";

package calculator;
option go_package="calculatorpb";


message Sum {
    int32 x = 1;
    int32 y = 2;
}

message SumRequest {
    Sum sum = 1;
}

message SumResponse {
    int32 result = 1;
}

service CalculatorService{
    //unary
    rpc Sum(SumRequest) returns (SumResponse) {};
}

```

generate.sh
```bash
#!/bin/bash

protoc --go_out=plugins=grpc:. *.proto

```

run generate.sh to get

calculator.pb.go
```go
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator.proto

package calculatorpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Sum struct {
	X                    int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sum) Reset()         { *m = Sum{} }
func (m *Sum) String() string { return proto.CompactTextString(m) }
func (*Sum) ProtoMessage()    {}
func (*Sum) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{0}
}

func (m *Sum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sum.Unmarshal(m, b)
}
func (m *Sum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sum.Marshal(b, m, deterministic)
}
func (m *Sum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sum.Merge(m, src)
}
func (m *Sum) XXX_Size() int {
	return xxx_messageInfo_Sum.Size(m)
}
func (m *Sum) XXX_DiscardUnknown() {
	xxx_messageInfo_Sum.DiscardUnknown(m)
}

var xxx_messageInfo_Sum proto.InternalMessageInfo

func (m *Sum) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Sum) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

type SumRequest struct {
	Sum                  *Sum     `protobuf:"bytes,1,opt,name=sum,proto3" json:"sum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{1}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetSum() *Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

type SumResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{2}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Sum)(nil), "calculator.Sum")
	proto.RegisterType((*SumRequest)(nil), "calculator.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "calculator.SumResponse")
}

func init() { proto.RegisterFile("calculator.proto", fileDescriptor_c686ea360062a8cf) }

var fileDescriptor_c686ea360062a8cf = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x4e, 0xcc, 0x49,
	0x2e, 0xcd, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0x29, 0x72, 0x31, 0x07, 0x97, 0xe6, 0x0a, 0xf1, 0x70, 0x31, 0x56, 0x48, 0x30, 0x2a, 0x30,
	0x6a, 0xb0, 0x06, 0x31, 0x56, 0x80, 0x78, 0x95, 0x12, 0x4c, 0x10, 0x5e, 0xa5, 0x92, 0x3e, 0x17,
	0x57, 0x70, 0x69, 0x6e, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x22, 0x17, 0x73, 0x71,
	0x69, 0x2e, 0x58, 0x2d, 0xb7, 0x11, 0xbf, 0x1e, 0x92, 0xe1, 0x20, 0x45, 0x20, 0x39, 0x25, 0x55,
	0x2e, 0x6e, 0xb0, 0x86, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x31, 0x2e, 0xb6, 0xa2, 0xd4,
	0xe2, 0xd2, 0x9c, 0x12, 0xa8, 0x05, 0x50, 0x9e, 0x91, 0x2f, 0x97, 0xa0, 0x33, 0x5c, 0x77, 0x70,
	0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x05, 0xc4, 0x3d, 0x62, 0xe8, 0x06, 0x43, 0x6c, 0x97,
	0x12, 0xc7, 0x10, 0x87, 0x58, 0xa2, 0xc4, 0xe0, 0xc4, 0x17, 0xc5, 0x83, 0x90, 0x2b, 0x48, 0x4a,
	0x62, 0x03, 0x7b, 0xd6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xe9, 0xf6, 0x95, 0xb9, 0x00, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	//unary
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	//unary
	Sum(context.Context, *SumRequest) (*SumResponse, error)
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Sum(ctx context.Context, req *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculator.proto",
}

```

calculator_server.go
```go
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
```

calculator_client.go
```go
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

```

# Grpc Server Streaming

## What is a Server Streaming API?

* Server Streaming RPC API are a NEW kind API enabled thanks to HTTP/2 

* The client will send one message to the server and will recieve many responses from the server, possibly an infinite number

* Streaming Server are well suited for
	* when the server needs to send a lot of data (big data)

	* When the server needs to "push" data to the client without the client requesting for more (live feed, chat, etc)

* in gRPC Server Streaming Calls are defined using the keyword "stream"

* As for each RPC call we have to define a "Request" message and a "Response" message 

example:

```proto
message GreetManyTimesRequest {
	Greeting greeting = 1;
}

message GreetManyTimesResponse {
	string result = 1;

}

service GreetService {
	//Streaming Server 
	rpc GreetManyTimes(GreetManyTimesRequest) returns (stream GreetManyTimesResponse) {};
}
```


## What GreetManyTimes API Definition:

* It will take ONE GreetManyTimesRequest that contains a 
greeting 

* It will return MANY GreetManyTimesResponse that contains a result string


adding this in greet.proto


```proto


message GreetManytimesResponse {
	string result = 1;
}

service GreetService {
	//unary
    rpc Greet(GreetRequest) returns (GreetResponse) {};

	//server streaming
	rpc GreeManyTimes(GreetManyTimeRequest) returns(stream GreetManyTimesResponse) {};
}


```

now run 

```sh
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.

```

and if you look in the greet.pb.go we have :

```go
type GreetServiceClient interface {
	// unary
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
	//server streaming
	GreeTManyTimes(ctx context.Context, in *GreetManyTimesRequest, opts ...grpc.CallOption) (GreetService_GreetManyTimesClient, error)
}
```

now we must implement GreetManyTimes



## Server Streaming API Server Implementation

server.go
```go
func (*server) GreetManyTimes(req *greetpb.GreetManytimesRequest, stream greetpb.GreetService_GreetManyTimesServer ) error {
	fmt.Printf("GreetManyTimes has been invoked")
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManytimesResponse {
			Result: result,
		}
	
	stream.Send(res)
	time.Sleep(1000 * time.Millisecond)
	}
	return nil
}
```

## Client Streaming API Client Implementation

in greet_client.go:

```go
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
```

now run on terminal to run client:

```sh
$ go run greet/greet_client/greet_client.go
Hello I'm the client
Starting server streaming rpc
2020/08/18 00:14:32 Response from GreetManyTimes: Hello Avery number 0
2020/08/18 00:14:33 Response from GreetManyTimes: Hello Avery number 1
2020/08/18 00:14:34 Response from GreetManyTimes: Hello Avery number 2
2020/08/18 00:14:35 Response from GreetManyTimes: Hello Avery number 3
2020/08/18 00:14:36 Response from GreetManyTimes: Hello Avery number 4
2020/08/18 00:14:37 Response from GreetManyTimes: Hello Avery number 5
2020/08/18 00:14:38 Response from GreetManyTimes: Hello Avery number 6
^Csignal: interrupt

```

# gRPC client streaming

## What is a client streaming api?

* the client will send many message, and the server
will send back one response

* streaming client are well suited for 
	* when the client needs to send a lot of data
	* when the server processing is expensive and should happen as the client sends data


* In gRPC Client Streaming Calls are defined using keyword "stream"

Lets implement a LongGreet API

* It will take many longGreetRequest

* It will return one LongGreetResponse that contains 
a result string

lets modify the greet.proto

```proto

message LongGreetRequest {
    Greeting greeting = 1;
}

message LongGreetResponse {
    string result = 1;
}


service GreetService {
	// unary
    rpc Greet(GreetRequest) returns (GreetResponse) {};

	//server streaming
	rpc GreetManyTimes(GreetManytimesRequest) returns(stream GreetManytimesResponse) {};

    // Client Streaming
    rpc LongGreet(stream LongGreetRequest) returns (LongGreetResponse) {}
}

```

now the server code:

```go
func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Printf("LongGreet has been invoked\n")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("Sending %v\n", result)
			return stream.SendAndClose(&greetpb.LongGreetResponse {
				Result : result,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "! "
	}
}
```

and lastly the client:

```go
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
```