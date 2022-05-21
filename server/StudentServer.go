package server

import "platzi.com/go/grpc/repository"

func NewStudentServer(repo repository.Repository) *Server {
	return &Server{repo: repo}
}
