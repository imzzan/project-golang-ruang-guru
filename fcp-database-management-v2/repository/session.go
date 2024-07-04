package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type SessionsRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailName(name string) error
	SessionAvailToken(token string) (model.Session, error)
}

type sessionsRepoImpl struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) *sessionsRepoImpl {
	return &sessionsRepoImpl{db}
}

func (s *sessionsRepoImpl) AddSessions(session model.Session) error {
	err := s.db.Create(&session).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *sessionsRepoImpl) DeleteSession(token string) error {
	var session model.Session
	err := s.db.Delete(&session, "token=?", token).Error
	if err == nil {
		return nil
	}

	return err
}

func (s *sessionsRepoImpl) UpdateSessions(session model.Session) error {
	err := s.db.Table("sessions").Where("username = ? AND expiry = ?", session.Username, session.Expiry).Update("token", session.Token).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *sessionsRepoImpl) SessionAvailName(name string) error {
	var session model.Session
	err := s.db.First(&session).Where("username=?", name).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *sessionsRepoImpl) SessionAvailToken(token string) (model.Session, error) {
	var session model.Session
	err := s.db.First(&session).Where("token=?", token).Error
	if err != nil {
		return model.Session{}, err
	}

	return session, nil
}
