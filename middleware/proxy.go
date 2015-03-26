package middleware

import (
	"io/ioutil"
	"net/http"

	"github.com/wchan2/go-service-layer"
	"golang.org/x/net/context"
)

func Proxy(host, port string) servicelayer.HTTPHandler {
	return func(context context.Context, rw http.ResponseWriter, req *http.Request) {
		var (
			err        error
			proxiedReq *http.Request
			response   *http.Response
		)
		proxiedReq, err = http.NewRequest(req.Method, host+":"+port+req.URL.Path, req.Body)
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
