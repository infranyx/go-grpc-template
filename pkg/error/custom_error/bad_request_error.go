package customError

import (
	"github.com/pkg/errors"
)

func NewBadRequestError(message string, code int, details map[string]string) error {
	br := &badRequestError{
		CustomError: NewCustomError(nil, code, message, details),
	}
	// stackErr := error.WithStack(br)

	return br
}

func NewBadRequestErrorWrap(err error, message string, code int, details map[string]string) error {
	br := &badRequestError{
		CustomError: NewCustomError(err, code, message, details),
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

	if errors.As(err, &badRequestError) {
		return badRequestError.IsBadRequestError()
	}

	return false
}