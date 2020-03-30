package app

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const databaseEnv = "DATABASE_URL"

type shrt struct {
	shrt string
	sha  string
	url  string
}

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
		BEGIN;
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

func queryBySha(db *sql.DB, sha string) (s shrt, err error) {
	query := `SELECT shrt, sha, url FROM shrts WHERE sha = $1;`
	err = db.QueryRow(query, sha).Scan(&s.shrt, &s.sha, &s.url)
	return
}

func insert(db *sql.DB, s shrt) (err error) {
	query := `INSERT INTO shrts (shrt, sha, url) VALUES ($1, $2, $3);`
	_, err = db.Exec(query, s.shrt, s.sha, s.url)
	return
}
