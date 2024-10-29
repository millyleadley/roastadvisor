package errors

import (
	"context"

	"github.com/pkg/errors"
)

// In the future we might want to make these more complex i.e.
// - Attach metadata from the context
// - Skip creating a stack trace when one already exists
// - Remove the bit in the stack trace about this errors package

func Wrap(ctx context.Context, cause error, msg string, args ...any) error {
	return errors.Wrapf(cause, msg, args...)
}

func New(ctx context.Context, msg string, args ...any) error {
	return errors.Errorf(msg, args...)
}
