package errors

import "errors"

var (
	ErrEmptyMacAddress   = errors.New("mac address must not be empty")
	ErrTimestampInFuture = errors.New("timestamp cannot be in the future")
)
