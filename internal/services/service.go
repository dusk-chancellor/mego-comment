package services

import (
	"github.com/dusk-chancellor/mego-comment/internal/models"
	"github.com/dusk-chancellor/mego-comment/internal/repositories"
)

type CommentService interface {
	Find(pageSize int, pageToken string) ([]*models.Comment, string, error)
	Create(comment models.Comment) (int64, error)
	Delete(id int64) (int64, error)
	Count(postId int64) (int32, error)
	GetById(id int64) (*models.Comment, error)
	Exists(id int64) (bool, error)
}

type commentService struct {
	commentRepo repositories.CommentRepository
}

func NewCommentService(commentRepo repositories.CommentRepository) CommentService {
	return &commentService{
		commentRepo: commentRepo,
	}
}
