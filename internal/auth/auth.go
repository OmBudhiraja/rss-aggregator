package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts ApiKey from headers
// Eg. Authorization: ApiKey {api_key}
func GetApiKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")

	if authHeader == "" {
		return "", errors.New("auth header not found")
	}

	splittedHeader := strings.Split(authHeader, " ")

	if len(splittedHeader) != 2 {
		return "", errors.New("malformed auth header")
	}

	if splittedHeader[0] != "ApiKey" {
		return "", errors.New("malformed auth header")
	}

	if splittedHeader[1] == "" {
		return "", errors.New("invalid ApiKey")
	}

	return splittedHeader[1], nil

}
