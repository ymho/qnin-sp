package aperrors

type MyAppError struct {
	ErrCode
	Message string //エラーメッセージ
	Err     err    `json:"-"` // エラーチェーン
}

func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}

func (myErr *MyAppError) Wrap(err error, message string) error {
	return &MyAppError{ErrCode: code, Message: message, Err: err}
}
