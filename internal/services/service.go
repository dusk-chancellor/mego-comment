package services

import (
	"github.com/dusk-chancellor/mego-comment/internal/models"
	"github.com/dusk-chancellor/mego-comment/internal/repositories"
)

type CommentService interface {
	Find(pageSize int, pageToken string) ([]*models.Comment, string, error)
	Create(comment models.Comment) (string, error)
	Delete(id string) (string, error)
	Count(postId string) (int32, error)
}

type commentService struct {
	commentRepo repositories.CommentRepository
}

func NewCommentService(commentRepo repositories.CommentRepository) CommentService {
	return &commentService{
		commentRepo: commentRepo,
	}
}
