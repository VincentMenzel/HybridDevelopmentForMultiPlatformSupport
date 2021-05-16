package user

import (
	"encoding/json"
	"github.com/VincentMenzel/HybridDevelopmentForMultiPlatformSupport/tree/development/services/api/src/datastore"
	"github.com/VincentMenzel/HybridDevelopmentForMultiPlatformSupport/tree/development/services/api/src/model/user"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func CreateUserSignInHandler(d *datastore.SqlLiteDatastore) func(w http.ResponseWriter, r *http.Request) {

	findUserStmt, err := d.DB.Prepare(`
		SELECT user_id, username, password
		FROM user
		WHERE username LIKE ?
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

			logrus.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return

		}

		var (
			id       int64
			username string
			password string
		)

		if err := findUserStmt.QueryRow(credentials.Username).Scan(&id, &username, &password); err != nil {

			logrus.Error(err)
			w.WriteHeader(http.StatusUnauthorized)
			return

		}

		if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(credentials.Password)); err != nil {

			logrus.Error(err)
			w.WriteHeader(http.StatusUnauthorized)
			return

		}

		err := json.NewEncoder(w).Encode(&user.User{
			ID:       id,
			Username: username,
		})

		if err != nil {

			logrus.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return

		}
	}
}
