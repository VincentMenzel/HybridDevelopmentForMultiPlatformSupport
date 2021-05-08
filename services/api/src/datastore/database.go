package datastore

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type SqlLiteDatastore struct {
	DB *sql.DB
}

func NewSqlLiteDatabase() *SqlLiteDatastore {
	db, err := sql.Open("sqlite3", "HybridDevelopmentForMultiPlatformSupportData.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	newStore := &SqlLiteDatastore{
		DB: db,
	}

	if err = newStore.createTablesIfNotExists(); err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return newStore
}
