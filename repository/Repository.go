package repository

import (
	"context"

	"platzi.com/go/grpc/models"
)

type Repository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) error
	GetTest(ctx context.Context, id string) (*models.Test, error)
	SetTest(ctx context.Context, test *models.Test) error
	SetQuestion(ctx context.Context, question *models.Question) error
	SetEnrollment(ctx context.Context, enrollment *models.Enrollment) error
	GetStudentsPerTest(ctx context.Context, testId string) ([]*models.Student, error)
	GetQuestionPerTest(ctx context.Context, testId string) ([]*models.Question, error)
	SetAnswer(ctx context.Context, answer *models.Answer) error
	GetTestScore(ctx context.Context, testId, studentId string) (*models.TestScore, error)
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}
