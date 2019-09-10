package serializer

//Response is the response serialization of http.Response.Body
type Response struct {
	Status int
	Msg    string
	Data   interface{}
	Error  error
}

//BuildResponse is the fatory of response.
func BuildResponse(code int, msg string, data interface{}, err error) *Response {
	return &Response{
		Status: code,
		Msg:    msg,
		Data:   data,
		Error:  err,
	}
}
