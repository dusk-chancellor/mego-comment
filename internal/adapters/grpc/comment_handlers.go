package grpc

import (
	"context"
	"log"

	"github.com/antibomberman/mego-protos/gen/go/comment"
	"github.com/dusk-chancellor/mego-comment/internal/dto"
)

const element = "comment_handlers"

func (s *serverAPI) Find(ctx context.Context, in *comment.FindRequest) (*comment.FindResponse, error) {
	pageSize := int(in.GetPageSize())
	pageToken := in.GetPageToken()

	comments, nextPageToken, err := s.service.Find(pageSize, pageToken)
	if err != nil {
		log.Printf("Element %s | Failed to find: %v", element, err)
		return nil, err
	}
	pbComments := dto.ToPbComments(comments)

	return &comment.FindResponse{
		Comments:      pbComments,
		NextPageToken: nextPageToken,
	}, nil
}

func (s *serverAPI) Create(ctx context.Context, in *comment.CreateRequest) (*comment.CreateResponse, error) {
	inComment := in.GetComment()
	modelComment := dto.ToModelComment(inComment)

	id, err := s.service.Create(modelComment)
	if err != nil {
		log.Printf("Element %s | Failed to create: %v", element, err)
		return nil, err
	}

	return &comment.CreateResponse{
		Id: id,
	}, nil
}

func (s *serverAPI) Delete(ctx context.Context, in *comment.DeleteRequest) (*comment.DeleteResponse, error) {
	id := in.GetId()

	id, err := s.service.Delete(id)
	if err != nil {
		log.Printf("Element %s | Failed to delete: %v", element, err)
		return nil, err
	}

	return &comment.DeleteResponse{
		Id: id,
	}, nil
}

func (s *serverAPI) Count(ctx context.Context, in *comment.CountRequest) (*comment.CountResponse, error) {
	postId := in.GetPostId()

	count, err := s.service.Count(postId)
	if err != nil {
		log.Printf("Element %s | Failed to count: %v", element, err)
		return nil, err
	}

	return &comment.CountResponse{
		Count: count,
	}, nil
}

func (s *serverAPI) GetById(ctx context.Context, in *comment.GetByIdRequest) (*comment.Comment, error) {
	id := in.GetId()

	comment, err := s.service.GetById(id)
	if err != nil {
		log.Printf("Element %s | Failed to get by id: %v", element, err)
		return nil, err
	}
	pbComment := dto.ToPbComment(comment)

	return pbComment, nil
}

func (s *serverAPI) Exists(ctx context.Context, in *comment.GetByIdRequest) (*comment.ExistsResponse, error) {
	id := in.GetId()

	exists, err := s.service.Exists(id)	
	if err != nil {
		log.Printf("Element %s | Failed to exists: %v", element, err)
		return nil, err
	}

	return &comment.ExistsResponse{
		Exists: exists,
	}, nil
}
