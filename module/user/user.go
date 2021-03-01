package user

import (
	"strconv"
	"udemy-fb-comment/module/comment"
	"udemy-fb-comment/module/post"
)

type IUser interface {
	GetID() string
	GetName() string
	AddCommentToPost(p post.IPost, c comment.IComment)
	ReplyToComment(p post.IPost, parentID string, c comment.IComment)
	EditComment(p post.IPost, parentID string, commentID string, desc string)
	DeleteComment(p post.IPost, parentID string, commentID string)
}

type User struct {
	ID   string
	Name string
}

var idIncrementor int

func GetInstance(name string) IUser {
	idIncrementor += 1

	var u User
	u.ID = "u_id_" + strconv.Itoa(idIncrementor)
	u.Name = name
	return &u
}

func (u *User) GetID() string {
	return u.ID
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) AddCommentToPost(p post.IPost, c comment.IComment) {
	p.AddComment(c)
}

func (u *User) ReplyToComment(
	p post.IPost, parentID string, c comment.IComment,
) {
	p.AddNestedComment(parentID, c)
}

func (u *User) EditComment(
	p post.IPost, parentID string, commentID string, desc string,
) {
	p.EditComment(commentID, desc)
}

func (u *User) DeleteComment(
	p post.IPost, parentID string, commentID string,
) {
	p.DeleteComment(commentID)
}
