package repository

import (
	"context"

	"platzi.com/go/grpc/models"
)

func SetQuestion(ctx context.Context, question *models.Question) error {
	return implementation.SetQuestion(ctx, question)
}
