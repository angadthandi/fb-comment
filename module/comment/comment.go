package comment

import "strconv"

type IComment interface {
	GetID() string
	GetDescription() string
	SetDescription(d string)
	GetComments() []IComment
	AddNestedComment(c IComment)
	EditNestedComment(commentID string, desc string)
	DeleteNestedComment(commentID string)
}

type Comment struct {
	ID       string
	ParentID string
	PostID   string
	UserID   string
	Desc     string
	Comments []IComment
}

var idIncrementor int

func GetInstance() IComment {
	idIncrementor += 1

	var c Comment
	c.ID = "c_id_" + strconv.Itoa(idIncrementor)
	return &c
}

func (c *Comment) GetID() string {
	return c.ID
}

func (c *Comment) GetDescription() string {
	return c.Desc
}

func (c *Comment) SetDescription(d string) {
	c.Desc = d
}

func (c *Comment) GetComments() []IComment {
	return c.Comments
}

func (c *Comment) AddNestedComment(nc IComment) {
	c.Comments = append(c.Comments, nc)
}

func (c *Comment) EditNestedComment(id string, desc string) {
	for i := 0; i < len(c.Comments); i++ {
		if c.Comments[i].GetID() == id {
			c.Comments[i].SetDescription(desc)
			break
		}
	}
}

func (c *Comment) DeleteNestedComment(id string) {
	delIdx := -1
	for i := 0; i < len(c.Comments); i++ {
		if c.Comments[i].GetID() == id {
			delIdx = i
			break
		}
	}

	if delIdx >= 0 {
		c.Comments = append(
			c.Comments[:delIdx], c.Comments[delIdx+1:]...,
		)
	}
}
