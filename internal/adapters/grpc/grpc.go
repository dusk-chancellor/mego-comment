package grpc

import (
	"github.com/dusk-chancellor/mego-comment/internal/config"
	"github.com/dusk-chancellor/mego-comment/internal/services"

	"github.com/antibomberman/mego-protos/gen/go/comment"
	"google.golang.org/grpc"
)

type serverAPI struct {
	comment.UnimplementedCommentServiceServer
	service services.CommentService
	cfg *config.Config
}


func RegisterGRPC(grpc *grpc.Server, service services.CommentService, cfg *config.Config) {
	comment.RegisterCommentServiceServer(grpc, &serverAPI{
		service: service,
		cfg: cfg,
	})
}
