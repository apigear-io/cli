package core

import "errors"

// Common errors returned by wscore operations
var (
	// ErrConnectionNotFound is returned when a connection with the given ID doesn't exist
	ErrConnectionNotFound = errors.New("connection not found")

	// ErrConnectionClosed is returned when attempting operations on a closed connection
	ErrConnectionClosed = errors.New("connection closed")

	// ErrPoolClosed is returned when attempting operations on a closed pool
	ErrPoolClosed = errors.New("connection pool closed")

	// ErrDuplicateConnection is returned when attempting to add a connection with an existing ID
	ErrDuplicateConnection = errors.New("connection with this ID already exists")
)
