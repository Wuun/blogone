package serializer

//CommentRequest is the comment request serializer
type CommentRequest struct {
	Nickname string `json:"nickname"`
	Comment  string `json:"comment"`
}
