package response

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Response struct {
	Status string `json:"status"`          // Буде вказано Error або OK
	Error  string `json:"error,omitempty"` // omitempty параметр який можна вказати в страктегові json, що вказує, коли параметр в json буде пустий це поле буде відсутнє
}

const (
	StatusOk    = "OK"
	StatusError = "Error"
)

func OK() Response {
	return Response{
		Status: StatusOk,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errMsg []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "requires":
			errMsg = append(errMsg, fmt.Sprintf("faild %s is a requires faild", err.Field()))

		case "url":
			errMsg = append(errMsg, fmt.Sprintf("field %s in not a valid URL", err.Field()))

		default:
			errMsg = append(errMsg, fmt.Sprintf("field %s is not valid", err.Field()))

		}
	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsg, ""),
	}
}
