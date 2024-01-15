package main

import (
	"fmt"
	"log"
	"net"
	"todos/internal/config"
	db "todos/internal/models"

	grpcServer "todos/internal/grpc"
	"todos/internal/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


func main() {
	// read env
	env := config.ReadEnv()

	// db
	db.ConnectDatabase(env)

	if (env.RUN_AUTO_MIGRATION) {
		db.RunAutoMigartion()
	}

	// create tcp server
	address := fmt.Sprintf(":%d", env.PORT)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	// register grpc server
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterTodoServiceServer(s, &grpcServer.Server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}