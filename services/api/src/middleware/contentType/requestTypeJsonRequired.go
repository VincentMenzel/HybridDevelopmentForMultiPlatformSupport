package contentType

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func RequestContentTypeJsonRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		contentType := request.Header.Get("content-type")

		if  contentType != "application/json" {
			_, err := writer.Write([]byte("only accepts json"))
			if err != nil {
				logrus.Error(err)
			}
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		next.ServeHTTP(writer, request)
	})
}
