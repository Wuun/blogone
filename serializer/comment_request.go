package serializer

type CommentRequest struct {
	Nickname string `json:"nickname"`
	Comment  string `json:"comment"`
}
