package repositories

import (
	"github.com/dusk-chancellor/mego-comment/internal/models"
	"github.com/jmoiron/sqlx"
)

type CommentRepository interface {
	Find(startIndex, pageSize int) ([]*models.Comment, error)
	Create(comment models.Comment) (string, error)
	Delete(inID string) (string, error)
	Count(postId string) (int32, error)
}

type commentRepository struct {
	db *sqlx.DB
}

func NewCommentRepository(db *sqlx.DB) CommentRepository {
	return &commentRepository{
		db: db,
	}
}