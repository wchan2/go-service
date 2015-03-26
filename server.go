package servicelayer

import "net/http"

type Server interface {
	Get(route string, handler HTTPHandler)
	Put(route string, handler HTTPHandler)
	Post(route string, handler HTTPHandler)
	Delete(route string, handler HTTPHandler)
	Patch(route string, handler HTTPHandler)
	Use(middleware ...HTTPHandler)
	Start(addr string)
}

func NewServiceServer() Server {
	return &serviceServer{router: NewRouter()}
}

type serviceServer struct {
	router Router
}

func (svr *serviceServer) Post(route string, handler HTTPHandler) {
	svr.router.Register("POST", route, handler)
}

func (svr *serviceServer) Get(route string, handler HTTPHandler) {
	svr.router.Register("GET", route, handler)
}

func (svr *serviceServer) Put(route string, handler HTTPHandler) {
	svr.router.Register("PUT", route, handler)
}

func (svr *serviceServer) Delete(route string, handler HTTPHandler) {
	svr.router.Register("DELETE", route, handler)
}

func (svr *serviceServer) Patch(route string, handler HTTPHandler) {
	svr.router.Register("PATCH", route, handler)
}

func (svr *serviceServer) Use(middleware ...HTTPHandler) {
	svr.router.Use(middleware...)
}

func (svr *serviceServer) Start(addr string) {
	http.ListenAndServe(addr, svr.router)
}
