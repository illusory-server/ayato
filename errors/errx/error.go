package errx

import (
	"fmt"
	"github.com/illusory-server/ayato/errors/codex"

	"github.com/pkg/errors"
)

type Error struct {
	err  error
	code codex.Code
}

func New(code codex.Code, msg string) error {
	return &Error{
		code: code,
		err:  errors.New(msg),
	}
}

func Newf(code codex.Code, format string, args ...interface{}) error {
	return &Error{
		code: code,
		err:  errors.Errorf(format, args...),
	}
}

func (e *Error) Error() string {
	if e.err == nil {
		return ""
	}
	return e.err.Error()
}

func (e *Error) Unwrap() error {
	return e.err
}

func (e *Error) Code() codex.Code {
	return e.code
}

func (e *Error) Cause() error {
	return e.err
}

func (e *Error) Format(f fmt.State, verb rune) {
	switch verb {
	case 'v':
		if f.Flag('+') {
			if stackTracer, ok := e.err.(interface{ StackTrace() errors.StackTrace }); ok {
				_, _ = fmt.Fprintf(f, "%s (code: %d)\n", e.Error(), e.code)
				for _, frame := range stackTracer.StackTrace() {
					_, _ = fmt.Fprintf(f, "%+v\n", frame)
				}
			} else {
				_, _ = fmt.Fprintf(f, "%s (code: %d)", e.Error(), e.code)
			}
		} else {
			_, _ = fmt.Fprintf(f, "%s (code: %d)", e.Error(), e.code)
		}
	case 's':
		_, _ = fmt.Fprintf(f, "%s", e.Error())
	case 'q':
		_, _ = fmt.Fprintf(f, "%q", e.Error())
	default:
		_, _ = fmt.Fprintf(f, "%s (code: %d)", e.Error(), e.code)
	}
}

func (e *Error) StackTrace() errors.StackTrace {
	if tr, ok := e.err.(interface{ StackTrace() errors.StackTrace }); ok {
		return tr.StackTrace()
	}
	return nil
}

func WrapWithCode(err error, code codex.Code, msg string) error {
	return &Error{
		code: code,
		err:  errors.WithMessage(err, msg),
	}
}

func WrapWithCodef(err error, code codex.Code, format string, args ...interface{}) error {
	return &Error{
		code: code,
		err:  errors.WithMessagef(err, format, args...),
	}
}

func Code(err error) codex.Code {
	if err == nil {
		return codex.Unknown
	}
	var e *Error
	if errors.As(err, &e) {
		return e.code
	}
	return codex.Unknown
}
