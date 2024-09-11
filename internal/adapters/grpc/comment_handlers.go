package grpc

import (
	"context"

	"github.com/antibomberman/mego-protos/gen/go/comment"
)

func (s *serverAPI) Find(ctx context.Context, in *comment.FindRequest) (*comment.FindResponse, error) {
	return &comment.FindResponse{}, nil
}

func (s *serverAPI) Create(ctx context.Context, in *comment.CreateRequest) (*comment.CreateResponse, error) {
	return &comment.CreateResponse{}, nil
}

func (s *serverAPI) Delete(ctx context.Context, in *comment.DeleteRequest) (*comment.DeleteResponse, error) {
	return &comment.DeleteResponse{}, nil
}

func (s *serverAPI) Count(ctx context.Context, in *comment.CountRequest) (*comment.CountResponse, error) {
	return &comment.CountResponse{}, nil
}
