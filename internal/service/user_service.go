package service

import (
	"simple-golang-social-media-app/internal/model"
	"simple-golang-social-media-app/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(email, username, password string) error
	Login(email, password string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{r}
}

func (s *userService) Register(email, username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &model.User{
		Email:    email,
		Username: username,
		Password: string(hashedPassword),
	}
	return s.repo.Create(user)
}

func (s *userService) Login(email, password string) (*model.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) FindByEmail(email string) (*model.User, error) {
	return s.repo.FindByEmail(email)
}
