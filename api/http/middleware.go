package http

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

var (
	allowedDomains []*regexp.Regexp
	allowMethods   = "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS"
)

func init() {
	dd := []string{
		"anglebroking.dev",
		"localhost",
	}

	allowedDomains = []*regexp.Regexp{}
	for _, d := range dd {
		r := fmt.Sprintf("%s(:[0-9]+)?$", d)
		reg, _ := regexp.Compile(r)
		allowedDomains = append(allowedDomains, reg)
	}

	methods := make([]string, 0, 12)
	// GET & HEAD are mandatory for any HTTP server
	methods = append(methods, http.MethodHead)
	methods = append(methods, http.MethodGet)
	methods = append(methods, http.MethodPost)

	allowMethods = strings.Join(methods, ",")
}
