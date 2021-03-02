package main

import (
	"fmt"
	"udemy-fb-comment/module/comment"
	"udemy-fb-comment/module/post"
	"udemy-fb-comment/module/user"
)

func main() {
	fmt.Println("fb-comment")

	// setup --------------------------------
	user1 := user.GetInstance("User 1")

	c1 := comment.GetInstance()
	c1.SetDescription("C 1")
	c2 := comment.GetInstance()
	c2.SetDescription("C 2")
	c11 := comment.GetInstance()
	c11.SetDescription("C 11")
	c21 := comment.GetInstance()
	c21.SetDescription("C 21")

	p1 := post.GetInstance()

	// add comments & replies --------------------------------
	user1.AddCommentToPost(p1, c1)
	user1.AddCommentToPost(p1, c2)
	user1.ReplyToComment(p1, c1.GetID(), c11)
	user1.ReplyToComment(p1, c2.GetID(), c21)

	user1.PrintPostComments(p1)

	// edit reply --------------------------------
	user1.EditReply(p1, c1.GetID(), c11.GetID(), "C 11 edited")

	user1.PrintPostComments(p1)

	// add more replies --------------------------------
	c12 := comment.GetInstance()
	c12.SetDescription("C 12")
	c13 := comment.GetInstance()
	c13.SetDescription("C 13")
	user1.ReplyToComment(p1, c1.GetID(), c12)
	user1.ReplyToComment(p1, c1.GetID(), c13)

	user1.PrintPostComments(p1)

	// delete reply --------------------------------
	user1.DeleteReply(p1, c1.GetID(), c12.GetID())

	user1.PrintPostComments(p1)

	// delete comment --------------------------------
	user1.DeleteComment(p1, c1.GetID())

	user1.PrintPostComments(p1)
}
