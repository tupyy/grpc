package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/tupyy/grpc/helloworld"

	"google.golang.org/grpc"
)

type helloServer struct {
	pb.GreeterServer
}

func (s *helloServer) SayHello(ctx context.Context, person *pb.Person) *pb.PersonResponse {
	fmt.Printf("Say hello: %+v\n", person)
	return &pb.PersonResponse{"hello to you"}
}

func main() {
	lis, _ := net.Listen("tcp", "localhost:7000")
	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer)
	grpcServer.Serve(lis)
}
