package repository

import (
	"ZebraX/apps/base/model"
	"context"
	"database/sql"
)

//StudentsInterface is interface for Students request
type StudentsInterface interface {
	InsertStudent(*model.StudentModel) (sql.Result, error)
	UpdateStudent(*model.StudentModel) (sql.Result, error)
	DeleteStudent(*model.StudentModel) (sql.Result, error)
	SelectStudent(*model.StudentModel) (*model.StudentModel, error)
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
	rsl, err := s.db.PrepareContext(s.ctx, "insert into student (ID, Name, Age) values(?, ?,?);")

	if err != nil {
		return nil, err
	}

	result, errs := rsl.ExecContext(s.ctx, d.ID, d.Name, d.Age)

	if errs != nil {
		return nil, errs
	}

	return result, nil
}

func (s studentRepository) UpdateStudent(d *model.StudentModel) (sql.Result, error) {
	rsl, err := s.db.PrepareContext(s.ctx, "update student set name=?, age=? where id=?;")

	if err != nil {
		return nil, err
	}

	result, errs := rsl.ExecContext(s.ctx, d.Name, d.Age, d.ID)

	if errs != nil {
		return nil, errs
	}

	return result, nil
}

func (s studentRepository) SelectStudent(d *model.StudentModel) (*model.StudentModel, error) {
	row := s.db.QueryRowContext(s.ctx, "select ID, Name, Age from student where id = ?;", d.ID)

	if err := row.Scan(&d.ID, &d.Name, &d.Age); err != nil {
		return nil, err
	}

	return d, nil
}

func (s studentRepository) DeleteStudent(d *model.StudentModel) (sql.Result, error) {
	rsl, err := s.db.PrepareContext(s.ctx, "delete from student where id= ?;")

	if err != nil {
		return nil, err
	}

	result, errs := rsl.ExecContext(s.ctx, d.ID)

	if errs != nil {
		return nil, errs
	}

	return result, nil
}
