package repo

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TeacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) TeacherRepo {
	return TeacherRepo{db}
}

func (t TeacherRepo) Save(data model.Teacher) error {
	err := t.db.Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (t TeacherRepo) Query() ([]model.Teacher, error) {
	var teacher []model.Teacher

	err := t.db.Find(&teacher).Error
	if err != nil {
		return nil, err
	}

	return teacher, nil
}

func (t TeacherRepo) Update(id uint, name string) error {
	var taecher model.Teacher
	return t.db.Model(&taecher).Where("id = ?", id).Update("name", name).Error
}

func (t TeacherRepo) Delete(id uint) error {
	teacher := model.Teacher{}
	err := t.db.Where("id = ?", id).Delete(&teacher).Error
	if err != nil {
		return err
	}
	return nil
}
