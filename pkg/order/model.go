package order

// Details ...
type Details struct {
	ID         int64  `json:"id"`
	CustomerID int64  `json:"customerID"`
	Amount     int64  `json:"amount"`
	Status     string `json:"status"`
}
