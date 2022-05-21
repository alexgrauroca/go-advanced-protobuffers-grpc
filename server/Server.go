package server

import (
	"platzi.com/go/grpc/repository"
	"platzi.com/go/grpc/studentpb"
)

type Server struct {
	repo repository.Repository
	studentpb.UnimplementedStudentServiceServer
}
