package service

import (
	"database/sql"
	"errors"

	"github.com/threadpulse/internal/auth/repository"
	"github.com/threadpulse/models"

	jwt "github.com/threadpulse/internal/JTW"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	AuthRepo *repository.AuthRepo
}

func NewAuthService(repo *repository.AuthRepo) *AuthService {
	return &AuthService{
		AuthRepo: repo,
	}
}

func (s *AuthService) Register(registerInput models.Register) error {

	existUser, err := s.AuthRepo.VerifyByEmail(registerInput.Email)
	if existUser != nil {
		return errors.New("user already exists")
	} else if err != sql.ErrNoRows {
		return err
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(registerInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	User := &models.User{
		Name:         registerInput.Name,
		Email:        registerInput.Email,
		HashPassword: string(hashedPass),
	}

	err = s.AuthRepo.RegisterNewUserRepo(User)
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthService) Login(user models.Login) (string, error) {
	ReturnedUser, err := s.AuthRepo.VerifyByEmail(user.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(ReturnedUser.HashPassword), []byte(user.Password))
	if err != nil {
		return "", errors.New("invalid user")
	}

	token, err := jwt.JWTinit(ReturnedUser.Id)
	if err != nil {
		return "", err
	}

	return token, nil

}
