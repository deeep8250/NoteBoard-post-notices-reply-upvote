package routes

import (
	"github.com/gin-gonic/gin"
	auth "github.com/threadpulse/internal/auth/handlers"
	"github.com/threadpulse/internal/middleware"
	replies "github.com/threadpulse/internal/replies/handlers"
	thread "github.com/threadpulse/internal/threads/handlers"
	upvote "github.com/threadpulse/internal/upvotes/handlers"
)

func Routes(r *gin.Engine, auth *auth.AuthHandler, ThreadHandler *thread.ThreadHandler, RepliesHandler *replies.RepliesHandler, upvote *upvote.UpvoteHandler) {

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

		//replies
		Protected.POST("/thread/:id/reply", RepliesHandler.CreateRepliesHandler)
		Protected.PATCH("/thread/reply/:id", RepliesHandler.UpdateRepliesHandler)
		Protected.DELETE("/thread/reply/:id", RepliesHandler.DeleteReplyHandler)

		//upvotes
		Protected.POST("/thread/:id/upvote", upvote.Upvote)

	}
	Public := r.Group("/public")
	{
		Public.GET("/threads", ThreadHandler.GetAllThreadHandler)
		Public.GET("/thread/:id", ThreadHandler.GetThreadByIdHandler)
		Public.GET("/thread/:id/replies", RepliesHandler.GetAllRepliesHandler)
		Public.GET("/thread/:id/upvotes", upvote.GetAllUpvotes)
	}

}
