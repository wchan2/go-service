package service

import (
	"net/http"

	"golang.org/x/net/context"
)

type HTTPHandler func(contxt context.Context, rw http.ResponseWriter, req *http.Request)

type Route struct {
	path    path
	handler HTTPHandler
}
