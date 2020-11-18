package http

import (
	"encoding/json"
	"net/http"

	"github.com/andanhm/anglebroking/pkg/product"
)

// ListProduct return the list of products
func (s *Service) ListProduct(w http.ResponseWriter, r *http.Request) {
	details, err := s.handler.Handler.Product.List(r.Context())
	if err != nil {
		s.response(w, map[string]string{
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}
	s.response(w, details, http.StatusOK)
}

// SaveProduct return the save of products
func (s *Service) SaveProduct(w http.ResponseWriter, r *http.Request) {
	details := make([]product.Details, 0)
	if err := json.NewDecoder(r.Body).Decode(&details); err != nil {
		s.response(w, map[string]string{
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}
	err := s.handler.Handler.Product.Save(r.Context(), details)
	if err != nil {
		s.response(w, map[string]string{
			"message": err.Error(),
		}, http.StatusBadRequest)
		return
	}
	s.response(w, details, http.StatusOK)
}
