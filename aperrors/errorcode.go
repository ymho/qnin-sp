package aperrors

type ErrCode string

const (
	Unknown             ErrCode = "U0000"
	InsertDataFailed    ErrCode = "S0001"
	GetDataFailed       ErrCode = "S0002"
	NAData              Errcode = "S0003"
	NoTargetData        ErrCode = "S0004" //TODO: 実装
	UpdateDataFailed    ErrCode = "S0005" //TODO: 実装
	ReqBodyDecodeFailed ErrCode = "R0001" //TODO: 実装
	BadParam            ErrCode = "R0002" //TODO: 実装

)
