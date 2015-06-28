package service

import (
	"net/http"

	"golang.org/x/net/context"
)

type HTTPHandler func(contxt context.Context, req *http.Request, rw http.ResponseWriter)

type Route struct {
	path    path
	handler HTTPHandler
}
