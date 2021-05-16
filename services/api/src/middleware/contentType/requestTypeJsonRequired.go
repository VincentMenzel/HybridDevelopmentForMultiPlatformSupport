package contentType

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)
const (
	Accepts = "Accept"
	ContentType = "Content-Type"
)
const (
	ApplicationJson = "application/json"
)

func RequestContentTypeJsonRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		writer.Header().Set(Accepts, ApplicationJson)

		contentType := request.Header.Get(ContentType)

		if !strings.Contains(contentType,  ApplicationJson) {

			writer.WriteHeader(http.StatusBadRequest)

			if _, err := writer.Write([]byte("invalid_content_type")); err != nil {
				logrus.Error(err)
			}
			return
		}

		next.ServeHTTP(writer, request)
	})
}
