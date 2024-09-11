package repositories

import (
	"github.com/dusk-chancellor/mego-comment/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type CommentRepository interface {
	Find(startIndex, pageSize int) ([]*models.Comment, error)
	Create(comment models.Comment) (string, error)
	Delete(comment models.Comment) (string, error)
	Count(postId string) (int32, error)
}

type commentRepository struct {
	db *sqlx.DB
	redis *redis.Client
}

func NewCommentRepository(db *sqlx.DB, redis *redis.Client) CommentRepository {
	return &commentRepository{
		db: db,
		redis: redis,
	}
}