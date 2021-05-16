package logging

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"net/http/httputil"
)

func LogRequestMiddleware(next http.Handler) http.Handler  {

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		if  body, err := httputil.DumpRequest(request, true); err != nil{
			logrus.Error(err)
		} else{
			logrus.Infof("%s: %s => %s", request.Method, request.RequestURI, string(body))
		}

		next.ServeHTTP(writer, request)
	})
}