package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/threadpulse/internal/auth/handlers"
)

func Routes(r *gin.Engine, auth *handler.AuthService) {

	authHandler := r.Group("/auth")
	{
		authHandler.POST("/register", auth.RegisterHandler)
		authHandler.POST("/login", auth.Login)
	}
}
