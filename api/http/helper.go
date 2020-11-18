package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bnkamalesh/webgo/v4"
)

// NotFound handler returns a 404 response based on the request data type
func (s *Service) NotFound(w http.ResponseWriter, r *http.Request) {
	s.response(w, fmt.Errorf("%s request resource not found", r.URL.Path), http.StatusNotFound)
}

// Response is used to respond to any request (JSON response) based on the code, data etc.
func (s *Service) response(w http.ResponseWriter, data interface{}, rCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(rCode)
	// To avoid multiple response.WriteHeader calls
	if data == nil {
		return
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		webgo.SendError(w, err.Error(), http.StatusBadRequest)
		return
	}
}
