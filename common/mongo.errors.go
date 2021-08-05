package common

import (
	"github.com/pkg/errors"
)

// ErrMongoDuplicatedDocument is throwed when a Document already exists in the repository.
var ErrMongoDuplicatedDocument = errors.New("Duplicated Document")
