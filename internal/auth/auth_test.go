package auth

import (
	"log"
	"net/http"
	"testing"
)

// func GetAPIKey(headers http.Header) (string, error) {
// 	authHeader := headers.Get("Authorization")
// 	if authHeader == "" {
// 		return "", ErrNoAuthHeaderIncluded
// 	}
// 	splitAuth := strings.Split(authHeader, " ")
// 	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
// 		return "", errors.New("malformed authorization header")
// 	}
//
// 	return splitAuth[1], nil
// }

func TestGetAPIKey(t *testing.T) {
	// Successul header
	header := http.Header{
		"Authorization": {"ApiKey thisismykey"},
	}
	want := "thisismykey"
	got, err := GetAPIKey(header)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if got != want {
		log.Fatalf("Not equal! Got: %s and want: %s", got, want)
	}

	// No header included
	header = http.Header{}

	_, err = GetAPIKey(header)
	if err != nil {
		log.Fatal("ErrNoAuthHeaderIncluded error required")
	}

	// malformed header
	header.Add("Authorization", "Bearer thisismykey")
	_, err = GetAPIKey(header)
	if err == nil {
		log.Fatal("malformed authorization header required")
	}

	header.Set("Authorization", "Justonesimplelement")
	_, err = GetAPIKey(header)
	if err == nil {
		log.Fatal("malformed authorization header required")
	}
}
