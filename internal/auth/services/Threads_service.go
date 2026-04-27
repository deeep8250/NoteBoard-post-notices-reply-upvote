package service

import (
	"github.com/threadpulse/internal/auth/repository"
	"github.com/threadpulse/models"
)

type ThreadsService struct {
	repo *repository.ThreadsRepo
}

func NewThreadsService(repo *repository.ThreadsRepo) *ThreadsService {
	return &ThreadsService{
		repo: repo,
	}

}

func (s *ThreadsService) CreateThread(userID int, input models.CreateThread) error {

	input.UserID = userID

	err := s.repo.CreateThreads(&input)
	if err != nil {
		return err
	}
	return nil
}

func (s *ThreadsService) GetAllThreads(page, limit int) ([]models.CreateThread, int, error) {
	offset := (page - 1) * limit

	threads, count, err := s.repo.GetAllThreads(limit, offset)
	if err != nil {
		return nil, 0, err
	}
	return *threads, count, nil

}

func (s *ThreadsService) GetThreadById(ThreadID int) (*models.CreateThread, error) {
	thread, err := s.repo.GetThreadByID(ThreadID)
	if err != nil {
		return nil, err
	}
	return thread, nil
}

func (s *ThreadsService) UpdateThread(threadID, userID int, input models.UpdateThread) error {
	err := s.repo.UpdateThread(threadID, userID, input)
	if err != nil {
		return err
	}
	return nil
}

func (s *ThreadsService) DeleteThread(ThreadID, UserID int) error {
	err := s.repo.DeleteThread(ThreadID, UserID)
	if err != nil {
		return err
	}
	return nil
}
