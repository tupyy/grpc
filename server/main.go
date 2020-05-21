package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/tupyy/grpc/helloworld/helloworld"

	"google.golang.org/grpc"
)

type helloServer struct {
	pb.GreeterServer
}

func (s *helloServer) SayHello(ctx context.Context, person *pb.Person) (*pb.PersonResponse, error) {
	fmt.Printf("Say hello: %+v\n", person)
	return &pb.PersonResponse{Message: "hello to you"}, nil
}

func main() {
	lis, _ := net.Listen("tcp", "localhost:50000")
	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &helloServer{})
	grpcServer.Serve(lis)
}
