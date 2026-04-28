package repository

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/threadpulse/models"
)

type RepliesRepo struct {
	Db *sqlx.DB
}

func NewRepliesRepo(db *sqlx.DB) *RepliesRepo {
	return &RepliesRepo{
		Db: db,
	}
}

func (r *RepliesRepo) CreateRepliesRepo(postID, userID int, reply string) error {
	query := `insert into replies(post_id,replied_user_id,reply) values($1,$2,$3)`
	status, err := r.Db.Exec(query, postID, userID, reply)
	if err != nil {
		return err
	}
	rowsAffected, _ := status.RowsAffected()

	if rowsAffected == 0 {
		return errors.New("something went wrong while creating your reply")
	}
	return nil
}

func (r *RepliesRepo) GetAllReplies(postID, limit, offset int) ([]models.Replies, error) {
	var replies []models.Replies
	query := `select * from replies where post_id=$1 limit $2 offset $3`
	err := r.Db.Select(&replies, query, postID, limit, offset)
	if err != nil {
		return nil, err
	}
	return replies, nil
}

func (r *RepliesRepo) UpdateReply(replyID, userID int, UpdatedReply string) error {
	query := `update replies set reply=$1 where id=$2 and replied_user_id=$3`
	status, err := r.Db.Exec(query, UpdatedReply, replyID, userID)
	if err != nil {
		return err
	}
	satusRowsAffected, _ := status.RowsAffected()
	if satusRowsAffected == 0 {
		return errors.New("cant update the reply something went wrong")
	}
	return nil
}

func (r *RepliesRepo) DeleteReply(replyID, userID int) error {
	query := `delete from replies where id=$1 and replied_user_id=$2`
	status, err := r.Db.Exec(query, replyID, userID)
	if err != nil {
		return err
	}
	statusRowsAffected, _ := status.RowsAffected()
	if statusRowsAffected == 0 {
		return errors.New("something went wrong while deleting the reply please try again later")
	}
	return nil
}
