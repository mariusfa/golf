package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func getMigrationConnectionString(dbConfig *DbConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
}

func getAppConnectionString(dbConfig *DbConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbConfig.AppUser, dbConfig.AppPassword, dbConfig.Host, dbConfig.Port, dbConfig.Name)
}

func Setup(dbConfig *DbConfig) (*sql.DB, error) {
	connectionString := getAppConnectionString(dbConfig)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Migrate(dbConfig *DbConfig, migrationsPath string) error {
	resolvedMigrationsPath := filepath.Join(migrationsPath, "resolved")
	if err := resolveAllTemplates(dbConfig, migrationsPath, resolvedMigrationsPath); err != nil {
		return err
	}

	m, err := migrate.New("file://"+resolvedMigrationsPath, getMigrationConnectionString(dbConfig))
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	deleteDirectory(resolvedMigrationsPath)
	return nil
}

func resolveAllTemplates(dbConfig *DbConfig, migrationsPath string, resolvedMigrationsPath string) error {
	if err := deleteDirectory(resolvedMigrationsPath); err != nil {
		return err
	}

	if err := createDirectory(resolvedMigrationsPath); err != nil {
		return err
	}

	if dbConfig.RunBaseLine == "true" {
		baselinePath := filepath.Join(migrationsPath, "baseline")
		if err := resolveTemplates(baselinePath, resolvedMigrationsPath, dbConfig); err != nil {
			return err
		}
	}

	standardPath := filepath.Join(migrationsPath, "standard")
	if err := resolveTemplates(standardPath, resolvedMigrationsPath, dbConfig); err != nil {
		return err
	}

	return nil
}

func resolveTemplates(fromPath string, toPath string, dbConfig *DbConfig) error {
	files, err := getFiles(fromPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if err := resolveFile(file, fromPath, toPath, dbConfig); err != nil {
			return err
		}
	}

	return nil
}

func getFiles(fromPath string) ([]string, error) {
	dir, err := os.ReadDir(fromPath)
	if err != nil {
		return nil, err
	}
	var files []string
	for _, file := range dir {
		println(file.Name())
		files = append(files, file.Name())
	}
	return files, nil
}

func resolveFile(file string, fromPath string, toPath string, dbConfig *DbConfig) error {
	// read file
	contentBytes, err := os.ReadFile(filepath.Join(fromPath, file))
	if err != nil {
		return err
	}

	// create new template
	tmpl, err := template.New(file).Parse(string(contentBytes))
	if err != nil {
		return err
	}

	// create new file to write resolved template
	resolvedFile, err := os.Create(filepath.Join(toPath, file))
	if err != nil {
		return err
	}
	defer resolvedFile.Close()

	// create a map of data to replace in the template
	data := map[string]string{
		"User":     dbConfig.AppUser,
		"Password": dbConfig.AppPassword,
	}

	// execute template with data and write to file
	err = tmpl.Execute(resolvedFile, data)
	if err != nil {
		return err
	}
	return nil
}

func deleteDirectory(path string) error {
	if _, err := os.Stat(path); err == nil {
		os.RemoveAll(path)
	}
	return nil
}

func createDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}
	return nil
}
