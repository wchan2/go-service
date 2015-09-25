package service

import (
	"net/http"

	"golang.org/x/net/context"
)

// HTTPHandler a handler to serve responses to HTTP requests
type HTTPHandler func(contxt context.Context, rw http.ResponseWriter, req *http.Request)

// Route routes a request to a specific handler based on the path
type Route struct {
	path    path
	handler HTTPHandler
}
