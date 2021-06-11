package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"gogrpc/helloworld"
)

type HelloServer struct {
	helloworld.UnimplementedHelloServiceServer
}

func (s *HelloServer) SayHello(ctx context.Context, req *helloworld.HelloReq) (*helloworld.HelloResp, error) {
	// get the client information
	md, _ := metadata.FromIncomingContext(ctx)
	fmt.Printf("[SayHello] %+v Name: %s\n", md, req.GetName())
	return &helloworld.HelloResp{Result: fmt.Sprintf("Hello, %s!", req.GetName())}, nil
}

func (s *HelloServer) SayHi(ctx context.Context, req *helloworld.HelloReq) (*helloworld.HelloResp, error) {
	// get the client information
	md, _ := metadata.FromIncomingContext(ctx)
	fmt.Printf("[SayHi] %+v Name: %s\n", md, req.GetName())

	if len(req.GetName()) >= 10 {
		return nil, status.Errorf(codes.InvalidArgument,
			"Length of `Name` cannot be more than 10 characters")
	}

	return &helloworld.HelloResp{Result: fmt.Sprintf("Hi, %s!", req.GetName())}, nil
}

type HealthServer struct {
	helloworld.UnimplementedHealthServer
}

func (h *HealthServer) Check(ctx context.Context, req *helloworld.HealthParameter) (*helloworld.HelloResp, error) {
	// get the client information
	md, _ := metadata.FromIncomingContext(ctx)
	fmt.Printf("[HealthCheck] %+v\n", md)
	return &helloworld.HelloResp{Result: "Healthy"}, nil
}

func Serve() {
	log.Println("gRPC Server is listening on 0.0.0.0:50051")
	addr := fmt.Sprintf("0.0.0.0:%d", 50051)
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Cannot listen to address %s", addr)
	}

	s := grpc.NewServer()
	helloworld.RegisterHelloServiceServer(s, &HelloServer{})
	helloworld.RegisterHealthServer(s, &HealthServer{})
	if err := s.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	Serve()
}
