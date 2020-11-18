package redis

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
)

const (
	defaultPort = 3306
)

// Handler is a aerospike client store
type Handler struct {
	mutex  sync.Mutex
	config *Config
	client *redis.Client
}

// Config struct has all the configurations used by MYSQL driver
type Config struct {
	DBName   string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	// Timeout is the connection timeout in seconds
	Timeout string `json:"timeout,omitempty"`
}

// New returns a new instance of the DB handler
func New(cfg *Config) (*Handler, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &Handler{
		config: cfg,
		client: rdb,
		mutex:  sync.Mutex{},
	}, nil
}

// Connect attempts to establish a connection with a AerospikeDB.
func (handler *Handler) Connect() error {
	return nil
}

// Ping returns an error if the aerospike database manager could not be reached.
func (handler *Handler) Ping() error {
	if handler == nil || handler.client == nil {
		return ErrNotInitialized
	}
	err := handler.client.Ping(context.Background()).Err()
	if err != nil {
		return err
	}
	return nil
}

// Close closes all client connections to server nodes.
func (handler *Handler) Close() {
	if handler == nil || handler.config == nil || handler.client == nil {
		return
	}
	handler.client.Close()
}

// Insert a new document created with the set and user defined key are converted to a digest before sending to the server.
func (handler *Handler) Insert(set string, id string, ttl int64, schemas interface{}) error {
	if handler == nil || handler.client == nil {
		return ErrNotInitialized
	}
	err := handler.client.Set(context.Background(), set+id, schemas, 0).Err()
	return err
}

// Find fetches document in the collection bin
// If there is an error, you can check if it is a connection/validation/datanotfound error using a error code assertion
func (handler *Handler) Find(set string, id string, data interface{}) error {
	if handler == nil || handler.client == nil {
		return ErrNotInitialized
	}
	err := handler.client.Get(context.Background(), set+id).Scan(data)
	if err != nil {
		return err
	}
	return nil
}

// Remove finds the document matching the provided set and id, removes them from the database.
func (handler *Handler) Remove(set string, id string) error {
	if handler == nil || handler.client == nil {
		return ErrNotInitialized
	}
	err := handler.client.Del(context.Background(), set+id).Err()
	if err != nil {
		return err
	}
	return nil
}
