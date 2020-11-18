package grpc

import (
	"context"

	"github.com/andanhm/anglebroking/pkg/order"

	"github.com/andanhm/anglebroking/api"

	v1 "github.com/andanhm/anglebroking/grpc/anglebroking/stock/order/v1"
)

// OrderService ...
type OrderService struct {
	handle *api.API
}

// NewOrder returns new access instance
func NewOrder(api *api.API) *OrderService {
	return &OrderService{
		handle: api,
	}
}
func getStatus(status string) v1.Status {
	switch status {
	case "FAILED":
		return v1.Status_FAILED
	case "SUCCESS":
		return v1.Status_SUCCESS
	}
	return v1.Status_PENDING
}

// Create ...
func (o *OrderService) Create(ctx context.Context, details *v1.Details) (*v1.Details, error) {
	order, err := o.handle.Handler.Order.Create(ctx, order.Details{
		CustomerID: details.CustomerID,
		Amount:     details.Amount,
	})
	if err != nil {
		return nil, err
	}
	details.Id = order.ID
	details.Status = getStatus(order.Status)
	return details, nil
}

// Update ...
func (o *OrderService) Update(ctx context.Context, details *v1.Details) (*v1.Details, error) {
	order, err := o.handle.Handler.Order.Update(ctx, order.Details{
		ID:     details.Id,
		Status: details.Status.String(),
	})
	if err != nil {
		return nil, err
	}
	details.Status = getStatus(order.Status)
	return details, nil
}

// Status ...
func (o *OrderService) Status(ctx context.Context, details *v1.Details) (*v1.Details, error) {
	order, err := o.handle.Handler.Order.Status(ctx, details.Id)
	if err != nil {
		return nil, err
	}
	details.Amount = order.Amount
	details.CustomerID = order.CustomerID
	details.Status = getStatus(order.Status)
	return details, nil
}
