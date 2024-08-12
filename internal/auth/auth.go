package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Header example: Authorization: APIKey {API Key}
func GetAPIKey(header http.Header) (string, error) {
	val := header.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication infomation found!")
	}
	apiKey := strings.Split(val, " ")
	if (len(apiKey) != 2) {
		return "", errors.New("wrong format of auth header")
	} 
	if (apiKey[0] != "APIKey") {
		return "", errors.New("wrong format of the first part of auth header")
	}
	return apiKey[1], nil
}