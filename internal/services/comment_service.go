package services

import (
	"log"

	"github.com/dusk-chancellor/mego-comment/internal/models"
	"github.com/dusk-chancellor/mego-comment/pkg/utils"
)

const element = "comment_service"

func (s *commentService) Find(pageSize int, pageToken string) ([]*models.Comment, string, error) {
	var err error
	if pageSize < 1 {
		pageSize = 10
	}
	startIndex := 0
	if pageToken != "" {
		startIndex, err = utils.DecodePageToken(pageToken)
		if err != nil {
			log.Printf("Element: %s | Failed to decode page token: %v", element, err)
			return nil, "", err
		}
	}

	comments, err := s.commentRepo.Find(startIndex, pageSize+1)
	if err != nil {
		log.Printf("Element: %s | Failed to find comments: %v", element, err)
		return nil, "", err
	}

	var nextPageToken string
	if len(comments) > pageSize {
		nextPageToken = utils.EncodePageToken(startIndex + pageSize)
		comments = comments[:pageSize]
	}

	return comments, nextPageToken, nil
}

func (s *commentService) Create(comment models.Comment) (int64, error) {
	id, err := s.commentRepo.Create(comment)
	if err != nil {
		log.Printf("Element: %s | Failed to create comment: %v", element, err)
		return 0, err
	}
	return id, nil
}

func (s *commentService) Delete(id int64) (int64, error) {
	id, err := s.commentRepo.Delete(id)
	if err != nil {
		log.Printf("Element: %s | Failed to delete comment: %v", element, err)
		return 0, err
	}
	return id, nil
}

func (s *commentService) Count(postId int64) (int32, error) {
	count, err := s.commentRepo.Count(postId)
	if err != nil {
		log.Printf("Element: %s | Failed to count comments: %v", element, err)
		return 0, err
	}
	return count, nil
}

func (s *commentService) GetById(id int64) (*models.Comment, error) {
	comment, err := s.commentRepo.GetById(id)
	if err != nil {
		log.Printf("Element: %s | Failed to get by id: %v", element, err)
		return nil, err
	}
	return comment, nil
}

func (s *commentService) Exists(id int64) (bool, error) {
	exists, err := s.commentRepo.Exists(id)
	if err != nil {
		log.Printf("Element: %s | Failed to exists: %v", element, err)
		return false, err
	}
	return exists, nil
}
