package service

import (
	"database/sql"

	"github.com/andanhm/anglebroking/pkg/order"
	"github.com/andanhm/anglebroking/pkg/platform/db"
	"github.com/andanhm/anglebroking/pkg/product"
)

// config stores the configuration necessary to connect to a remote key/value store.
type config interface {
	Get(key string) (string, error)
	GetInt64(key string) (int64, error)
	GetBool(key string) (bool, error)
	Unmarshal(key string, v interface{}) error
}

// Handler struct holds all the section services
type Handler struct {
	Config  config
	Store   db.Database
	Product product.Service
	Order   order.Service
}

// New return the instance of the service
func New(config config, db db.Database) *Handler {
	product := product.New(db)
	return &Handler{
		Config:  config,
		Store:   db,
		Product: product,
	}
}

// NewOrder return the instance of the service
func NewOrder(config config, db *sql.DB) *Handler {
	order := order.New(db)
	return &Handler{
		Config: config,
		Order:  order,
	}
}
