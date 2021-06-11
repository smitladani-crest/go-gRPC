# Go-gRPC

Sample gRPC go application for AWS Application Load Balancer gRPC functionality demo
    
## Instructions

Generate protobuf files:

    $ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative hello.proto

Build server and client, and run:

    $ go build server.go
    $ go build client.go
    $ ./server &
    $ ./client

## Reference

[grpc.io](https://grpc.io/docs/languages/go/quickstart/)
[gRPC Errors](https://github.com/avinassh/grpc-errors)
[exampleloadbalancer.com](https://exampleloadbalancer.com/albgrpc_demo.html)