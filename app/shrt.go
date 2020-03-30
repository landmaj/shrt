package app

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"regexp"
)

var urlRegex = regexp.MustCompile(`^(https?|ftp)://[^\s/$.?#].[^\s]*$`)

func generateId(s string) (string, error) {
	if !urlRegex.Match([]byte(s)) {
		return "", errors.New("invalid URL")
	}
	id := fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
	return id, nil
}
