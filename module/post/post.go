package post

import (
	"strconv"
	"udemy-fb-comment/module/comment"
)

type IPost interface {
	GetComments() []comment.IComment
	AddComment(c comment.IComment)
	AddNestedComment(parentID string, c comment.IComment)
	EditComment(commentID string, desc string)
	EditNestedComment(parentID string, commentID string, desc string)
	DeleteComment(commentID string)
	DeleteNestedComment(parentID string, commentID string)
}

// var PostInstanceVar IPost

type Post struct {
	ID       string
	Comments []comment.IComment
}

var idIncrementor int

func GetInstance() IPost {
	// if PostInstanceVar == nil {
	// 	var p Post
	// 	PostInstanceVar = &p
	// }

	// return &PostInstanceVar

	idIncrementor += 1

	var p Post
	p.ID = "p_id_" + strconv.Itoa(idIncrementor)
	return &p
}

func (p *Post) GetComments() []comment.IComment {
	return p.Comments
}

func (p *Post) AddComment(c comment.IComment) {
	p.Comments = append(p.Comments, c)
}

func (p *Post) AddNestedComment(parentID string, c comment.IComment) {
	for i := 0; i < len(p.Comments); i++ {
		if p.Comments[i].GetID() == parentID {
			p.Comments[i].AddNestedComment(c)
			break
		}
	}
}

func (p *Post) EditComment(commentID string, d string) {
	for i := 0; i < len(p.Comments); i++ {
		if p.Comments[i].GetID() == commentID {
			p.Comments[i].SetDescription(d)
			break
		}
	}
}

func (p *Post) EditNestedComment(
	parentID string, commentID string, desc string,
) {
	for i := 0; i < len(p.Comments); i++ {
		if p.Comments[i].GetID() == parentID {
			p.Comments[i].EditNestedComment(commentID, desc)
			break
		}
	}
}

func (p *Post) DeleteComment(commentID string) {
	delIdx := -1
	for i := 0; i < len(p.Comments); i++ {
		if p.Comments[i].GetID() == commentID {
			delIdx = i
			break
		}
	}

	if delIdx >= 0 {
		p.Comments = append(
			p.Comments[:delIdx], p.Comments[delIdx+1:]...,
		)
	}
}

func (p *Post) DeleteNestedComment(
	parentID string, commentID string,
) {
	for i := 0; i < len(p.Comments); i++ {
		if p.Comments[i].GetID() == parentID {
			p.Comments[i].DeleteNestedComment(commentID)
			break
		}
	}
}
