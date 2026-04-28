package service

import (
	"errors"

	"github.com/threadpulse/internal/auth/repository"
	"github.com/threadpulse/models"
)

type RepliesService struct {
	service *repository.RepliesRepo
}

func NewRepliesService(serv *repository.RepliesRepo) *RepliesService {
	return &RepliesService{
		service: serv,
	}
}

func (s *RepliesService) CreateRepliesService(postID, userID int, reply string) error {
	err := s.service.CreateRepliesRepo(postID, userID, reply)
	if err != nil {
		return err
	}
	return nil
}

func (s *RepliesService) GetAllRepliessService(postID int, limit, page int) ([]models.Replies, int, error) {
	offset := (page - 1) * limit
	replies, count, err := s.service.GetAllReplies(postID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	return replies, count, nil
}

func (s *RepliesService) UpdateReplyService(postID, userID int, UpdatedReply string) error {
	if UpdatedReply == "" {
		return errors.New("there is nothing to update")
	}
	err := s.service.UpdateReply(postID, userID, UpdatedReply)
	if err != nil {
		return err
	}
	return nil
}

func (s *RepliesService) DeleteReplyService(replyId, userID int) error {
	err := s.service.DeleteReply(replyId, userID)
	if err != nil {
		return err
	}
	return nil
}
