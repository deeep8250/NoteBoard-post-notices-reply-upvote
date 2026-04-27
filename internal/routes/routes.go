package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/threadpulse/internal/auth/handlers"
	"github.com/threadpulse/internal/middleware"
)

func Routes(r *gin.Engine, auth *handler.AuthHandler, ThreadHandler *handler.ThreadHandler) {

	authHandler := r.Group("/auth")
	{
		authHandler.POST("/register", auth.RegisterHandler)
		authHandler.POST("/login", auth.Login)
	}

	Protected := r.Group("/private", middleware.Miiddleware())
	{
		Protected.POST("/thread", ThreadHandler.CreateThreadHandler)
		Protected.PATCH("/thread/:id", ThreadHandler.UpdateThreadHandler)
		Protected.DELETE("/thread/:id", ThreadHandler.DeleteThreadHandler)

	}
	Public := r.Group("/public")
	{
		Public.GET("/threads", ThreadHandler.GetAllThreadHandler)
		Public.GET("/thread/:id", ThreadHandler.GetThreadByIdHandler)
	}

}
