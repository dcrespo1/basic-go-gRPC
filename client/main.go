package main

import (
	"log"

	pb "github.com/dcrespo1/basic-go-gRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = ":8080"

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := pb.NamesList{
		Names: []string{"poopy", "stinky", "Alice", "Bob"},
	}

	// callSayHello(client) //Unary Request Function Call
	callSayHelloServerStream(client, names)
}
