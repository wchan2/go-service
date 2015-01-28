package router

import "regexp"

type path interface {
	uri() string
	match(path string) (bool, error)
}

type namedPath struct {
	uriPath string
}

func newNamedPath(uriPath string) path {
	return &namedPath{uriPath: uriPath}
}

func (p *namedPath) match(path string) (bool, error) {
	namedParamRegEx, err := regexp.Compile(`(\(\?)?:\w+`)
	if err != nil {
		return false, err
	}
	regExURI := namedParamRegEx.ReplaceAllString(p.uriPath, `[a-zA-Z0-9]+`) + `/?$`
	return regexp.MatchString(regExURI, path)
}

func (p *namedPath) uri() string {
	return p.uriPath
}
