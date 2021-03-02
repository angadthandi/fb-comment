package user

import (
	"fmt"
	"strconv"
	"udemy-fb-comment/module/comment"
	"udemy-fb-comment/module/post"
)

type IUser interface {
	GetID() string
	GetName() string
	AddCommentToPost(p post.IPost, c comment.IComment)
	ReplyToComment(p post.IPost, parentID string, c comment.IComment)
	EditComment(p post.IPost, commentID string, desc string)
	EditReply(p post.IPost, parentID string, commentID string, desc string)
	DeleteComment(p post.IPost, commentID string)
	DeleteReply(p post.IPost, parentID string, commentID string)
	PrintPostComments(p post.IPost)
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
	p post.IPost, commentID string, desc string,
) {
	p.EditComment(commentID, desc)
}

func (u *User) EditReply(
	p post.IPost, parentID string, commentID string, desc string,
) {
	p.EditNestedComment(parentID, commentID, desc)
}

func (u *User) DeleteComment(
	p post.IPost, commentID string,
) {
	p.DeleteComment(commentID)
}

func (u *User) DeleteReply(
	p post.IPost, parentID string, commentID string,
) {
	p.DeleteNestedComment(parentID, commentID)
}

func (u *User) PrintPostComments(
	p post.IPost,
) {
	pComments := p.GetComments()
	for i := 0; i < len(pComments); i++ {
		fmt.Printf(
			": %s. Comment: %s\n",
			pComments[i].GetID(),
			pComments[i].GetDescription(),
		)

		comments := pComments[i].GetComments()
		for j := 0; j < len(comments); j++ {
			fmt.Printf(
				"\t: %s. Reply: %s\n",
				comments[j].GetID(),
				comments[j].GetDescription(),
			)
		}
	}

	if len(pComments) > 0 {
		fmt.Println("******************************************")
	}
}
