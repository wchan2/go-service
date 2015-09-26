package service_test

import (
	"net/http"
	"reflect"
	"testing"

	"golang.org/x/net/context"

	"github.com/wchan2/go-service-framework"
)

type mockServer struct {
	calledWith []interface{}
	called     bool
}

func (m *mockServer) Call(args ...interface{}) {
	m.calledWith = args
	m.called = true
}

func (m *mockServer) Called() bool {
	return m.called
}

func (m *mockServer) CalledWith() []interface{} {
	return m.calledWith
}

type fakeRouter struct {
	m *mockServer
}

func (f *fakeRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.m.Call(w, r)
}

func (f *fakeRouter) Register(method string, path string, handler service.HTTPHandler) {
	f.m.Call(method, path, handler)
}

func (f *fakeRouter) Use(handlers ...service.HTTPHandler) {
	f.m.Call(handlers)
}

func TestServiceServerInterface(t *testing.T) {
	var _ service.ServiceServer = (*service.Server)(nil)
}

func TestServerGet(t *testing.T) {
	var (
		router          *fakeRouter           = &fakeRouter{m: &mockServer{}}
		server          service.ServiceServer = service.NewServer(router)
		expectedhandler service.HTTPHandler   = func(_ context.Context, _ http.ResponseWriter, _ *http.Request) {}
	)

	server.Get("test_url", expectedhandler)
	if router.m.CalledWith()[0].(string) != "GET" {
		t.Error("server did not register proper handler method with router")
	}

	if router.m.CalledWith()[1].(string) != "test_url" {
		t.Error("server did not register proper url with router")
	}

	if reflect.DeepEqual(router.m.CalledWith()[2].(service.HTTPHandler), expectedhandler) {
		t.Error("server did not register proper handler with router")
	}
}

func TestServerPost(t *testing.T) {
	var (
		router          *fakeRouter           = &fakeRouter{m: &mockServer{}}
		server          service.ServiceServer = service.NewServer(router)
		expectedhandler service.HTTPHandler   = func(_ context.Context, _ http.ResponseWriter, _ *http.Request) {}
	)

	server.Post("test_url", expectedhandler)
	if router.m.CalledWith()[0].(string) != "POST" {
		t.Error("server did not register proper handler method with router")
	}

	if router.m.CalledWith()[1].(string) != "test_url" {
		t.Error("server did not register proper url with router")
	}

	if reflect.DeepEqual(router.m.CalledWith()[2].(service.HTTPHandler), expectedhandler) {
		t.Error("server did not register proper handler with router")
	}
}

func TestServerPatch(t *testing.T) {
	var (
		router          *fakeRouter           = &fakeRouter{m: &mockServer{}}
		server          service.ServiceServer = service.NewServer(router)
		expectedhandler service.HTTPHandler   = func(_ context.Context, _ http.ResponseWriter, _ *http.Request) {}
	)

	server.Patch("test_url", expectedhandler)
	if router.m.CalledWith()[0].(string) != "PATCH" {
		t.Error("server did not register proper handler method with router")
	}

	if router.m.CalledWith()[1].(string) != "test_url" {
		t.Error("server did not register proper url with router")
	}

	if reflect.DeepEqual(router.m.CalledWith()[2].(service.HTTPHandler), expectedhandler) {
		t.Error("server did not register proper handler with router")
	}
}

func TestServerPut(t *testing.T) {
	var (
		router          *fakeRouter           = &fakeRouter{m: &mockServer{}}
		server          service.ServiceServer = service.NewServer(router)
		expectedhandler service.HTTPHandler   = func(_ context.Context, _ http.ResponseWriter, _ *http.Request) {}
	)

	server.Put("test_url", expectedhandler)
	if router.m.CalledWith()[0].(string) != "PUT" {
		t.Error("server did not register proper handler method with router")
	}

	if router.m.CalledWith()[1].(string) != "test_url" {
		t.Error("server did not register proper url with router")
	}

	if reflect.DeepEqual(router.m.CalledWith()[2].(service.HTTPHandler), expectedhandler) {
		t.Error("server did not register proper handler with router")
	}
}

func TestServerDelete(t *testing.T) {
	var (
		router          *fakeRouter           = &fakeRouter{m: &mockServer{}}
		server          service.ServiceServer = service.NewServer(router)
		expectedhandler service.HTTPHandler   = func(_ context.Context, _ http.ResponseWriter, _ *http.Request) {}
	)

	server.Delete("test_url", expectedhandler)
	if router.m.CalledWith()[0].(string) != "DELETE" {
		t.Error("server did not register proper handler method with router")
	}

	if router.m.CalledWith()[1].(string) != "test_url" {
		t.Error("server did not register proper url with router")
	}

	if reflect.DeepEqual(router.m.CalledWith()[2].(service.HTTPHandler), expectedhandler) {
		t.Error("server did not register proper handler with router")
	}
}

func TestServerUse(t *testing.T) {
	var (
		router          *fakeRouter           = &fakeRouter{m: &mockServer{}}
		server          service.ServiceServer = service.NewServer(router)
		expectedhandler service.HTTPHandler   = func(_ context.Context, _ http.ResponseWriter, _ *http.Request) {}
	)

	server.Use(expectedhandler)
	if reflect.DeepEqual(router.m.CalledWith()[0].([]service.HTTPHandler), []service.HTTPHandler{expectedhandler}) {
		t.Error("server did not register the proper middleware with router")
	}
}
