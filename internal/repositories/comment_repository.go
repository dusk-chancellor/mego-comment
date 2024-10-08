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
	
	
func (r *commentRepository)	Create(comment models.Comment) (int64, error) {
	q := `
		INSERT INTO comments (id, content, post_id, parent_id, author_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id;
	`
	var id int64
	if err := r.db.QueryRow(
		q, comment.Id, comment.Content, comment.PostId, comment.ParentId, comment.AuthorId,
		).Scan(&id); err != nil {
		log.Printf("Element: %s | Failed to create comment: %v", element, err)
		return 0, err
	}
	return id, nil
}
	
	
func (r *commentRepository)	Delete(inID int64) (int64, error) {
	q := `
		DELETE FROM comments
		WHERE id = $1
		RETURNING id;
	`
	var outID int64
	if err := r.db.QueryRow(q, inID).Scan(&outID); err != nil {
		log.Printf("Element: %s | Failed to delete comment: %v", element, err)
		return 0, err
	}
	return outID, nil
}
	
	
func (r *commentRepository)	Count(postId int64) (int32, error) {
	q := `SELECT COUNT(*) FROM comments WHERE post_id = $1;`

	var count int32
	if err := r.db.Get(&count, q, postId); err != nil {
		log.Printf("Element: %s | Failed to count comments: %v", element, err)
		return 0, err
	}
	return count, nil
}

func (r *commentRepository)	GetById(id int64) (*models.Comment, error) {
	q := `SELECT * FROM comments WHERE id = $1;`

	var comment models.Comment
	if err := r.db.Get(&comment, q, id); err != nil {
		log.Printf("Element: %s | Failed to get by id: %v", element, err)
		return nil, err
	}
	return &comment, nil
}

func (r *commentRepository)	Exists(id int64) (bool, error) {
	q := `SELECT EXISTS(SELECT 1 FROM comments WHERE id = $1);`

	var exists bool
	if err := r.db.Get(&exists, q, id); err != nil {
		log.Printf("Element: %s | Failed to exists: %v", element, err)
		return false, err
	}
	return exists, nil
}
