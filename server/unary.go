package main

import (
	"context"

	pb "github.com/dcrespo1/basic-go-gRPC/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello Message",
	}, nil
}
