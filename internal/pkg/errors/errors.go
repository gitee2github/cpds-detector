package errors

import "fmt"

const (
	SUCCESS        = 0
	DATABASE_ERROR = 101
	SOCKET_ERROR   = 102

	RULES_GET_ERROR    = 1001
	RULES_CREATE_ERROR = 1002
	RULES_UPDATE_ERROR = 1003
	RULES_DELETE_ERROR = 1004
)

var AnalyzerResultCodeMap = map[uint16]string{
	SUCCESS:        "Success",
	DATABASE_ERROR: "Database Error",
	SOCKET_ERROR:   "Network Error",

	RULES_GET_ERROR:    "Failed to get rule list",
	RULES_CREATE_ERROR: "Failed to create rule",
	RULES_UPDATE_ERROR: "Failed to update rule",
	RULES_DELETE_ERROR: "Failed to delete rule",
}

type Error struct {
	Err        error
	ResultCode uint16
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", AnalyzerResultCodeMap[e.ResultCode], e.Err.Error())
}

func NewError(resultCode uint16, err error) error {
	return &Error{
		ResultCode: resultCode,
		Err:        err,
	}
}

func IsErrorWithCode(err error, desiredResultCode uint16) bool {
	if err == nil {
		return false
	}

	serverError, ok := err.(*Error)
	if !ok {
		return false
	}

	return serverError.ResultCode == desiredResultCode
}
