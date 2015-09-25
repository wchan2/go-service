package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wchan2/go-service"
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
