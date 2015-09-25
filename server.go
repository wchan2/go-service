package service

import (
	"fmt"

	"net/http"
)

const addressString = `%s:%s`

// ServiceServer is the interface for convenience methods that registers handlers to a given route and runs the server
type ServiceServer interface {
	Get(route string, handler HTTPHandler)
	Put(route string, handler HTTPHandler)
	Post(route string, handler HTTPHandler)
	Delete(route string, handler HTTPHandler)
	Patch(route string, handler HTTPHandler)
	Use(middleware ...HTTPHandler)
	Run(host string, port string) error
}

// Server serves HTTP requests
type Server struct {
	router ServiceRouter
}

// NewServer creates a new server to serve HTTP requests
func NewServer(router ServiceRouter) *Server {
	return &Server{router: router}
}

// Post registers a handler to serve a request that matches the POST method and the supplied route
func (svr *Server) Post(route string, handler HTTPHandler) {
	svr.router.Register("POST", route, handler)
}

// Get registers a handler to serve a request that matches the GET method and the supplied route
func (svr *Server) Get(route string, handler HTTPHandler) {
	svr.router.Register("GET", route, handler)
}

// Put registers a handler to serve a request that matches the Put method and the supplied route
func (svr *Server) Put(route string, handler HTTPHandler) {
	svr.router.Register("PUT", route, handler)
}

// Delete registers a handler to serve a request that matches the DELETE method and the supplied route
func (svr *Server) Delete(route string, handler HTTPHandler) {
	svr.router.Register("DELETE", route, handler)
}

// Patch registers a handler to serve a request that matches the PATCH method and the supplied route
func (svr *Server) Patch(route string, handler HTTPHandler) {
	svr.router.Register("PATCH", route, handler)
}

// Use registers handlers as middleware
func (svr *Server) Use(middleware ...HTTPHandler) {
	svr.router.Use(middleware...)
}

// Run runs the server on host and a port
func (svr *Server) Run(host string, port string) error {
	return http.ListenAndServe(fmt.Sprintf(addressString, host, port), http.Handler(svr.router))
}
