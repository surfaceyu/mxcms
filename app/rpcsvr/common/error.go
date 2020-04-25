package common

type Error struct {
	ErrMsg  string
}

func NewError(code int, msg string) *Error {
	return &Error{ErrMsg: msg}
}

func (err *Error) Error() string {
	return err.ErrMsg
}
