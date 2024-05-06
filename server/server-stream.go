package main

import (
	"log"
	"time"

	pb "github.com/dcrespo1/basic-go-gRPC/proto"
)

func (s *helloServer) callSayHelloServerStream(req *pb.NamesList, stream *pb.) error {
	log.Printf("got request with names: %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello" + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
