## Go Service ![build](https://travis-ci.org/wchan2/go-service-framework.svg?branch=master)

A minimal service library implemented in Go

## Dependencies

Install the dependencies via the below commands.

	go get golang.org/x/net/context

## Example

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wchan2/go-service-framework"
	"golang.org/x/net/context"
)

func main() {
	router := service.NewRouter()
	server := service.NewServer(router)

	// setting up a regular route
	server.Get("/healthcheck", func(contxt context.Context, rw http.ResponseWriter, req *http.Request) {
		response, err := json.Marshal(map[string]string{"status": "ok"})
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write(response)
	})

	// setting up a route with named parameters
	server.Get("/users/1", func(contxt context.Context, rw http.ResponseWriter, req *http.Request) {
		response, err := json.Marshal(map[string]string{"status": "ok"})
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.WriteHeader(http.StatusOK)
		rw.Write(response)
	})

	log.Print("server running on localhost:8080")
	log.Fatal(server.Run("localhost", "8080"))
}

```

## License

service-layer is released under the [MIT License](http://www.opensource.org/licenses/MIT).
