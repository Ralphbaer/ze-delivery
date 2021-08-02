package common

import (
	"github.com/pkg/errors"
)

// ErrMongoDuplicatedDocument is throwed when a Document already exists in the repository.
var ErrMongoDuplicatedDocument = errors.New("Duplicated Document")

// SuppressNotFoundError return suppress errors when mongo: no documents in result are returned by mongo drive
func SuppressNotFoundError(err error) error {
	if err.Error() == "Entity not found" {
		return nil
	}
	return err
}
