package app

import (
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"regexp"
)

var urlRegex = regexp.MustCompile(`^(https?|ftp)://[^\s/$.?#].[^\s]*$`)

func generateShrt(db *sql.DB, url string) (string, error) {
	if !urlRegex.Match([]byte(url)) {
		return "", errors.New("Invalid URL")
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
		return row.url, err
	}

	err = insert(db, row)
	if err != nil {
		return row.url, err
	}
	return row.shrt, nil
}
