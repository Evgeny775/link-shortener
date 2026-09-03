package user

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	Repository *UserRepository
}

func NewUserService(repository *UserRepository) *UserService {
	return &UserService{
		Repository: repository,
	}
}

func (s *UserService) Register(email string, password string, userName string) (*User, error) {
	user, err := s.Repository.FindByEmail(email)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("error checking email: %w", err)
	}

	if user != nil {
		return nil, userAlreadyExists
	}

	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	newUser := User{
		Email:    email,
		Password: string(encryptedPass),
		UserName: userName,
	}

	user, err = s.Repository.CreateUser(&newUser)

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (s *UserService) Login(email string) (*User, error) {

	user, err := s.Repository.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil

}
