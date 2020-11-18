package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/bnkamalesh/webgo/v4"

	"github.com/andanhm/anglebroking/api"
)

// Service struct has all the dependencies required by the packages
type Service struct {
	handler *api.API
}

// New returns a new instance of the HTTP struct
func New(handler *api.API) *Service {
	return &Service{
		handler: handler,
	}
}

// Health return the status of the underlying services that seo-engine relies on.
func (s *Service) Health() http.HandlerFunc {
	health := s.handler.Health(context.Background())
	bytes, _ := json.Marshal(health)
	str := string(bytes)
	return func(w http.ResponseWriter, r *http.Request) {
		webgo.Send(w, webgo.JSONContentType, str, http.StatusOK)
	}
}

// CorsOptions is a CORS middleware only for OPTIONS request method
func (s *Service) CorsOptions(rw http.ResponseWriter, req *http.Request) {
	allowedOrigin := ""
	reqOrigin := []byte(req.Header.Get("Origin"))
	if len(reqOrigin) == 0 {
		reqOrigin = []byte(req.Host)
	}

	for _, reg := range allowedDomains {
		if reg.Match(reqOrigin) {
			allowedOrigin = string(reqOrigin)
			break
		}
	}

	rw.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
	rw.Header().Set("Access-Control-Allow-Methods", allowMethods)
	rw.Header().Set("Access-Control-Allow-Credentials", "true")

	// Adding allowed headers
	rw.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,Access-Control-Request-Headers,"+req.Header.Get("Access-Control-Request-Headers"))
	webgo.SendHeader(rw, 200)
}
