package auth

import (
	"testing"
	"time"
)

func TestCreateAndParseToken(t *testing.T) {
	secret := "secret"
	token, err := CreateToken("user", secret, nil)
	if err != nil {
		t.Errorf("Error creating token: %s", err)
	}

	parsedUser, err := ParseToken(token, secret)
	if err != nil {
		t.Errorf("Error parsing token: %s", err)
	}
	if parsedUser != "user" {
		t.Errorf("Expected user to be 'user', got '%s'", parsedUser)
	}
}

func TestParseTokenInvalidSecret(t *testing.T) {
	secret := "secret"
	token, err := CreateToken("user", secret, nil)
	if err != nil {
		t.Errorf("Error creating token: %s", err)
	}

	_, err = ParseToken(token, "invalid")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestExpiredToken(t *testing.T) {
	secret := "secret"

	expires := time.Now().Add(-24 * time.Hour)
	
	token, err := CreateToken("user", secret, &expires)
	if err != nil {
		t.Errorf("Error creating token: %s", err)
	}

	_, err = ParseToken(token, secret)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
