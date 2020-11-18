package mysql

import (
	"database/sql"
	"fmt"
	"net/url"
	"sync"

	// mssql package has the drivers for MYSQL
	_ "github.com/go-sql-driver/mysql"
)

const (
	defaultPort = 3306
)

// Handler is a aerospike client store
type Handler struct {
	mutex  sync.Mutex
	config *Config
	Client *sql.DB
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

// String returns the connection string computed from the configuration
func (c *Config) String() string {
	query := url.Values{}

	dbName := c.DBName

	// Timeout for establishing connections, aka dial timeout
	// if its empty then 0s is the default time out
	if c.Timeout == "" {
		c.Timeout = "0s"
	}
	query.Add("timeout", fmt.Sprintf("%s", c.Timeout))
	query.Add("multiStatements", "true")
	if c.Port == 0 {
		c.Port = defaultPort
	}

	userInfo := url.UserPassword(c.Username, c.Password)
	host := fmt.Sprintf("%s:%d", c.Host, c.Port)
	connConfig := fmt.Sprintf("%s@tcp(%s)/%s?%s", userInfo.String(), host, dbName, query.Encode())
	return connConfig
}

// New returns a new instance of the DB handler
func New(cfg *Config) (*Handler, error) {
	db, err := sql.Open("mysql", cfg.String())
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(0)

	return &Handler{
		config: cfg,
		Client: db,
		mutex:  sync.Mutex{},
	}, nil
}

// Connect attempts to establish a connection with a AerospikeDB.
func (handler *Handler) Connect() error {

	return nil
}

// Ping returns an error if the aerospike database manager could not be reached.
func (handler *Handler) Ping() error {
	if handler == nil || handler.Client == nil {
		return ErrNotInitialized
	}
	err := handler.Client.Ping()
	if err != nil {
		return err
	}
	return nil
}

// Close closes all client connections to server nodes.
func (handler *Handler) Close() {
	if handler == nil || handler.config == nil || handler.Client == nil {
		return
	}
	handler.Client.Close()
}

// Insert a new document created with the set and user defined key are converted to a digest before sending to the server.
func (handler *Handler) Insert(set string, id string, ttl int64, schemas interface{}) error {
	if handler == nil || handler.Client == nil {
		return ErrNotInitialized
	}

	return nil
}

// Find fetches document in the collection bin
// If there is an error, you can check if it is a connection/validation/datanotfound error using a error code assertion
func (handler *Handler) Find(set string, id string, data interface{}) error {
	if handler == nil || handler.Client == nil {
		return ErrNotInitialized
	}

	return nil
}

// Remove finds the document matching the provided set and id, removes them from the database.
func (handler *Handler) Remove(set string, id string) error {
	if handler == nil || handler.Client == nil {
		return ErrNotInitialized
	}

	return nil
}
