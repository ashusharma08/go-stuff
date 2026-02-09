package error

type ErrCode int

const (
	ERR_KEY_NOT_FOUND ErrCode = 1 + iota
)

type Error struct {
	Code    ErrCode
	Message string
}

func (e *Error) Error() string {
	switch e.Code {
	case ERR_KEY_NOT_FOUND:
		return "Key not found"
	default:
		return ""
	}
}
func (e *Error) String() string {
	switch e.Code {
	case ERR_KEY_NOT_FOUND:
		return "Key not found"
	default:
		return ""
	}
}

func NewError(code ErrCode, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
