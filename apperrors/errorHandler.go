package apperrors

import (
	"encoding/json"
	"errors"
	"github.com/ymho/qnin-sp/api/middlewares"
	"log"
	"net/http"
)

// ErrorHandler コントローラ層から呼び出されるエラーハンドラ
// TODO: コントローラ層での実装
func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	var appErr *MyAppError
	if !errors.As(err, &appErr) {
		appErr = &MyAppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	traceID := middlewares.GetTraceID(req.Context())
	log.Printf("[%d]error: %s\n", traceID, appErr)

	var statusCode int

	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed, BadParam:
		statusCode = http.StatusBadRequest
	case RequiredAuthorizationHeader, Unauthorized:
		statusCode = http.StatusUnauthorized
	default:
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}
