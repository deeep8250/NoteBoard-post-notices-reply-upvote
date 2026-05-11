package services

import (
	"github.com/threadpulse/internal/replies/repository"
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

func (s *RepliesService) CreateRepliesService(reply models.Replies) error {
	err := s.service.CreateRepliesRepo(reply)
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

func (s *RepliesService) UpdateReplyService(UpdatedReply models.Replies) error {

	err := s.service.UpdateReply(UpdatedReply)
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
