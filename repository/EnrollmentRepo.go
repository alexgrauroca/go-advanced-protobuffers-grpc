package repository

import (
	"context"

	"platzi.com/go/grpc/models"
)

func SetEnrollment(ctx context.Context, enrollment *models.Enrollment) error {
	return implementation.SetEnrollment(ctx, enrollment)
}
