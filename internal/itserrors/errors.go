package itserrors

type Error struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	HTTPStatus int    `json:"http_status"`
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrNotFound         = Error{Code: "CLIENT_0001", Message: "Not found", HTTPStatus: 404}
	ErrInvalidSignature = Error{Code: "CLIENT_0002", Message: "Invalid signature", HTTPStatus: 400}
)
