package errors

// Error 业务错误
type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

// New 创建业务错误
func New(code int, message string) *Error {
	return &Error{Code: code, Message: message}
}

// IsBizError 判断是否为业务错误
func IsBizError(err error) bool {
	_, ok := err.(*Error)
	return ok
}
