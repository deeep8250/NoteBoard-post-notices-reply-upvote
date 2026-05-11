package repositories

import "github.com/jmoiron/sqlx"

type UpvotesRepository struct {
	db *sqlx.DB
}

func NewUpvotesRepository(Db *sqlx.DB) *UpvotesRepository {
	return &UpvotesRepository{
		db: Db,
	}
}

func (r *UpvotesRepository) CreateUpvote(postID, userID int) error {
	query := `insert into upvotes(post_id,user_id) values($1,$2)`
	_, err := r.db.Exec(query, postID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (r *UpvotesRepository) GetUpvotes(postId int) (int, error) {
	var UpvotesCount int
	query := `select count(*) from upvotes where post_id=$1`
	err := r.db.Get(&UpvotesCount, query, postId)
	if err != nil {
		return 0, err
	}
	return UpvotesCount, nil
}

func (r *UpvotesRepository) CheckUpvote(postID, userID int) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM upvotes WHERE post_id=$1 AND user_id=$2`
	err := r.db.Get(&count, query, postID, userID)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
