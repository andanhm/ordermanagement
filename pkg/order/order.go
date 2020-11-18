package order

import (
	"context"
	"database/sql"
)

// Service is the interface to store implemented in package db. Store represent repository of the schema implementation supported by the datastore layer
type Service interface {
	// Create fetches document in the product details.
	// If there is an error, you can check if it is a connection/validation/datanotfound error using a error code assertion
	Create(context context.Context, details Details) (Details, error)

	// Update Now you'll have a new document in the schema collection.
	// If there is an error, you can check if it is a connection/validation/datanotfound error using a error code assertion
	Update(context context.Context, details Details) (Details, error)

	// Status Now you'll have a new document in the schema collection.
	// If there is an error, you can check if it is a connection/validation/datanotfound error using a error code assertion
	Status(context context.Context, id int64) (Details, error)
}

type handle struct {
	db    *sql.DB
	table string
}

// New return the product instance
func New(db *sql.DB) Service {
	return &handle{
		table: "tblProduct",
		db:    db,
	}
}

func (h *handle) Create(context context.Context, details Details) (Details, error) {
	details.Status = "PENDING"
	statement, err := h.db.Prepare(`INSERT INTO tblCustomerOrder (customer_id, amount, status) VALUES (?,?,?)`)
	if err != nil {
		return details, err
	}
	result, err := statement.Exec(details.CustomerID, details.Amount, details.Status)
	if err != nil {
		return details, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return details, err
	}

	details.ID = id
	return details, nil
}

func (h *handle) Update(context context.Context, details Details) (Details, error) {
	statement, err := h.db.Prepare(`UPDATE tblCustomerOrder SET status=? WHERE id=?`)
	if err != nil {
		return details, err
	}
	_, err = statement.Exec(details.Status, details.ID)
	if err != nil {
		return details, err
	}
	return details, nil
}

func (h *handle) Status(context context.Context, id int64) (Details, error) {
	details := Details{}
	row, err := h.db.Query(`SELECT customer_id,amount,status FROM tblCustomerOrder WHERE id=?`, id)
	if err != nil {
		return details, err
	}
	defer row.Close()
	for row.Next() {
		row.Scan(&details.CustomerID, &details.Amount, &details.Status)
	}

	return details, nil
}
