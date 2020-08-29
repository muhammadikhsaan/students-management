package repository

import (
	"ZebraX/apps/base/model"
	"context"
	"database/sql"
)

type StudentsInterface interface {
	InsertStudent(*model.StudentModel) (sql.Result, error)
}

type studentRepository struct {
	ctx context.Context
	db  *sql.DB
}

//NewStudentRepository create student repository
func NewStudentRepository(ctx context.Context, db *sql.DB) *studentRepository {
	return &studentRepository{
		ctx: ctx,
		db:  db,
	}
}

func (s studentRepository) InsertStudent(d *model.StudentModel) (sql.Result, error) {
	rsl, err := s.db.PrepareContext(s.ctx, "insert into student (Name, Age) values(?,?);")

	if err != nil {
		return nil, err
	}

	result, errs := rsl.ExecContext(s.ctx, d.Name, d.Age)

	if errs != nil {
		return nil, err
	}

	return result, nil
}
