package services

import "github.com/dusk-chancellor/mego-comment/internal/models"

func (s *commentService) Find(pageSize int, pageToken string) ([]*models.Comment, string, error) {
	return nil, "", nil
}

func (s *commentService) Create(comment models.Comment) (string, error) {
	return "", nil
}

func (s *commentService) Delete(comment models.Comment) (string, error) {
	return "", nil
}

func (s *commentService) Count(postId string) (int32, error) {
	return 0, nil
}
