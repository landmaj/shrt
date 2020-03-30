package app

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const databaseEnv = "DATABASE_URL"

func newDatabase() *sql.DB {
	dbUrl, exists := os.LookupEnv(databaseEnv)
	if !exists {
		log.Fatalln("missing environmental variable:", databaseEnv)
	}
	dbUrl += "?sslmode=disable"
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
