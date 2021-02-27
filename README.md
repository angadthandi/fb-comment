# fb-comment (OOD)

## Entities:

- Post
- Comment
- User

## Classes:

### Post
- PostID INT
- Comments Comment[]

### Comment
- CommentID INT
- ParentID INT
- PostID INT
- UserID INT
- Desc STRING
- Comments Comment[]

### User
- UserID INT
- Name STRING

--------------------------------------

### Test FbComment Flow: