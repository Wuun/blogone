package serializer

//Response is the response serialization of http.Response.Body
type Response struct {
	Code    int
	Content interface{}
}

func BuildResponse(code int, content interface{}) *Response {
	return &Response{
		Code:    code,
		Content: content,
	}
}
