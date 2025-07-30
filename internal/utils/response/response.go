package response

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOK    = "OK"
	StatusError = "ERROR"
)

func ValidationError(errs validator.ValidationErrors) Response {

	var errMsgs []string

	for _, err := range errs {

		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is required field", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is Invalid!", err.Field()))
		}

	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ", "),
	}

}