package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserService interface {
	Register(user *model.User) (model.User, error)
	Login(user *model.User) (token *string, err error)
	GetUserTaskCategory() ([]model.UserTaskCategory, error)
}

type userService struct {
	userRepo repo.UserRepository
}

func NewUserService(userRepository repo.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) Register(user *model.User) (model.User, error) {
	dbUser, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return *user, err
	}

	if dbUser.Email != "" || dbUser.ID != 0 {
		return *user, errors.New("email already exists")
	}

	user.CreatedAt = time.Now()

	newUser, err := s.userRepo.CreateUser(*user)
	if err != nil {
		return *user, err
	}

	return newUser, nil
}

func (s *userService) Login(user *model.User) (token *string, err error) {
	dbUser, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if user.Password != dbUser.Password || user.Email != dbUser.Email {
		return nil, errors.New("wrong email or password")
	}

	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &model.Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := tokens.SignedString(model.JwtKey)
	if err != nil {
		return nil, err
	}

	return &ss, nil
}

func (s *userService) GetUserTaskCategory() ([]model.UserTaskCategory, error) {
	userCategory, err := s.userRepo.GetUserTaskCategory()
	if err != nil {
		return nil, err
	}

	return userCategory, nil
}
