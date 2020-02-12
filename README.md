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
---------------------------------------s
