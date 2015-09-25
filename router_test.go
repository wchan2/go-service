package service_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/wchan2/go-service"
	"golang.org/x/net/context"
)

func TestRouterInterface(t *testing.T) {
	var (
		_ service.ServiceRouter = service.NewRouter()
		_ service.ServiceRouter = (*service.Router)(nil)
	)
}

func TestRouterServeHTTP(t *testing.T) {
	middlewareCalled := false
	router := service.NewRouter()
	router.Use(func(contxt context.Context, rw http.ResponseWriter, req *http.Request) {
		middlewareCalled = true
	})
	router.Register("GET", "/healthcheck", func(contxt context.Context, rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})
	router.Register("GET", "/users/:param", func(contxt context.Context, rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})

	// test not found
	notFoundRecorder := httptest.NewRecorder()
	notFoundRequest, err := http.NewRequest("GET", "/notfound", nil)
	if err != nil {
		t.Fatal("unable to create HTTP request")
	}
	router.ServeHTTP(notFoundRecorder, notFoundRequest)
	if notFoundRecorder.Code != http.StatusNotFound {
		t.Error("expected route to not be found")
	}
	if middlewareCalled {
		t.Error("expected middleware not to be called on not found request, but called")
	}

	// test route without wildcards
	recorder := httptest.NewRecorder()
	httpRequest, err := http.NewRequest("GET", "/healthcheck", nil)
	if err != nil {
		t.Fatal("unable to create HTTP request")
	}
	router.ServeHTTP(recorder, httpRequest)
	if recorder.Code != http.StatusOK {
		t.Error("expected handler to be called, but it was not called")
	}
	if !middlewareCalled {
		t.Error("expected middleware to be called, but it was not called")
	}

	// test route with wildcards
	middlewareCalled = false
	namedParamRecorder := httptest.NewRecorder()
	namedParamHTTPRequest, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal("unable to create HTTP request")
	}
	router.ServeHTTP(namedParamRecorder, namedParamHTTPRequest)
	if namedParamRecorder.Code != http.StatusOK {
		t.Error("expected handler with a named parameter to be called, but it was not called")
	}
	if !middlewareCalled {
		t.Error("expected middleware to be called, but it was not called")
	}
}
