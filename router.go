package servicelayer

import (
	"fmt"
	"strings"

	"golang.org/x/net/context"

	"net/http"
)

type AppRouter interface {
	ServeHTTP(rw http.ResponseWriter, request *http.Request)
	Register(method string, path string, handler HTTPHandler)
	Use(middlewares ...HTTPHandler)
}

type appRouter struct {
	middlewares []HTTPHandler
	routes      map[string][]*route
}

func NewRouter() AppRouter {
	return &appRouter{routes: map[string][]*route{}}
}

func (router *appRouter) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	matchedRoute, err := router.match(request.Method, request.URL.Path)
	if err != nil {
		http.NotFound(rw, request)
	} else {
		contxt := context.TODO()
		for _, middleware := range router.middlewares {
			middleware(contxt, rw, request)
		}
		matchedRoute(contxt, rw, request)
	}
}

func (router *appRouter) Register(method string, path string, handler HTTPHandler) {
	router.routes[strings.ToUpper(method)] = append(router.routes[strings.ToUpper(method)], &route{
		path:    newNamedPath(path),
		handler: handler,
	})
}

func (router *appRouter) Use(middlewares ...HTTPHandler) {
	router.middlewares = middlewares
}

func (router *appRouter) match(method string, path string) (HTTPHandler, error) {
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
