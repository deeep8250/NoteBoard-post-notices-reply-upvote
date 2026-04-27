package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	handler "github.com/threadpulse/internal/auth/handlers"
	"github.com/threadpulse/internal/middleware"
)

func Routes(r *gin.Engine, auth *handler.AuthHandler) {

	authHandler := r.Group("/auth")
	{
		authHandler.POST("/register", auth.RegisterHandler)
		authHandler.POST("/login", auth.Login)
	}

	Protected := r.Group("/protected")
	Protected.Use(middleware.Miiddleware())
	Protected.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "middleware and jwt worked",
		})

	})
}
