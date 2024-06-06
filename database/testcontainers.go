package database

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

/**
 * This function is used to run a test that requires a single postgres container
 * The function will setup the container, run the test, and clean up the container
 */
func SinglePostgresTestMain(m *testing.M, dbConfig *DbConfig, migrationPath string) {
	cleanUp, err := SetupContainer(dbConfig)
	if err != nil {
		log.Fatalf("Failed to setup container: %v", err)
	}
	defer cleanUp()

	if err := Migrate(dbConfig, migrationPath); err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}

	code := m.Run()
	os.Exit(code)
}

/**
 * This will setup a new postgres container and it's cleanup function
 * The function will also change the host and port of the dbConfig to the container's host and port
 */
func SetupContainer(dbConfig *DbConfig) (func(), error) {
	container, err := createPostgresContainer(dbConfig)
	if err != nil {
		return func() {}, err
	}

	cleanUp := func() {
		if err := container.Terminate(context.Background()); err != nil {
			panic(err)
		}
	}

	getHostPortConfig(container, dbConfig)

	return cleanUp, nil
}

func createPostgresContainer(dbConfig *DbConfig) (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Image:        "postgres:15-alpine",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     dbConfig.User,
			"POSTGRES_PASSWORD": dbConfig.Password,
			"POSTGRES_DB":       dbConfig.Name,
		},
		WaitingFor: wait.ForListeningPort("5432/tcp").WithStartupTimeout(time.Duration(5 * time.Second)),
	}
	container, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}
	return container, nil
}

func getHostPortConfig(container testcontainers.Container, dbConfig *DbConfig) error {
	host, err := container.Host(context.Background())
	if err != nil {
		return err
	}
	port, err := container.MappedPort(context.Background(), "5432")
	if err != nil {
		return err
	}

	dbConfig.Host = host
	dbConfig.Port = port.Port()
	return nil
}
