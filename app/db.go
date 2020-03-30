package app

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const databaseEnv = "DATABASE_URL"

func NewDatabase() *sql.DB {
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

func CreateDatabase(db *sql.DB) {
	query := `
		BEGIN ;
		CREATE TABLE IF NOT EXISTS shrts
		(
			id   SERIAL,
			shrt TEXT NOT NULL,
			sha  TEXT NOT NULL,
			url  TEXT NOT NULL,
			PRIMARY KEY (id)
		);
		CREATE UNIQUE INDEX IF NOT EXISTS ix_shrts_shrt ON shrts (shrt);
		CREATE UNIQUE INDEX IF NOT EXISTS ix_shrts_sha ON shrts (sha);
		COMMIT;
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalln(err)
	}
}
