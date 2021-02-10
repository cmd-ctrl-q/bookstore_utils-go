package rest_errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// RestErr is a rest error struct
type restErr struct {
	message string        `json:"message"`
	status  int           `json:"status"`
	error   string        `json:"error"`
	causes  []interface{} `json:"causes"` // causes of the error
}

// RestErr interface
type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

func (e restErr) Message() string {
	return e.message
}

func (e restErr) Status() int {
	return e.status
}

func (e restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: [%v ]",
		e.message, e.status, e.error, e.causes)
}

func (e restErr) Causes() []interface{} {
	return e.causes
}

// NewRestError is a custom error function
func NewRestError(message string, status int, err string, causes []interface{}) RestErr {
	return restErr{
		message: message,
		status:  status,
		error:   err,
		causes:  causes,
	}
}

// NewError returns a general message of the error.
// NewError is largely used to send a vague description back to an external caller.
func NewError(msg string) error {
	return errors.New(msg)
}

// NewRestErrorFromBytes attempts to create a RestErr
// If bytes object cannot be unmarshalled then return an
// invalid json error back to caller
func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var apiErr restErr
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

// NewBadRequestError returns a status bad request
func NewBadRequestError(message string) RestErr {
	return restErr{
		message: message,
		status:  http.StatusBadRequest,
		error:   "bad_request",
	}
}

// NewNotFoundError returns a status not found
func NewNotFoundError(message string) RestErr {
	return restErr{
		message: message,
		status:  http.StatusNotFound,
		error:   "not_found",
	}
}

// NewUnauthorizedError returns a rest error for unauthorized user
func NewUnauthorizedError(message string) RestErr {
	return restErr{
		message: message,
		status:  http.StatusUnauthorized,
		error:   "unauthorized",
	}
}

// NewInternalServerError returns an internal server error
func NewInternalServerError(message string, err error) RestErr {
	result := restErr{
		message: message,
		status:  http.StatusInternalServerError,
		error:   "internal_server_error",
	}
	if err != nil {
		result.causes = append(result.causes, err.Error())
	}
	return result
}
