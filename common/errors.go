package common

import (
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

// EntityNotFoundError records an error indicating an entity was not found in any case that caused it.
// You can use it to representing a Database not found, cache not found or any other repository.
type EntityNotFoundError struct {
	EntityType string
	Message    string
	Err        error
}

// NewEntityNotFoundError creates an instance of EntityNotFoundError
func NewEntityNotFoundError(entityType string) EntityNotFoundError {
	return EntityNotFoundError{
		EntityType: entityType,
		Message:    "",
		Err:        nil,
	}
}

// WrapEntityNotFoundError creates an instance of EntityNotFoundError
func WrapEntityNotFoundError(entityType string, err error) EntityNotFoundError {
	return EntityNotFoundError{
		EntityType: entityType,
		Message:    "",
		Err:        err,
	}
}

// Error implements the error interface
func (e EntityNotFoundError) Error() string {
	if strings.TrimSpace(e.Message) == "" {
		if strings.TrimSpace(e.EntityType) != "" {
			return fmt.Sprintf("Entity %s not found", e.EntityType)
		}
		if e.Err != nil && strings.TrimSpace(e.Message) == "" {
			return e.Err.Error()
		}
		return "entity not found"
	}
	return e.Message
}

// Unwrap implements the error interface introduced in Go 1.13 to unwrap the internal error
func (e EntityNotFoundError) Unwrap() error {
	return e.Err
}

// ValidationError records an error indicating an entity was not found in any case that caused it.
// You can use it to representing a Database not found, cache not found or any other repository.
type ValidationError struct {
	Message string
	Code    string
	Err     error
}

// Error implements the error interface
func (e ValidationError) Error() string {
	if strings.TrimSpace(e.Code) != "" {
		return fmt.Sprintf("%s - %s", e.Code, e.Message)
	}
	return e.Message
}

// Unwrap implements the error interface introduced in Go 1.13 to unwrap the internal error
func (e ValidationError) Unwrap() error {
	return e.Err
}

// EntityConflictError records an error indicating an entity already exists in some repository
// You can use it to representing a Database conflict, cache or any other repository.
type EntityConflictError struct {
	Message string
	Err     error
}

// Error implements the error interface
func (e EntityConflictError) Error() string {
	if e.Err != nil && strings.TrimSpace(e.Message) == "" {
		return e.Err.Error()
	}
	return e.Message
}

// Unwrap implements the error interface introduced in Go 1.13 to unwrap the internal error
func (e EntityConflictError) Unwrap() error {
	return e.Err
}

// ContainsError return true if the specific mongo error code match the WriteException and false if its not
func ContainsError(errs mongo.WriteException, code int) bool {
	for _, e := range errs.WriteErrors {
		if e.Code == code {
			return true
		}
	}
	return false
}

// UnauthorizedError indicates an operation that couldn't be performant because there's no user authenticated
type UnauthorizedError struct {
	Message string
	Code    string
	Err     error
}

func (e UnauthorizedError) Error() string {
	return e.Message
}

// ForbiddenError indicates an operation that couldn't be performant because the authenticated user has no sufficient privileges
type ForbiddenError struct {
	Message string
	Code    string
	Err     error
}

func (e ForbiddenError) Error() string {
	return e.Message
}

// UnprocessableOperationError indicates an operation that couldn't be performant because it's invalid
type UnprocessableOperationError struct {
	Message string
	Code    string
	Err     error
}

func (e UnprocessableOperationError) Error() string {
	return e.Message
}

//HTTPError indicates an http error raised in a http client
type HTTPError struct {
	Message string
	Code    int
	Err     error
}

func (e HTTPError) Error() string {
	return e.Message
}
