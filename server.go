package server

import (
	"net/http"

	"github.com/wchan2/service-layer/router"
)

type Server interface {
	Get(route string, handler http.HandlerFunc)
	Put(route string, handler http.HandlerFunc)
	Post(route string, handler http.HandlerFunc)
	Delete(route string, handler http.HandlerFunc)
	Patch(route string, handler http.HandlerFunc)
	Start(addr string)
}

func NewAppServer() Server {
	return &appServer{appRouter: router.NewRouter()}
}

type appServer struct {
	appRouter router.AppRouter
}

func (svr *appServer) Post(route string, handler http.HandlerFunc) {
	svr.appRouter.Register("POST", route, handler)
}

func (svr *appServer) Get(route string, handler http.HandlerFunc) {
	svr.appRouter.Register("GET", route, handler)
}

func (svr *appServer) Put(route string, handler http.HandlerFunc) {
	svr.appRouter.Register("PUT", route, handler)
}

func (svr *appServer) Delete(route string, handler http.HandlerFunc) {
	svr.appRouter.Register("DELETE", route, handler)
}

func (svr *appServer) Patch(route string, handler http.HandlerFunc) {
	svr.appRouter.Register("PATCH", route, handler)
}

func (svr *appServer) Start(addr string) {
	http.ListenAndServe(addr, svr.appRouter)
}
