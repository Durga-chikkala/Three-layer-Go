package stores

import (
	"context"
	"threeLayer/models"
)

type store struct {
}

func New() store {
	return store{}
}

func (s store) CalculateMarks(ctx context.Context, student models.Student) (models.Student, error) {
	_ = "INSERT INTO student (id, name, marks, grade) VALUES (?, ?, ?, ?)"

	return student, nil
}
