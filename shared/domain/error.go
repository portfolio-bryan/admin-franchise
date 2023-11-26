package domain

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/google/uuid"
)

type Error struct {
	ErrorID    string
	Err        error
	Message    string
	StackTrace []byte
	OccurredOn time.Time
	Misc       map[string]interface{}
}

type ErrorOpts struct {
	Message string
	Misc    map[string]interface{}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s; ErrorID: %s", e.Message, e.ErrorID)
}

func WrapError(err error, opts ErrorOpts) Error {
	if opts.Message == "" {
		opts.Message = err.Error()
	}

	return Error{
		ErrorID:    uuid.NewString(),
		Err:        err,
		Message:    opts.Message,
		StackTrace: debug.Stack(),
		OccurredOn: time.Now(),
		Misc:       opts.Misc,
	}
}
