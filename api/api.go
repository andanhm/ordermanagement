package api

import (
	"github.com/andanhm/anglebroking/service"
)

// API struct has all the dependencies required by the packages
type API struct {
	// protocol standard used to exchanging of data (http,gRPC)
	Protocol string
	Handler  *service.Handler
}

// New returns a new instance of the API struct
func New(service *service.Handler) *API {
	return &API{
		Protocol: "https",
		Handler:  service,
	}
}
