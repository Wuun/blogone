package serializer

//LoginAuthentication is use to check user's password.
type LoginAuthentication struct {
	Password string `json:"password"`
}
