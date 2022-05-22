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
