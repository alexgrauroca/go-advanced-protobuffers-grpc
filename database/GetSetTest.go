package database

import (
	"context"
	"log"

	"platzi.com/go/grpc/models"
)

func (repo *PostgresRepository) SetTest(ctx context.Context, test *models.Test) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO tests (id, name) VALUES ($1, $2)", test.Id, test.Name)
	return err
}

func (repo *PostgresRepository) GetTest(ctx context.Context, id string) (*models.Test, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name FROM tests WHERE id = $1", id)

	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	var test = models.Test{}
	for rows.Next() {
		err := rows.Scan(&test.Id, &test.Name)

		if err != nil {
			return nil, err
		}

		break
	}

	return &test, nil
}
