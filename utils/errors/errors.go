package errors

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

var (
	ErrBadRequest = CustomError{
		Message:  "Bad Request",
		HTTPCode: http.StatusBadRequest,
	}

	ErrUnauthorized = CustomError{
		Message:  "Unauthorized",
		HTTPCode: http.StatusUnauthorized,
	}

	ErrForbidden = CustomError{
		Message:  "Forbidden",
		HTTPCode: http.StatusForbidden,
	}

	ErrNotFound = CustomError{
		Message:  "Record not exist",
		HTTPCode: http.StatusNotFound,
	}

	ErrUnprocessableEntity = CustomError{
		Message:  "Unprocessable Entity",
		HTTPCode: http.StatusUnprocessableEntity,
	}

	ErrFailedAuthentication = CustomError{
		Message:  "Invalid Credentials",
		HTTPCode: http.StatusUnauthorized,
	}

	ErrInternalServerError = CustomError{
		Message:  "Internal Server Error",
		HTTPCode: http.StatusInternalServerError,
	}
)

// CustomError holds data for customized error
type CustomError struct {
	ErrWithStack error       `json: error`
	Message      interface{} `json:"message"`
	HTTPCode     int         `json:"code"`
}

// Error is a function to convert error to string.
// It exists to satisfy error interface
func (c CustomError) Error() string {
	return fmt.Sprint(c.Message)
}

// New, Errorf, Wrap, and Wrapf record a stack trace at the point they are invoked.
// StackTracer can be used to retrieve this information
type StackTracer interface {
	StackTrace() errors.StackTrace
}

// New returns an error with the supplied message.
// New also records the stack trace at the point it was called.
func New(message string) error {
	return errors.New(message)
}

// withStackCustom annotates custom err with a stack trace at the point WithStack was called.
func withStackCustom(stackTraced error, err CustomError, message string) CustomError {
	return CustomError{
		ErrWithStack: stackTraced,
		HTTPCode:     err.HTTPCode,
		Message:      message,
	}
}

func CustomNew(customErr CustomError, message string) CustomError {
	return withStackCustom(WithStack(customErr), customErr, message)
}

func CustomWrap(err error, customErr CustomError, message string) CustomError {
	return withStackCustom(Wrap(err, customErr.Error()), customErr, message)
}

func CustomWithStack(customErr CustomError) CustomError {
	return withStackCustom(WithStack(customErr), customErr, customErr.Error())
}

func CustomWithMessage(customErr CustomError, message string) CustomError {
	newMessage := customErr.Error()
	if message != "" {
		newMessage = message
	}
	return withStackCustom(WithStack(customErr), customErr, newMessage)
}

// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
// Errorf also records the stack trace at the point it was called.
func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args)
}

// WithStack annotates err with a stack trace at the point WithStack was called.
// If err is nil, WithStack returns nil.
func WithStack(err error) error {
	return errors.WithStack(err)
}

// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called, and the supplied message.
// If err is nil, Wrap returns nil.
func Wrap(err error, message string) error {
	return errors.Wrap(err, message)
}

// Wrapf returns an error annotating err with a stack trace
// at the point Wrapf is called, and the format specifier.
// If err is nil, Wrapf returns nil.
func Wrapf(err error, format string, args ...interface{}) error {
	return errors.Wrapf(err, format, args)
}

// WithMessage annotates err with a new message.
// If err is nil, WithMessage returns nil.
func WithMessage(err error, message string) error {
	return errors.WithMessagef(err, message)
}

// WithMessagef annotates err with the format specifier.
// If err is nil, WithMessagef returns nil.
func WithMessagef(err error, format string, args ...interface{}) error {
	return errors.WithMessagef(err, format, args)
}

// Cause returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the following
// interface:
//
//	type causer interface {
//	       Cause() error
//	}
//
// If the error does not implement Cause, the original error will
// be returned. If the error is nil, nil will be returned without further
// investigation.
func Cause(err error) error {
	return errors.Cause(err)
}
