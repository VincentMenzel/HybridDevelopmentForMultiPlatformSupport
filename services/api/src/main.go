package main

import (
	"github.com/VincentMenzel/HybridDevelopmentForMultiPlatformSupport/tree/development/services/api/src/datastore"
	"github.com/VincentMenzel/HybridDevelopmentForMultiPlatformSupport/tree/development/services/api/src/handler/user"
	"github.com/VincentMenzel/HybridDevelopmentForMultiPlatformSupport/tree/development/services/api/src/middleware/contentType"
	"github.com/VincentMenzel/HybridDevelopmentForMultiPlatformSupport/tree/development/services/api/src/middleware/logging"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func main() {
	store := datastore.NewSqlLiteDatastore()
	logrus.Info("SqlLiteDatastore created")

	router := mux.NewRouter()
	router.Use(
		logging.LogRequestMiddleware,
		// added to avoid complications with preflight checks triggered by "Content-Type": "application/json"
		//contentType.RequestContentTypeJsonRequired,
		contentType.SetResponseContentTypeJson,
	)
	router.HandleFunc("/signup", user.CreateUserSignUpHandler(store)).Methods(http.MethodPost)
	router.HandleFunc("/signIn", user.CreateUserSignInHandler(store)).Methods(http.MethodPost, http.MethodOptions)

	headersOk := handlers.AllowedHeaders([]string{"Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	server := &http.Server{
		Addr:              "192.168.5.129:8080",
		Handler:           handlers.CORS(headersOk, originsOk, methodsOk)(router),
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       1 * time.Second,
	}
	logrus.Fatal(server.ListenAndServe())
}
