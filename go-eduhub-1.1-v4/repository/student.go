package repository

import (
	"a21hc3NpZ25tZW50/model"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	Store(student *model.Student) error
	ResetStudentRepo()
}

type studentRepository struct {
	students []model.Student
}

func NewStudentRepo() *studentRepository {
	return &studentRepository{}
}

func (s *studentRepository) FetchAll() ([]model.Student, error) {
	return s.students, nil
}

func (s *studentRepository) Store(student *model.Student) error {
	students := model.Student{
		ID:       student.ID,
		Name:     student.Name,
		Email:    student.Email,
		Phone:    student.Phone,
		CourseID: student.CourseID,
	}
	s.students = append(s.students, students)
	return nil
}

func (s *studentRepository) ResetStudentRepo() {
	s.students = []model.Student{}
}
