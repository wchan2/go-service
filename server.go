package servicelayer

import "net/http"

type Server interface {
	Get(route string, handler HTTPHandler)
	Put(route string, handler HTTPHandler)
	Post(route string, handler HTTPHandler)
	Delete(route string, handler HTTPHandler)
	Patch(route string, handler HTTPHandler)
	Start(addr string)
}

func NewAppServer() Server {
	return &appServer{appRouter: NewRouter()}
}

type appServer struct {
	appRouter AppRouter
}

func (svr *appServer) Post(route string, handler HTTPHandler) {
	svr.appRouter.Register("POST", route, handler)
}

func (svr *appServer) Get(route string, handler HTTPHandler) {
	svr.appRouter.Register("GET", route, handler)
}

func (svr *appServer) Put(route string, handler HTTPHandler) {
	svr.appRouter.Register("PUT", route, handler)
}

func (svr *appServer) Delete(route string, handler HTTPHandler) {
	svr.appRouter.Register("DELETE", route, handler)
}

func (svr *appServer) Patch(route string, handler HTTPHandler) {
	svr.appRouter.Register("PATCH", route, handler)
}

func (svr *appServer) Start(addr string) {
	http.ListenAndServe(addr, svr.appRouter)
}
