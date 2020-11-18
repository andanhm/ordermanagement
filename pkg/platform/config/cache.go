package config

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// cacheConfig is a prioritized configuration registry.
type cacheConfig struct {
	config *redis.Client
}

// newCache returns the redis based config which is guaranteed to be a singleton.
func newCache() (*cacheConfig, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	client := &cacheConfig{
		config: rdb,
	}
	return client, nil
}

// Type determines the type config
func (*cacheConfig) Type() string {
	return "environment"
}

// Get function returns the string config value.
func (c *cacheConfig) Get(key string) (string, error) {
	value := ""
	err := c.config.Get(context.Background(), key).Scan(&value)
	if err != nil {
		return value, err
	}
	return value, nil
}
