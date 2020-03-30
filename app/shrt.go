package app

import (
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"regexp"
)

const minShrtLen = 3

var errURL = errors.New("Invalid URL")
var errDB = errors.New("Backend error")

var urlRegex = regexp.MustCompile(`^(https?|ftp)://[^\s/$.?#].[^\s]*$`)

func createShrt(db *sql.DB, url string) (string, error) {
	if !urlRegex.Match([]byte(url)) {
		return "", errURL
	}
	sha := fmt.Sprintf("%x", sha256.Sum256([]byte(url)))
	row := shrt{
		sha: sha,
		url: url,
	}

	existing, err := queryBySha(db, sha)
	if err == nil {
		return existing.shrt, nil
	} else if err != sql.ErrNoRows {
		log.Println(err)
		return row.url, errDB
	}

	row.generateShrt(db)
	err = insert(db, row)
	if err != nil {
		log.Println(err)
		return row.url, errDB
	}
	return row.shrt, nil
}

func (sh *shrt) generateShrt(db *sql.DB) {
	for i := minShrtLen; i < len(sh.sha); i++ {
		fmt.Println("Loop", i)
		s := sh.sha[:i]
		_, err := queryByShrt(db, s)
		if err == sql.ErrNoRows {
			sh.shrt = s
			break
		}
	}
}
