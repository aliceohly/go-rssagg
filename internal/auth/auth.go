// why it is called auth??
package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(header http.Header) (string, error) {
	apiKey := header.Get("Authorization")

	if apiKey == "" {
		return "", errors.New("missing Authorization header")
	}

	splitAuth := strings.Split(apiKey, " ")
	if len(splitAuth) != 2 {
		return "", errors.New("invalid Authorization header")
	}
	if splitAuth[0] != "ApiKey" {
		return "", errors.New("invalid Authorization header")
	}

	return splitAuth[1], nil
}
