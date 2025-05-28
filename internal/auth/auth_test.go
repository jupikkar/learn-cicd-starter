package auth

import (
	"testing"
	"net/http"
)

func TestGetAPIKeyEmpty(t *testing.T) {
	// Arrange
	headers := http.Header{}

	// Act
	_, err := GetAPIKey(headers)

	// Assert
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if err != ErrNoAuthHeaderIncluded {
        t.Errorf("Expected error '%v', got '%v'", ErrNoAuthHeaderIncluded, err)
    }
}

func TestGetAPIKeyExistsInvalid(t *testing.T) {
	// Arrange
	headers := http.Header{}
	headers.Add("Authorization", "test_key_lol")

	// Act
	_, err := GetAPIKey(headers)

	// Assert
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if err != ErrMalformedAuthHeader {
        t.Errorf("Expected error '%v', got '%v'", ErrMalformedAuthHeader, err)
    }
}

func TestGetAPIKeyExistsValid(t *testing.T) {
	// Arrange
	headers := http.Header{}
	headers.Add("Authorization", "ApiKe test_key_lol")

	// Act
	key, err := GetAPIKey(headers)

	// Assert
	if err != nil {
		t.Errorf("Got error %s, expected nil", err)
	}
	if key == "" {
		t.Errorf("Expected API key, got empty string")
	}
}
