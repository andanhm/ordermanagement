// Package config application config
package config

import (
	"encoding/json"
	"log"
	"strconv"
	"sync"
)

// config stores the configuration necessary to connect to a remote key/value store.
type config interface {
	Type() string
	Get(key string) (string, error)
}

// Handler is a provider that reads from one or more config sources.
type Handler struct {
	// client Store configration connection
	client config
	// enableLog error logs need to be enabled or disabled
	enableLog bool
	// mutex will synchronize access to `appDebug`.
	mutex *sync.Mutex
}

// New returns the instance of the config
func New() (*Handler, error) {
	c, err := newCache()
	if err != nil {
		log.Println(
			"fatal",
			"InternalServices",
			"CONFIG.NEW.ERROR",
			err.Error(),
			map[string]interface{}{},
			true,
		)
		return nil, err
	}
	return &Handler{
		client:    c,
		enableLog: true,
		mutex:     &sync.Mutex{},
	}, nil
}

// Enable receives a bool value, Modifies debug status.
// To safely access debug data across multiple goroutines.
func (c *Handler) Enable(debug bool) {
	// `Lock()` the `mutex` to ensure exclusive access to the `appDebug`,
	c.mutex.Lock()
	c.enableLog = debug
	//`Unlock()` the mutex, and change the value of the `appDebug`
	c.mutex.Unlock()
}

// Get function returns the string config value.
func (c *Handler) Get(key string) (string, error) {
	if c == nil {
		return "", ErrNotInitialized
	}
	val, err := c.client.Get(string(key))
	if err != nil && c.enableLog {
		log.Println(
			"error",
			"InternalServices",
			"CONFIG.GET.ERROR",
			err.Error(),
			map[string]interface{}{},
			true,
		)
		return "", err
	}
	return val, nil
}

// GetInt64 function returns the integer config value.
func (c *Handler) GetInt64(key string) (int64, error) {
	mode, err := c.Get(key)
	if err != nil {
		return 0, err
	}
	parsedValue, err := strconv.ParseInt(mode, 10, 64)
	if err != nil && c.enableLog {
		log.Println(
			"error",
			"InternalServices",
			"CONFIG.GET.ERROR",
			err.Error(),
			map[string]interface{}{},
			true,
		)
		return 0, err
	}
	return parsedValue, nil
}

// GetBool function returns the boolean config value.
func (c *Handler) GetBool(key string) (bool, error) {
	mode, err := c.Get(key)
	if err != nil {
		return false, err
	}
	parsedValue, err := strconv.ParseBool(mode)
	if err != nil && c.enableLog {
		log.Println(
			"error",
			"InternalServices",
			"CONFIG.BOOL.PARSER",
			err.Error(),
		)
		return false, err
	}
	return parsedValue, nil
}

// Unmarshal parses the JSON-encoded data and stores the result
// in the value pointed to by v. If v is nil or not a pointer,
// Unmarshal returns an Error.
func (c *Handler) Unmarshal(key string, v interface{}) error {
	val, err := c.Get(key)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(val), &v)
	if err != nil && c.enableLog {
		log.Println(
			"error",
			"InternalServices",
			"CONFIG.UNMARSHAL.ERROR",
			err.Error(),
			map[string]interface{}{},
			true,
		)
		return err
	}
	return nil
}
