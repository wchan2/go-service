package service

import (
	"net/http"

	"regexp"

	"golang.org/x/net/context"
)

var namedParamRegEx *regexp.Regexp

func init() {
	var err error
	namedParamRegEx, err = regexp.Compile(`(\(\?)?:\w+`)
	if err != nil {
		panic(err)
	}
}

// HTTPHandler a handler to serve responses to HTTP requests
type HTTPHandler func(contxt context.Context, rw http.ResponseWriter, req *http.Request)

// Route routes a request to a specific handler based on the path
type Route struct {
	method  string
	path    string
	handler HTTPHandler
}

func (r *Route) matchPath(path string) (bool, error) {
	regExURI := namedParamRegEx.ReplaceAllString(r.path, `[a-zA-Z0-9]+`) + `/?$`
	return regexp.MatchString(regExURI, path)
}
