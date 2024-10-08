package dto

import (
	"github.com/antibomberman/mego-protos/gen/go/comment"
	"github.com/dusk-chancellor/mego-comment/internal/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToPbComments(comments []*models.Comment) (pbComments []*comment.Comment) {
	for _, modelComment := range comments {
		pbComment := comment.Comment{
			Id:        modelComment.Id,
			Content:   modelComment.Content,
			PostId:    modelComment.PostId,
			ParentId:  modelComment.ParentId,
			Author: &comment.Author{
				Id:         modelComment.Author.Id,
				FirstName:  modelComment.Author.FirstName,
				MiddleName: modelComment.Author.MiddleName,
				LastName:   modelComment.Author.LastName,
				Email:      modelComment.Author.Email,
				Avatar:     modelComment.Author.Avatar,
			},
			CreatedAt: timestamppb.New(modelComment.CreatedAt),
		}

		pbComments = append(pbComments, &pbComment)
	}
	return
}

func ToModelComment(pbComment *comment.Comment) models.Comment {
	return models.Comment{
		Id:        pbComment.Id,
		Content:   pbComment.Content,
		PostId:    pbComment.PostId,
		ParentId:  pbComment.ParentId,
		Author: models.Author{
			Id:         pbComment.Author.Id,
			FirstName:  pbComment.Author.FirstName,
			MiddleName: pbComment.Author.MiddleName,
			LastName:   pbComment.Author.LastName,
			Email:      pbComment.Author.Email,
			Avatar:     pbComment.Author.Avatar,
		},
	}
}

func ToPbComment(modelComment *models.Comment) *comment.Comment {
	return &comment.Comment{
		Id:        modelComment.Id,
		Content:   modelComment.Content,
		PostId:    modelComment.PostId,
		ParentId:  modelComment.ParentId,
		Author: &comment.Author{
			Id:         modelComment.Author.Id,
			FirstName:  modelComment.Author.FirstName,
			MiddleName: modelComment.Author.MiddleName,
			LastName:   modelComment.Author.LastName,
			Email:      modelComment.Author.Email,
			Avatar:     modelComment.Author.Avatar,
		},
	}
}
