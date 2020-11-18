// Package db is a productive data access layer for Go that provides a common interface to work with different data sources
package db

import (
	"database/sql"

	"github.com/andanhm/anglebroking/pkg/platform/db/mysql"
	"github.com/andanhm/anglebroking/pkg/platform/db/redis"
)

// Database is an interface that defines methods that must be satisfied by
// all database adapters.
// Its primary job is to wrap existing implementations of such primitives,
// such as those in package aerospike etc..
type Database interface {

	// Open attempts to establish a connection with a DBMS.
	Connect() error

	// Ping returns an error if the database manager could not be reached.
	Ping() error

	// Close closes all client connections to database server nodes.
	Close()

	// Find prepares a query using the provided document.
	// The document may be a map or a struct value
	Find(collection string, id string, data interface{}) error

	// Insert inserts one or more documents in the respective collection.
	// ttl determines record expiration in seconds. Also known as TTL (Time-To-Live).
	// Seconds record will live before being removed by the server.
	Insert(collection string, id string, ttl int64, docs interface{}) error

	// Remove finds the documents matching the provided selector document and removes them from the database.
	Remove(collection string, id string) error
}

// config interface handler that allows to get the the configuration from key value store.
type config interface {
	Unmarshal(key string, v interface{}) error
}

// New returns new instance for the instance database
func New(config config) (Database, error) {

	c := new(redis.Config)
	err := config.Unmarshal("", c)
	if err != nil {
		return nil, err
	}
	handler, err := redis.New(c)
	if err != nil {
		return nil, err
	}
	if err := handler.Connect(); err != nil {
		return nil, err
	}
	if err := handler.Ping(); err != nil {
		return nil, err
	}
	return handler, nil
}

// NewSQL returns new instance for the instance database
func NewSQL(config config) (*sql.DB, error) {

	c := new(mysql.Config)
	err := config.Unmarshal("", c)
	if err != nil {
		return nil, err
	}
	handler, err := mysql.New(c)
	if err != nil {
		return nil, err
	}

	return handler.Client, nil
}
