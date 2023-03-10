package apperrors

type ErrCode string

const (
	Unknown                     ErrCode = "U0000"
	InsertDataFailed            ErrCode = "S0001"
	GetDataFailed               ErrCode = "S0002"
	NAData                      ErrCode = "S0003"
	NoTargetData                ErrCode = "S0004" //TODO: 実装
	UpdateDataFailed            ErrCode = "S0005" //TODO: 実装
	ReqBodyDecodeFailed         ErrCode = "R0001" //TODO: 実装
	BadParam                    ErrCode = "R0002" //TODO: 実装
	RequiredAuthorizationHeader ErrCode = "A0001"
	CannotMakeValidator         ErrCode = "A0002"
	Unauthorized                ErrCode = "A0003"
)

func (code ErrCode) Wrap(err error, message string) error {
	return &MyAppError{ErrCode: code, Message: message, Err: err}
}
