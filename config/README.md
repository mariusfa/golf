# Config

This package provides environment-based configuration with automatic struct field mapping.

## Usage

The `GetConfig` function takes 2 arguments:
- `filename` - path to the `.env` file (optional, can be empty)
- `config` - pointer to a struct that will be populated

```go
type Config struct {
    Port string                    // Required by default
    DatabaseHost string            // Maps to DATABASE_HOST
    OptionalSetting string `required:"false"` // Optional field
}

var config Config
err := GetConfig(".env", &config)
if err != nil {
    log.Fatal(err)
}
```

## Field Mapping

Struct field names are automatically converted to SNAKE_CASE environment variables:
- `ServerPort` → `SERVER_PORT`
- `DatabaseHost` → `DATABASE_HOST`

## Configuration

- All fields must be strings
- Fields are required by default
- Use `required:"false"` tag for optional fields
- Missing required environment variables return an error

## Environment File

Example `.env` file:
```bash
PORT=8080
DATABASE_HOST=localhost
```
