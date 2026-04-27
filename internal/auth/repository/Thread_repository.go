package repository

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/threadpulse/models"
)

type ThreadsRepo struct {
	Db *sqlx.DB
}

func NewThreadRepo(db *sqlx.DB) *ThreadsRepo {
	return &ThreadsRepo{
		Db: db,
	}
}

func (r *ThreadsRepo) CreateThreads(input *models.CreateThread) error {
	query := `insert into posts(user_id,title,content) values($1,$2,$3)`
	status, err := r.Db.Exec(query, input.UserID, input.Title, input.Content)

	if err != nil {
		return err
	}

	rowsAffected, _ := status.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("Something went wrong , while creating the post")
	}

	return nil
}

func (r *ThreadsRepo) GetAllThreads(limit, offset int) (*[]models.CreateThread, int, error) {
	var ThreadsAll []models.CreateThread
	query := `select * from posts order by created_at desc limit $1 offset $2`
	err := r.Db.Select(&ThreadsAll, query, limit, offset)
	if err != nil {
		return nil, 0, err

	}
	var count int
	query = `select count(*) from posts`
	err = r.Db.Get(&count, query)
	if err != nil {
		return nil, 0, err
	}

	return &ThreadsAll, count, nil
}

func (r *ThreadsRepo) GetThreadByID(threadId int) (*models.CreateThread, error) {
	var Thread models.CreateThread
	query := `select * from posts where id=$1`
	err := r.Db.Get(&Thread, query, threadId)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errors.New("Thread not found")
		} else {

			return nil, err
		}
	}
	return &Thread, nil
}

func (r *ThreadsRepo) UpdateThread(ThreadId, userID int, input models.UpdateThread) error {
	sql := `update posts set title=$1 , content=$2 where id=$3 and user_id=$4`
	status, err := r.Db.Exec(sql, input.Title, input.Content, ThreadId)
	if err != nil {
		return err
	}
	rowsAffected, _ := status.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("cant update the ROWS")
	}
	return nil
}

func (r *ThreadsRepo) DeleteThread(ThreadID, UserID int) error {
	sql := `delete from posts where id=$1 and user_id=$2`
	status, err := r.Db.Exec(sql, ThreadID, UserID)
	if err != nil {
		return err
	}
	RowsAffected, _ := status.RowsAffected()
	if RowsAffected == 0 {
		return errors.New("failed to delete the thread")
	}
	return nil
}
