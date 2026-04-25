package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/threadpulse/internal/auth/services"
	"github.com/threadpulse/models"
)

type AuthService struct {
	services service.ServiceStructInterFace
}

func NewAuthHandler(serv service.ServiceStructInterFace) *AuthService {
	return &AuthService{services: serv}
}

func (h *AuthService) RegisterHandler(c *gin.Context) {
	var UserRegisterInput models.Register
	err := c.ShouldBind(&UserRegisterInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.services.Register(UserRegisterInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"response": "user created",
	})
}
func (h *AuthService) Login(c *gin.Context) {

	var UserLoginInput models.Login

	err := c.ShouldBindJSON(&UserLoginInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := h.services.Login(UserLoginInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
