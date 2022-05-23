package repository

import (
	"context"

	"platzi.com/go/grpc/models"
)

func SetQuestion(ctx context.Context, question *models.Question) error {
	return implementation.SetQuestion(ctx, question)
}

func SetAnswer(ctx context.Context, question *models.Question) error {
	return implementation.SetAnswer(ctx, question)
}
