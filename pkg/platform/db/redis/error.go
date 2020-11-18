package redis

import "errors"

var (
	// ErrNotInitialized tells that connection not initialized
	ErrNotInitialized = errors.New(
		"Sorry, unable to connect to the data source",
	)
)
