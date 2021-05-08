package main

import (
	"github.com/VincentMenzel/HybridDevelopmentForMultiPlatformSupport/tree/development/services/api/src/datastore"
	"github.com/VincentMenzel/HybridDevelopmentForMultiPlatformSupport/tree/development/services/api/src/handler/user"
	"github.com/VincentMenzel/HybridDevelopmentForMultiPlatformSupport/tree/development/services/api/src/middleware/contentType"
	"github.com/VincentMenzel/HybridDevelopmentForMultiPlatformSupport/tree/development/services/api/src/middleware/logging"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func main() {
	store := datastore.NewSqlLiteDatastore()
	logrus.Info("SqlLiteDatastore created")

	router := mux.NewRouter()
	router.Use(logging.LogRequestMiddleware)
	router.Use(contentType.RequestContentTypeJsonRequired)
	router.Use(contentType.SetResponseContentTypeJson)
	router.HandleFunc("/signup", user.CreateUserSignUpHandler(store))
	router.HandleFunc("/signIn", user.CreateUserSignInHandler(store))

	server := &http.Server{
		Addr:              "192.168.5.129:8080",
		Handler:           router,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       1 * time.Second,
	}

	logrus.Fatal(server.ListenAndServe())
}