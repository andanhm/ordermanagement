package product

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"

	"github.com/andanhm/anglebroking/pkg/platform/db"
)

// Service is the interface to store implemented in package db. Store represent repository of the schema implementation supported by the datastore layer
type Service interface {
	// List fetches document in the product details.
	// If there is an error, you can check if it is a connection/validation/datanotfound error using a error code assertion
	List(context context.Context) ([]Details, error)

	// Save Now you'll have a new document in the schema collection.
	// If there is an error, you can check if it is a connection/validation/datanotfound error using a error code assertion
	Save(context context.Context, details []Details) error
}
type handle struct {
	db    db.Database
	table string
}

// New return the product instance
func New(db db.Database) Service {
	return &handle{
		table: "tblProduct",
		db:    db,
	}
}

func (h *handle) List(context context.Context) ([]Details, error) {
	data := make([]byte, 0)
	err := h.db.Find(h.table, "", &data)
	if err != nil {
		if err == redis.Nil {
			return nil, ErrNoProducts
		}
		return nil, err
	}
	details := make([]Details, 0)
	err = json.Unmarshal(data, &details)
	if err != nil {
		return nil, err
	}
	return details, nil
}

func (h *handle) Save(context context.Context, details []Details) error {
	data, _ := json.Marshal(details)
	err := h.db.Insert(h.table, "", 0, data)
	if err != nil {
		return err
	}
	return nil
}
