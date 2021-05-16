package datastore

import "log"
import (
	"github.com/sirupsen/logrus"
)

func (d SqlLiteDatastore) createTablesIfNotExists() error {
	d.createUserTable()

	logrus.Info("Table Creation concluded")

	return nil
}

func (d *SqlLiteDatastore) createUserTable() {

	res, err := d.DB.Exec(`
		CREATE TABLE IF NOT EXISTS "user" (
			user_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			username TEXT(20) NOT NULL,
			password TEXT(255) NOT NULL,
			CONSTRAINT user_UN UNIQUE (username)
		);
	`)

	if err != nil {
		log.Fatal(err)
	}

	if _, err := res.RowsAffected(); err != nil {
		log.Fatal(err)
	}

	logrus.Info("Table 'user' created")

}
