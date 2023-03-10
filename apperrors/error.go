package apperrors

type MyAppError struct {
	ErrCode
	Message string //エラーメッセージ
	Err     error  `json:"-"` // エラーチェーン
}

func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}
