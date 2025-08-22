# Database

This package provides PostgreSQL database management with template-based migrations and Docker integration.

## Migration Strategy

The package supports two migration scenarios:

**Scenario 1: New application with fresh database**
- Use only `standard/` migrations
- Both local and production use the same migration files
- `RunBaseLine: "false"` (or omit)

**Scenario 2: Migrating existing application to Golf**
- Use `baseline/` to represent existing database state
- Use `standard/` for new changes going forward
- Example: Migrating from Spring Boot/Flyway to Go/Golf
- `RunBaseLine: "true"` for fresh local setup

## Template Support

Migrations support credential injection:
```sql
CREATE USER {{.User}} WITH PASSWORD '{{.Password}}';
```

## Configuration

```go
type DbConfig struct {
    Host         string // Database host
    Port         string // Database port  
    Name         string // Database name
    User         string // Migration user (admin privileges)
    Password     string // Migration user password
    AppUser      string // Application user (limited privileges)
    AppPassword  string // Application user password
    RunBaseLine  string // "true" when migrating from existing systems
    StartupLocal string // "true" to start Docker container locally
}
```

## Migration Structure

```
migrations/
├── baseline/           # Existing database state (only for migrations from other systems)
│   └── 1_existing_schema.up.sql
└── standard/          # New migrations (used in all scenarios)
    └── 2_new_feature.up.sql
```

## Usage

### New Application (Scenario 1)
```go
dbConfig := &database.DbConfig{
    Host:         "localhost",
    Port:         "5432", 
    Name:         "myapp",
    User:         "postgres",
    Password:     "postgres",
    AppUser:      "app_user", 
    AppPassword:  "app_password",
    // RunBaseLine omitted (defaults to false)
    StartupLocal: "true",
}

// Start container and run migrations
if dbConfig.StartupLocal == "true" {
    cleanup, err := database.SetupContainer(dbConfig)
    if err != nil {
        panic(err)
    }
    defer cleanup()
}

err := database.Migrate(dbConfig, "migrations/")
if err != nil {
    panic(err)
}

db, err := database.Setup(dbConfig)
if err != nil {
    panic(err)
}
defer db.Close()
```

### Migrating from Existing System (Scenario 2)
```go
dbConfig := &database.DbConfig{
    // ... same config as above
    RunBaseLine: "true", // Include baseline for fresh local setup
}

// Same setup code as above
```

### Testing
```go
func TestMain(m *testing.M) {
    dbConfig := database.DbConfig{
        User:        "test",
        Password:    "test",
        Name:        "test", 
        AppUser:     "app_user",
        AppPassword: "app_password",
    }
    database.SinglePostgresTestMain(m, &dbConfig, "migrations/")
}

func TestRepository(t *testing.T) {
    db, err := database.Setup(&dbConfig)
    if err != nil {
        t.Fatalf("Failed to setup database: %v", err)
    }
    defer db.Close()
    
    // Use db for testing
}
```

## Dependencies
- Docker (for local development)
- golang-migrate
- testcontainers-go
- lib/pq PostgreSQL driver
