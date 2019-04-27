package domain

// Error represents a custom error of this app
type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// NewError returns a error
func NewError(code int, mes string) *Error {
	return &Error{
		Message: mes,
		Code:    code,
	}
}

// Error returns the message of the error
func (e *Error) Error() string {
	return e.Message
}
