package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
	FetchWithClass() (*[]model.StudentClass, error)
}

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	var student []model.Student

	err := s.db.Find(&student).Error

	if err != nil {
		return nil, err
	}

	return student, err
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	err := s.db.Create(&student).Error
	if err != nil {
		return err
	}

	return nil

}

func (s *studentRepoImpl) Update(id int, student *model.Student) error {
	err := s.db.Table("students").Where("id = ?", id).Updates(&model.Student{Name: student.Name, Address: student.Address, ClassId: student.ClassId}).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *studentRepoImpl) Delete(id int) error {
	var student model.Student
	err := s.db.Delete(&student, "id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	var student model.Student
	err := s.db.Find(&student, "id=?", id).Error
	if err != nil {
		return nil, err
	}
	return &student, err
}

func (s *studentRepoImpl) FetchWithClass() (*[]model.StudentClass, error) {
	var students []model.Student
	var classes []model.Class
	var studentClasses []model.StudentClass

	err := s.db.Find(&students).Error
	if err != nil {
		return nil, err
	}

	err = s.db.Find(&classes).Error
	if err != nil {
		return nil, err
	}

	if len(students) == 0 {
		return &[]model.StudentClass{}, nil
	}

	err = s.db.Table("students").
		Select("students.name, students.address, classes.name as class_name, classes.professor, classes.room_number").
		Joins("JOIN classes ON students.class_id = classes.id").Scan(&studentClasses).Error
	if err != nil {
		return nil, err
	}

	return &studentClasses, nil
}
