package config

import (
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {
	type Config struct {
		Port string
	}

	var config Config

	err := GetConfig(".env", &config)
	if err != nil {
		t.Fatal(err)
	}

	if config.Port != "8080" {
		t.Errorf("expected port to be 8080, got %v", config.Port)
	}
	os.Clearenv()
}

func TestGetConfigNoEnvFile(t *testing.T) {
	os.Setenv("PORT", "8080")
	type Config struct {
		Port string
	}

	var config Config

	err := GetConfig(".notExists", &config)
	if err != nil {
		t.Fatal(err)
	}

	if config.Port != "8080" {
		t.Errorf("expected port to be 8080, got %v", config.Port)
	}
	os.Clearenv()
}

func TestGetConfigMissingPort(t *testing.T) {
	type Config struct {
		Port string
	}

	var config Config

	err := GetConfig(".notExists", &config)
	if err == nil {
		t.Errorf("expected an error, got nil")
	}
	os.Clearenv()
}

func TestGetConfigWhenOptional(t *testing.T) {
	type Config struct {
		Port string `required:"false"`
	}

	var config Config

	err := GetConfig(".notExists", &config)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if config.Port != "" {
		t.Errorf("expected port to be empty, got %v", config.Port)
	}
	os.Clearenv()
}
