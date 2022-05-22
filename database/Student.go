package database

import (
	"context"
	"log"

	"platzi.com/go/grpc/models"
)

func (repo *PostgresRepository) SetStudent(ctx context.Context, student *models.Student) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO students (id, name, age) VALUES ($1, $2, $3)", student.Id, student.Name, student.Age)
	return err
}

func (repo *PostgresRepository) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, age FROM students WHERE id = $1", id)

	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	var student = models.Student{}
	for rows.Next() {
		err := rows.Scan(&student.Id, &student.Name, &student.Age)

		if err != nil {
			return nil, err
		}

		break
	}

	return &student, nil
}
