package common

// IDGenerator represents the interface to generate IDs in use cases
type IDGenerator interface {
	// New generates a new ID
	New() string
}
