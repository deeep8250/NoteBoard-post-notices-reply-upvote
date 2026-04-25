package service

import (
	"github.com/threadpulse/models"
)

type ServiceStructInterFace interface {
	Register(registerInput models.Register) error
	Login(user models.Login) (string, error)
}
