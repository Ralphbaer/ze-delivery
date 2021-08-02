package db

import (
	"errors"
)

// ErrRecordNotFound is throwed when a record was not found.
var ErrRecordNotFound = errors.New("record not found")
