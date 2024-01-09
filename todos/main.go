package main

import (
	"fmt"
	"log"
	"net"
	"todos/internal/config"
	pb "todos/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedTodoServiceServer
}

func main()  {
	// env
	env := config.ReadEnv()

	// grpc server
	address := fmt.Sprintf(":%d", env.PORT)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterTodoServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}