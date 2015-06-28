package service

import (
	"fmt"

	"net/http"
)

const addressString = `%s:%s`

type ServiceServer interface {
	Get(route string, handler HTTPHandler)
	Put(route string, handler HTTPHandler)
	Post(route string, handler HTTPHandler)
	Delete(route string, handler HTTPHandler)
	Patch(route string, handler HTTPHandler)
	Use(middleware ...HTTPHandler)
	Run(host string, port string) error
}

type Server struct {
	router ServiceRouter
}

func NewServer(router ServiceRouter) *Server {
	return &Server{router: router}
}

func (svr *Server) Post(route string, handler HTTPHandler) {
	svr.router.Register("POST", route, handler)
}

func (svr *Server) Get(route string, handler HTTPHandler) {
	svr.router.Register("GET", route, handler)
}

func (svr *Server) Put(route string, handler HTTPHandler) {
	svr.router.Register("PUT", route, handler)
}

func (svr *Server) Delete(route string, handler HTTPHandler) {
	svr.router.Register("DELETE", route, handler)
}

func (svr *Server) Patch(route string, handler HTTPHandler) {
	svr.router.Register("PATCH", route, handler)
}

func (svr *Server) Use(middleware ...HTTPHandler) {
	svr.router.Use(middleware...)
}

func (svr *Server) Run(host string, port string) error {
	return http.ListenAndServe(fmt.Sprintf(addressString, host, port), http.Handler(svr.router))
}
