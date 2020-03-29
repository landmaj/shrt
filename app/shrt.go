package app

import (
	"errors"
	"regexp"
)

var urlRegex = regexp.MustCompile(`^(https?|ftp)://[^\s/$.?#].[^\s]*$`)

func generateId(s string) (string, error) {
	if !urlRegex.Match([]byte(s)) {
		return "", errors.New("invalid URL")
	}
	return "PLACEHOLDER", nil
}
