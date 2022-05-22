package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"platzi.com/go/grpc/database"
	"platzi.com/go/grpc/server"
	"platzi.com/go/grpc/testpb"
)

func main() {
	list, err := net.Listen("tcp", ":5070")

	if err != nil {
		log.Fatal(err)
	}

	defer list.Close()
	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	server := server.NewTestServer(repo)
	s := grpc.NewServer()

	testpb.RegisterTestServiceServer(s, server)
	reflection.Register(s)

	if err := s.Serve(list); err != nil {
		log.Fatal(err)
	}
}
