package user

import (
	"encoding/json"
	"github.com/VincentMenzel/HybridDevelopmentForMultiPlatformSupport/tree/development/services/api/src/datastore"
	"github.com/VincentMenzel/HybridDevelopmentForMultiPlatformSupport/tree/development/services/api/src/response"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func CreateUserSignUpHandler(d *datastore.SqlLiteDatastore) func(w http.ResponseWriter, r *http.Request) {
	insertStmt, err := d.DB.Prepare(`
		INSERT INTO user 
		    (username, password) 
		VALUES (?,?)
	`)

	if err != nil {
		logrus.Panic(err)
	}

	searchStmt, err := d.DB.Prepare(`
		SELECT count(*) > 0
		FROM user
		WHERE username = ?
	`)

	if err != nil {
		logrus.Panic(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {

		var credentials struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		errorHandler := response.NewParameterErrorResponse(w)

		errorHandler.ValidateStringFieldIsNotMissingOrEmpty(credentials.Username, "username")
		errorHandler.ValidateStringFieldIsNotMissingOrEmpty(credentials.Password, "password")

		if errorHandler.WriteIfHasError() {
			errorHandler.PrintWarningToConsole()
			return
		}

		exists := false
		if err := searchStmt.QueryRow(credentials.Username).Scan(&exists); err != nil {
			logrus.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if exists {

			errorHandler.AddError(response.NewParameterAlreadyInUseError("username"))
			errorHandler.WriteIfHasError()
			return

		}

		if passwordHash, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost); err != nil {

			logrus.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return

		} else if result, err := insertStmt.Exec(credentials.Username, string(passwordHash)); err != nil {

			logrus.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return

		} else if id, err := result.LastInsertId(); err != nil {

			logrus.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return

		} else {

			w.WriteHeader(http.StatusCreated)
			logrus.Infof("User created: '%d'", id)

			if err := json.NewEncoder(w).Encode(response.NewCreatedResponse(id)); err != nil {
				logrus.Error(err)
				w.WriteHeader(http.StatusInternalServerError)

			}

		}
	}
}
