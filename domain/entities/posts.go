package entities

import (
	"time"
)

type PostDataModel struct {
	PostID  string          `json:"post_id" bson:"post_id,omitempty"`
	Content string          `json:"content" bson:"content,omitempty"`
	User    UserDetailModel `json:"user" bson:"user,omitempty"`
	Likes   int             `json:"likes" bson:"likes,omitempty"`
	Comment []Comment       `json:"comment" bson:"comment,omitempty"`
	PostdAt time.Time       `json:"created_at" bson:"created_at,omitempty"`
}

type UserPostModel struct {
	Content string `json:"content" bson:"content,omitempty"`
}

type Comment struct {
	CommentID string          `json:"comment_id" bson:"comment_id,omitempty"`
	Content   string          `json:"content" bson:"content,omitempty"`
	User      UserDetailModel `json:"user" bson:"user,omitempty"`
	CommentAt time.Time       `json:"created_at" bson:"created_at,omitempty"`
	Edited    bool            `json:"edited" bson:"edited,omitempty"`
}
