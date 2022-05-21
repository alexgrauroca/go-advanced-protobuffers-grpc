package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"platzi.com/go/grpc/database"
	"platzi.com/go/grpc/server"
	"platzi.com/go/grpc/studentpb"
)

func main() {
	list, err := net.Listen("tcp", ":5060")

	if err != nil {
		log.Fatal(err)
	}

	defer list.Close()
	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	server := server.NewStudentServer(repo)
	s := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(s, server)

	reflection.Register(s)

	if err := s.Serve(list); err != nil {
		log.Fatal(err)
	}
}
