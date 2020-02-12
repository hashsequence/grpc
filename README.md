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

'''proto
syntax = "proto3"
'''proto
