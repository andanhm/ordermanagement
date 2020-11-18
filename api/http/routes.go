package http

import (
	"encoding/json"
	"net/http"

	"github.com/bnkamalesh/webgo/v4"
)

// Routes returns a list of handlers defined for webGo
func Routes(s *Service) []*webgo.Route {
	return []*webgo.Route{
		{
			Name:          "*",
			Method:        http.MethodOptions,
			Pattern:       "/:w*",
			Handlers:      []http.HandlerFunc{s.CorsOptions},
			TrailingSlash: true,
		},
		{
			Name:    "status.r",
			Method:  http.MethodGet,
			Pattern: "/",
			Handlers: []http.HandlerFunc{func(w http.ResponseWriter, r *http.Request) {
				data, _ := json.Marshal(map[string]string{
					"version": "0.1.0",
					"app":     "Angle-Broker",
				})
				webgo.Send(w, webgo.JSONContentType, string(data), http.StatusOK)
			}},
			TrailingSlash: true,
		},
		{
			Name:          "health.r",
			Method:        http.MethodGet,
			Pattern:       "/health",
			Handlers:      []http.HandlerFunc{s.Health()},
			TrailingSlash: true,
		},
		{
			Name:          "products.r",
			Method:        http.MethodGet,
			Pattern:       "/products",
			Handlers:      []http.HandlerFunc{s.CorsOptions, s.ListProduct},
			TrailingSlash: true,
		},
		{
			Name:          "products.c",
			Method:        http.MethodPost,
			Pattern:       "/products",
			Handlers:      []http.HandlerFunc{s.SaveProduct},
			TrailingSlash: true,
		},
	}
}
