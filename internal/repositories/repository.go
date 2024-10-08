package repositories

import (
	"github.com/dusk-chancellor/mego-comment/internal/models"
	"github.com/jmoiron/sqlx"
)

type CommentRepository interface {
	Find(startIndex, pageSize int) ([]*models.Comment, error)
	Create(comment models.Comment) (int64, error)
	Delete(inID int64) (int64, error)
	Count(postId int64) (int32, error)
	GetById(id int64) (*models.Comment, error)
	Exists(id int64) (bool, error)
}

type commentRepository struct {
	db *sqlx.DB
}

func NewCommentRepository(db *sqlx.DB) CommentRepository {
	return &commentRepository{
		db: db,
	}
}