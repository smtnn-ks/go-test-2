package db

import (
	"database/sql"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
)

var Client *reform.DB
var db *sql.DB

func Init() {
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL is not specified")
	}

	var err error
	db, err = sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Unable to connect to DB")
	}

	logger := log.New().WithField("src", "SQL")

	Client = reform.NewDB(db, postgresql.Dialect, reform.NewPrintfLogger(logger.Infof))
	logger.Info("Connected")
}

func Close() error {
	return db.Close()
}
