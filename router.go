package service

import (
	"fmt"
	"strings"

	"golang.org/x/net/context"

	"net/http"
)

// ServiceRouter is the interface includes methods for serving HTTP requests and registers handlers and middleware
type ServiceRouter interface {
	ServeHTTP(rw http.ResponseWriter, request *http.Request)
	Register(method string, path string, handler HTTPHandler)
	Use(middlewares ...HTTPHandler)
}

// Router for serving HTTP handlers and middleware
type Router struct {
	middlewares []HTTPHandler
	routes      map[string][]*Route
}

// NewRouter creates a router
func NewRouter() *Router {
	return &Router{routes: map[string][]*Route{}}
}

// ServeHTTP serves an HTTP response to an HTTP request to a matching router
func (router *Router) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	matchedRoute, err := router.match(request.Method, request.URL.Path)
	if err != nil {
		http.NotFound(rw, request)
	} else {
		contxt := context.Background()
		for _, middleware := range router.middlewares {
			middleware(contxt, rw, request)
		}
		matchedRoute(contxt, rw, request)
	}
}

// Register adds a handler that can be reached when a request matches the specified HTTP method and URL path
func (router *Router) Register(method string, path string, handler HTTPHandler) {
	router.routes[strings.ToUpper(method)] = append(router.routes[strings.ToUpper(method)], &Route{
		path:    newNamedPath(path),
		handler: handler,
	})
}

// Use adds a middleware handler that will preprocess the request when reaching a matching HTTP handler
func (router *Router) Use(middlewares ...HTTPHandler) {
	for _, middleware := range middlewares {
		router.middlewares = append(router.middlewares, middleware)
	}
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
