package rest_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalServerError(t *testing.T) {
	// because an internal server error can have multiple causes,
	// the NewInternalServerError should have two parameters.
	err := NewInternalServerError("this is the message", errors.New("database error"))

	// not expecting the error to be nil.
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)

	// not expecting a cause
	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "database error", err.Causes[0])

	// errBytes, _ := json.Marshal(err)
	// fmt.Println(string(errBytes))
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("a bad request error")

	// not expecting the error to be nil.
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "a bad request error", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("cannot find data")

	// not expecting the error to be nil.
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "cannot find data", err.Message)
	assert.EqualValues(t, "not_found", err.Error)
}

func TestNewError(t *testing.T) {
	err := NewError("some new error")

	// not expecting the error to be nil.
	assert.NotNil(t, err)
	assert.EqualValues(t, "some new error", err.Error())
}
