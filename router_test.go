package service_test

import (
	"testing"

	"github.com/wchan2/go-service"
)

func TestingRouter(t *testing.T) {
	var (
		_ service.ServiceRouter = service.NewRouter()
		_ service.ServiceRouter = (*service.Router)(nil)
	)
}
