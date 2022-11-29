package customErrors

import (
	"github.com/pkg/errors"
)

func NewBadRequestError(message string, code int) error {
	br := &badRequestError{
		CustomError: NewCustomError(nil, code, message),
	}
	stackErr := errors.WithStack(br)

	return stackErr
}

func NewBadRequestErrorWrap(err error, code int, message string) error {
	br := &badRequestError{
		CustomError: NewCustomError(err, code, message),
	}
	stackErr := errors.WithStack(br)

	return stackErr
}

type badRequestError struct {
	CustomError
}

type BadRequestError interface {
	CustomError
	IsBadRequestError() bool
}

func (b *badRequestError) IsBadRequestError() bool {
	return true
}

func IsBadRequestError(err error) bool {
	var badRequestError BadRequestError
	//us, ok := grpc_errors.Cause(err).(BadRequestError)
	if errors.As(err, &badRequestError) {
		return badRequestError.IsBadRequestError()
	}

	return false
}