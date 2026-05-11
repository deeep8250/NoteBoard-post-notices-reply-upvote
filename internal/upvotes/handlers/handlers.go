package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/threadpulse/internal/upvotes/services"
)

type UpvoteHandler struct {
	service *services.UpvoteService
}

func NewUpvoteHandler(Serv *services.UpvoteService) *UpvoteHandler {
	return &UpvoteHandler{
		service: Serv,
	}
}

func (h *UpvoteHandler) Upvote(c *gin.Context) {
	postIDstring := c.Param("id")
	postID, err := strconv.Atoi(postIDstring)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	userID, ok := c.Get("userID")
	if !ok {
		c.Error(errors.New("unauthorized user"))
		c.Abort()
		return
	}

	h.service.SubmitUpvote(postID, userID.(int))

	c.JSON(http.StatusOK, gin.H{
		"status": "upvoted",
	})
}

func (h *UpvoteHandler) GetAllUpvotes(c *gin.Context) {
	postIDstring := c.Param("id")
	postID, err := strconv.Atoi(postIDstring)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	countVotes, err := h.service.GetUpvotes(postID)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"upvotes": countVotes,
	})
}
