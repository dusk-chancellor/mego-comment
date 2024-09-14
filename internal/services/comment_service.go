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

func (s *commentService) Create(comment models.Comment) (string, error) {
	id, err := s.commentRepo.Create(comment)
	if err != nil {
		log.Printf("Element: %s | Failed to create comment: %v", element, err)
		return "", err
	}
	return id, nil
}

func (s *commentService) Delete(id string) (string, error) {
	id, err := s.commentRepo.Delete(id)
	if err != nil {
		log.Printf("Element: %s | Failed to delete comment: %v", element, err)
		return "", err
	}
	return id, nil
}

func (s *commentService) Count(postId string) (int32, error) {
	count, err := s.commentRepo.Count(postId)
	if err != nil {
		log.Printf("Element: %s | Failed to count comments: %v", element, err)
		return 0, err
	}
	return count, nil
}
