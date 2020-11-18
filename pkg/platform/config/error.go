package config

import "errors"

var (
	// ErrNotInitialized tells the config library not initialized
	ErrNotInitialized = errors.New("configs: app config not initialized")
)
