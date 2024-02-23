package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/smtnn-ks/go-test-2/db"
	"github.com/smtnn-ks/go-test-2/router"
)

func main() {
	godotenv.Load(".env")
	db.Init()
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	app := router.Init()
	log.Fatal(app.Listen(":8000"))
}
