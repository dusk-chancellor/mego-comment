package repositories

import "github.com/dusk-chancellor/mego-comment/internal/models"

func (r *commentRepository)	Find(startIndex, pageSize int) ([]*models.Comment, error) {
	return nil, nil
}
	
	
func (r *commentRepository)	Create(comment models.Comment) (string, error) {
	return "", nil
}
	
	
func (r *commentRepository)	Delete(comment models.Comment) (string, error) {
	return "", nil
}
	
	
func (r *commentRepository)	Count(postId string) (int32, error) {
	return 0, nil
}
