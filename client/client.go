package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"gogrpc/helloworld"
)

func main() {

	// config := &tls.Config{
	// 	InsecureSkipVerify: true,
	// }
	// conn, err := grpc.Dial("127.0.0.1:443", grpc.WithTransportCredentials(credentials.NewTLS(config)))

	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	c := helloworld.NewHelloServiceClient(conn)
	// ideally, you should handle error here too, for brevity
	// I am ignoring that
	resp, _ := c.SayHello(
		context.Background(),
		&helloworld.HelloReq{Name: "Smit"},
	)
	fmt.Println(resp.GetResult())

	resp, _ = c.SayHi(
		context.Background(),
		&helloworld.HelloReq{Name: "Ladani"},
	)
	fmt.Println(resp.GetResult())

	healthCheck := helloworld.NewHealthClient(conn)
	resp, _ = healthCheck.Check(context.Background(), new(helloworld.HealthParameter))
	fmt.Println(resp.GetResult())
}
