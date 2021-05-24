package user

import (
	"database/sql"
	"encoding/json"
	"github.com/VincentMenzel/HybridDevelopmentForMultiPlatformSupport/tree/development/services/api/src/datastore"
	"github.com/VincentMenzel/HybridDevelopmentForMultiPlatformSupport/tree/development/services/api/src/model/user"
	"github.com/sirupsen/logrus"
	"net/http"
)

func CreateGetUserHandler(d *datastore.SqlLiteDatastore) func(w http.ResponseWriter, r *http.Request)  {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := d.DB.Query(`
			SELECT *
			FROM user`)

		defer func(stmt *sql.Rows) {
			err := stmt.Close()
			if err != nil {
				logrus.Error(err)
			}
		}(rows)

		if err != nil {
			logrus.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var users []user.User

		for rows.Next() {
			var selectedUser = &user.User{}
			if err = rows.Scan(&selectedUser.ID, &selectedUser.Username, &selectedUser.Password); err != nil {
				logrus.Error(err)
			}
			users = append(users, *selectedUser)
		}

		if err := json.NewEncoder(w).Encode(users); err != nil {
			logrus.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
