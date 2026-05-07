package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/threadpulse/internal/auth/handlers"
	"github.com/threadpulse/internal/middleware"
)

func Routes(r *gin.Engine, auth *handler.AuthHandler, ThreadHandler *handler.ThreadHandler, RepliesHandler *handler.RepliesHandler) {

	authHandler := r.Group("/auth", middleware.ErrorHandler())
	{
		authHandler.POST("/register", auth.RegisterHandler)
		authHandler.POST("/login", auth.Login)
	}

	Protected := r.Group("/private", middleware.Miiddleware(), middleware.ErrorHandler())
	{
		Protected.POST("/thread", ThreadHandler.CreateThreadHandler)
		Protected.PATCH("/thread/:id", ThreadHandler.UpdateThreadHandler)
		Protected.DELETE("/thread/:id", ThreadHandler.DeleteThreadHandler)

		//replies
		Protected.POST("/thread/:id/reply", RepliesHandler.CreateRepliesHandler)
		Protected.PATCH("/thread/reply/:id", RepliesHandler.UpdateRepliesHandler)
		Protected.DELETE("/thread/reply/:id", RepliesHandler.DeleteReplyHandler)

	}
	Public := r.Group("/public", middleware.ErrorHandler())
	{
		Public.GET("/threads", ThreadHandler.GetAllThreadHandler)
		Public.GET("/thread/:id", ThreadHandler.GetThreadByIdHandler)
		Public.GET("/thread/:id/replies", RepliesHandler.GetAllRepliesHandler)
	}

}
