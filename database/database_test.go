package database

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveFile(t *testing.T) {
	path := "resolve_file_test"
	fromPath := filepath.Join(path, "from")
	toPath := filepath.Join(path, "resolved")
	dbConfig := DbConfig{
		AppUser:     "app_user",
		AppPassword: "app_password",
	}

	deleteDirectory(toPath)
	createDirectory(toPath)
	resolveFile("1_setup.up.sql", fromPath, toPath, &dbConfig)

	content, err := os.ReadFile(filepath.Join(toPath, "1_setup.up.sql"))
	if err != nil {
		t.Errorf("failed to read file 1_setup.up.sql in resolved directory: %v", err)
	}

	expectedContent := "USER app_user PASSWORD app_password;"
	if string(content) != expectedContent {
		t.Errorf("file 1_setup.up.sql contains incorrect content. Expected: %s, Got: %s", expectedContent, string(content))
	}

	deleteDirectory(toPath)
}

func TestResolveAllTemplates(t *testing.T) {
	path := "resolve_templates_test"
	toPath := filepath.Join(path, "resolved")
	dbConfig := DbConfig{
		AppUser:     "app_user",
		AppPassword: "app_password",
		RunBaseLine: "true",
	}
	if err := resolveAllTemplates(&dbConfig, path, toPath); err != nil {
		t.Errorf("failed to resolve all templates: %v", err)
	}

	content1, err := os.ReadFile(filepath.Join(toPath, "1_setup.up.sql"))
	if err != nil {
		t.Errorf("failed to read file 1_setup.up.sql in resolved directory: %v", err)
	}
	expectedContent1 := "USER app_user PASSWORD app_password;"
	if string(content1) != expectedContent1 {
		t.Errorf("file 1_setup.up.sql contains incorrect content. Expected: %s, Got: %s", expectedContent1, string(content1))
	}

	content2, err := os.ReadFile(filepath.Join(toPath, "2_setup.up.sql"))
	if err != nil {
		t.Errorf("failed to read file 2_setup.up.sql in resolved directory: %v", err)
	}
	expectedContent2 := "HELLO;"
	if string(content2) != expectedContent2 {
		t.Errorf("file 2_setup.up.sql contains incorrect content. Expected: %s, Got: %s", expectedContent2, string(content2))
	}

	deleteDirectory(toPath)
}

func TestResolveWithoutBaseline(t *testing.T) {
	path := "resolve_without_baseline_test"
	toPath := filepath.Join(path, "resolved")
	dbConfig := DbConfig{
		AppUser:     "app_user",
		AppPassword: "app_password",
		RunBaseLine: "false",
	}
	if err := resolveAllTemplates(&dbConfig, path, toPath); err != nil {
		t.Errorf("failed to resolve all templates: %v", err)
	}

	_, err := os.ReadFile(filepath.Join(toPath, "1_setup.up.sql"))
	if err == nil {
		t.Errorf("file 1_setup.up.sql should not exist in resolved directory")
	}

	content, err := os.ReadFile(filepath.Join(toPath, "2_setup.up.sql"))
	expectedContent := "HELLO;"
	if string(content) != expectedContent {
		t.Errorf("file 1_setup.up.sql contains incorrect content. Expected: %s, Got: %s", expectedContent, string(content))
	}

	deleteDirectory(toPath)
}
