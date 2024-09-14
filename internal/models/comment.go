package models

import "time"

type Comment struct {
	Id 		  string 	`db:"id"`
	Content   string 	`db:"content"`
	PostId 	  string 	`db:"post_id"`
	ParentId  string 	`db:"parent_id"`
	CommentId string 	`db:"comment_id"`
	AuthorId  string 	`db:"author_id"`
	Author 	  Author 	`db:"author"`
	CreatedAt time.Time `db:"created_at"`
}

type Author struct {
	Id 		   string `db:"id"`
	FirstName  string `db:"first_name"`
	MiddleName string `db:"middle_name"`
	LastName   string `db:"last_name"`
	Email 	   string `db:"email"`
	Avatar 	   string `db:"avatar"`
}
