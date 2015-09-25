package middleware

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/wchan2/go-service"

	"golang.org/x/net/context"
)

func Proxy(host string, port string) service.HTTPHandler {
	return func(context context.Context, rw http.ResponseWriter, req *http.Request) {
		var (
			err        error
			proxiedReq *http.Request
			response   *http.Response
		)
		proxiedReq, err = http.NewRequest(req.Method, fmt.Sprintf("%s:%s%s", host, port, req.URL.Path), req.Body)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		proxiedReq.Form = req.Form
		proxiedReq.PostForm = req.PostForm
		proxiedReq.Header = req.Header

		response, err = http.DefaultClient.Do(req)
		if err != nil {
			rw.WriteHeader(http.StatusBadGateway)
			return
		}

		responseBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.WriteHeader(response.StatusCode)
		rw.Write(responseBody)
	}
}
