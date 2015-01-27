package router

import (
	"errors"
	"fmt"
	"strings"

	"net/http"
)

type AppRouter interface {
	ServeHTTP(w http.ResponseWriter, request *http.Request)
	Register(method string, path string, handler http.HandlerFunc)
}

type route struct {
	path    path
	handler http.HandlerFunc
}

type appRouter struct {
	routes map[string][]*route
}

func NewRouter() AppRouter {
	return &appRouter{routes: map[string][]*route{}}
}

func (router *appRouter) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	matchedRoute, err := router.match(request.Method, request.URL.Path)
	if err != nil {
		http.NotFound(w, request)
	} else {
		matchedRoute.ServeHTTP(w, request)
	}

}

func (router *appRouter) Register(method string, path string, handler http.HandlerFunc) {
	router.routes[strings.ToUpper(method)] = append(router.routes[strings.ToUpper(method)], &route{
		path:    newNamedPath(path),
		handler: handler,
	})
}

func (router *appRouter) match(method string, path string) (http.HandlerFunc, error) {
	routes := router.routes[strings.ToUpper(method)]
	if len(routes) == 0 {
		return nil, errors.New(fmt.Sprintf("Path %s, not found", path))
	}
	for _, route := range routes {
		matched, err := route.path.match(path)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Unable to match Path, %s with %s", path, route.path.uri()))
		}
		if matched {
			return route.handler, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Path %s not found", path))
}
