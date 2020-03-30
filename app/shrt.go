package app

import (
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"regexp"
)

var errURL = errors.New("Invalid URL")
var errDB = errors.New("Backend error")

var urlRegex = regexp.MustCompile(`^(https?|ftp)://[^\s/$.?#].[^\s]*$`)

func generateShrt(db *sql.DB, url string) (string, error) {
	if !urlRegex.Match([]byte(url)) {
		return "", errURL
	}
	sha := fmt.Sprintf("%x", sha256.Sum256([]byte(url)))
	row := shrt{
		shrt: sha[:8],
		sha:  sha,
		url:  url,
	}

	_, err := queryBySha(db, sha)
	if err == nil {
		return row.shrt, nil
	} else if err != sql.ErrNoRows {
		log.Println(err)
		return row.url, errDB
	}

	err = insert(db, row)
	if err != nil {
		log.Println(err)
		return row.url, errDB
	}
	return row.shrt, nil
}
