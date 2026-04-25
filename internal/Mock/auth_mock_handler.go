package mock

import (
	"github.com/threadpulse/models"
)

type ServiceMock struct {
	RegisterFunc func(registerInput models.Register) error
	LoginFunc    func(user models.Login) error
}

func (h *ServiceMock) Register(registerInput models.Register) error {
	return h.RegisterFunc(registerInput)
}

func (h *ServiceMock) Login(user models.Login) error {
	return h.LoginFunc(user)
}
