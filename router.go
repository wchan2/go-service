package service

import (
	"fmt"
	"strings"

	"golang.org/x/net/context"

	"net/http"
)

type ServiceRouter interface {
	ServeHTTP(rw http.ResponseWriter, request *http.Request)
	Register(method string, path string, handler HTTPHandler)
	Use(middlewares ...HTTPHandler)
}

type Router struct {
	middlewares []HTTPHandler
	routes      map[string][]*Route
}

func NewRouter() *Router {
	return &Router{routes: map[string][]*Route{}}
}

func (router *Router) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	matchedRoute, err := router.match(request.Method, request.URL.Path)
	if err != nil {
		http.NotFound(rw, request)
	} else {
		contxt := context.TODO()
		for _, middleware := range router.middlewares {
			middleware(contxt, request, rw)
		}
		matchedRoute(contxt, request, rw)
	}
}

func (router *Router) Register(method string, path string, handler HTTPHandler) {
	router.routes[strings.ToUpper(method)] = append(router.routes[strings.ToUpper(method)], &Route{
		path:    newNamedPath(path),
		handler: handler,
	})
}

func (router *Router) Use(middlewares ...HTTPHandler) {
	router.middlewares = middlewares
}

func (router *Router) match(method string, path string) (HTTPHandler, error) {
	routes := router.routes[strings.ToUpper(method)]
	if len(routes) == 0 {
		return nil, fmt.Errorf("Path %s, not found", path)
	}
	for _, route := range routes {
		matched, err := route.path.match(path)
		if err != nil {
			return nil, fmt.Errorf("Unable to match Path, %s with %s", path, route.path.uri())
		}
		if matched {
			return route.handler, nil
		}
	}
	return nil, fmt.Errorf("Path %s not found", path)
}
