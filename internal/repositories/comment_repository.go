package repositories

import (
	"log"

	"github.com/dusk-chancellor/mego-comment/internal/models"
)

const element = "comment_repository"

func (r *commentRepository)	Find(startIndex, pageSize int) ([]*models.Comment, error) {
	// redis
	
	// db
	commentQuery := `SELECT * FROM comments LIMIT $1 OFFSET $2;`
	authorQuery := `SELECT * FROM authors WHERE id = $1;`

	var comments []*models.Comment
	if err := r.db.Select(&comments, commentQuery, pageSize, startIndex); err != nil {
		log.Printf("Element: %s | Failed to find comments: %v", element, err)
		return nil, err
	}
	if len(comments) == 0 {
		return []*models.Comment{}, nil
	}

	for i, comment := range comments {
		var author models.Author
		if err := r.db.Get(&author, authorQuery, comment.AuthorId); err != nil {
			log.Printf("Element: %s | Failed to find author: %v", element, err)
			return nil, err
		}
		comments[i].Author = author
	}

	return comments, nil
}
	
	
func (r *commentRepository)	Create(comment models.Comment) (string, error) {
	q := `
		INSERT INTO comments (id, content, post_id, parent_id, comment_id, author_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id;
	`
	var id string
	if err := r.db.QueryRow(
		q, comment.Id, comment.Content, comment.PostId, comment.ParentId, comment.CommentId, comment.AuthorId,
		).Scan(&id); err != nil {
		log.Printf("Element: %s | Failed to create comment: %v", element, err)
		return "", err
	}
	return id, nil
}
	
	
func (r *commentRepository)	Delete(inID string) (string, error) {
	q := `
		DELETE FROM comments
		WHERE id = $1
		RETURNING id;
	`
	var outID string
	if err := r.db.QueryRow(q, inID).Scan(&outID); err != nil {
		log.Printf("Element: %s | Failed to delete comment: %v", element, err)
		return "", err
	}
	return outID, nil
}
	
	
func (r *commentRepository)	Count(postId string) (int32, error) {
	q := `SELECT COUNT(*) FROM comments WHERE post_id = $1;`

	var count int32
	if err := r.db.Get(&count, q, postId); err != nil {
		log.Printf("Element: %s | Failed to count comments: %v", element, err)
		return 0, err
	}
	return count, nil
}
