package models

import "time"

type Comment struct {
	Id 		  int64 	`db:"id"`
	Content   string 	`db:"content"`
	PostId 	  int64 	`db:"post_id"`
	ParentId  int64 	`db:"parent_id"`
	AuthorId  int64 	`db:"author_id"`
	Author 	  Author 	`db:"author"`
	CreatedAt time.Time `db:"created_at"`
}

type Author struct {
	Id 		   int64 `db:"id"`
	FirstName  string `db:"first_name"`
	MiddleName string `db:"middle_name"`
	LastName   string `db:"last_name"`
	Email 	   string `db:"email"`
	Avatar 	   string `db:"avatar"`
}
