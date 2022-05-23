package repository

import (
	"context"

	"platzi.com/go/grpc/models"
)

func SetTest(ctx context.Context, test *models.Test) error {
	return implementation.SetTest(ctx, test)
}

func GetTest(ctx context.Context, id string) (*models.Test, error) {
	return implementation.GetTest(ctx, id)
}

func GetStudentsPerTest(ctx context.Context, testId string) ([]*models.Student, error) {
	return implementation.GetStudentsPerTest(ctx, testId)
}

func GetQuestionPerTest(ctx context.Context, testId string) ([]*models.Question, error) {
	return implementation.GetQuestionPerTest(ctx, testId)
}
