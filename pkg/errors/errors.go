package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	ErrorInternalServer = &Error{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
	}

	ErrorBadRequest = &Error{
		Code:    http.StatusBadRequest,
		Message: "Bad Request",
	}
)

type Error struct {
	Code    int
	Message string
}

func (err *Error) Error() string {
	return err.String()
}

func (err *Error) String() string {
	if err == nil {
		return ""
	}

	return fmt.Sprintf("error: code=%s message=%s", http.StatusText(err.Code), err.Message)
}

func (err *Error) Json() []byte {
	if err == nil {
		return []byte("{}")
	}

	res, _ := json.Marshal(err)
	return res
}

func (err *Error) StatusCode() int {
	if err == nil {
		return http.StatusOK
	}
	return err.Code
}
