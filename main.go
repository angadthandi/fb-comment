package main

import (
	"fmt"
	"udemy-fb-comment/module/comment"
	"udemy-fb-comment/module/post"
	"udemy-fb-comment/module/user"
)

func main() {
	fmt.Println("fb-comment")

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

	user1.AddCommentToPost(p1, c1)
	user1.AddCommentToPost(p1, c2)
	user1.ReplyToComment(p1, c1.GetID(), c11)
	user1.ReplyToComment(p1, c2.GetID(), c21)

	p1Comments := p1.GetComments()
	for i := 0; i < len(p1Comments); i++ {
		fmt.Printf(
			"Comment: %s\n",
			p1Comments[i].GetDescription(),
		)

		comments := p1Comments[i].GetComments()
		for j := 0; j < len(comments); j++ {
			fmt.Printf(
				"\tReply: %s\n",
				comments[j].GetDescription(),
			)
		}
	}
}
