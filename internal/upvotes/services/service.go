package services

import (
	"github.com/threadpulse/internal/upvotes/repositories"
)

type UpvoteService struct {
	repo   *repositories.UpvotesRepository
	worker *repositories.UpvoteWorker
}

func NewUpvoteService(Repo *repositories.UpvotesRepository, Worker *repositories.UpvoteWorker) *UpvoteService {
	return &UpvoteService{
		repo:   Repo,
		worker: Worker,
	}
}

func (s *UpvoteService) SubmitUpvote(postID, userID int) {
	s.worker.Submit(postID, userID)

}

func (s *UpvoteService) GetUpvotes(postID int) (int, error) {
	upvotes, err := s.repo.GetUpvotes(postID)
	if err != nil {
		return 0, err
	}
	return upvotes, nil

}
