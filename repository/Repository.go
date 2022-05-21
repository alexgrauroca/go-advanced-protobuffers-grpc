package repository

import (
	"context"

	"platzi.com/go/grpc/models"
)

type Repository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func SetStudent(ctx context.Context, student *models.Student) error {
	return implementation.SetStudent(ctx, student)
}

func GetStudent(ctx context.Context, id string) (*models.Student, error) {
	return implementation.GetStudent(ctx, id)
}
