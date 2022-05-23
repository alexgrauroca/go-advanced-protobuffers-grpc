package database

import (
	"context"

	"platzi.com/go/grpc/models"
)

func (repo *PostgresRepository) SetQuestion(ctx context.Context, question *models.Question) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO questions (id, test_id, question, answer) VALUES ($1, $2, $3, $4)", question.Id, question.TestId, question.Question, question.Answer)
	return err
}

func (repo *PostgresRepository) SetAnswer(ctx context.Context, question *models.Question) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE questions SET answer = $1 WHERE id = $2", question.Answer, question.Id)
	return err
}
