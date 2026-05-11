package repositories

import (
	"log"
)

type UpvoteJob struct {
	postID int
	userID int
}

type UpvoteWorker struct {
	channel chan UpvoteJob
	repo    *UpvotesRepository
}

func NewUpvoteWorker(Repo *UpvotesRepository) *UpvoteWorker {
	return &UpvoteWorker{
		channel: make(chan UpvoteJob, 100),
		repo:    Repo,
	}
}

func (w *UpvoteWorker) Start() {
	go func() {
		for job := range w.channel {
			err := w.repo.CreateUpvote(job.postID, job.userID)
			if err != nil {
				log.Println("upvote worker error:", err.Error())
			}
		}
	}()
}

func (w *UpvoteWorker) Submit(PostID, UserID int) {
	w.channel <- UpvoteJob{postID: PostID, userID: UserID}
}
