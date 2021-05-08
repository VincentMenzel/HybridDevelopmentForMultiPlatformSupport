package main

import (
	"github.com/VincentMenzel/HybridDevelopmentForMultiPlatformSupport/tree/development/services/api/src/datastore"
	"github.com/VincentMenzel/HybridDevelopmentForMultiPlatformSupport/tree/development/services/api/src/middleware/contentType"
	"github.com/VincentMenzel/HybridDevelopmentForMultiPlatformSupport/tree/development/services/api/src/routes/user"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func main() {
	store := datastore.NewSqlLiteDatabase()
	logrus.Info("SqlLiteDatastore created")

	router := mux.NewRouter()
	router.Use(contentType.RequestContentTypeJsonRequired)
	router.Use(contentType.SetResponseContentTypeJson)
	router.HandleFunc("/signup", user.CreateUserSignUpHandler(store)).Methods(http.MethodPost)
	router.HandleFunc("/signIn", user.CreateUserSignInHandler(store)).Methods(http.MethodPost)

	server := &http.Server{
		Addr:              "127.0.0.1:8080",
		Handler:           router,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       1 * time.Second,
	}
	
	logrus.Fatal(server.ListenAndServe())
}