# Config package
This is a package for default config setup

## Usage
The `GetConfig` function takes 3 arguments:
- envFile - a string that represents the path to the `.env` file
- config - a pointer to a struct that represents the configuration

The `GetConfig` function returns an error if the configuration is not valid.

The `config` struct requires all the member variables to be strings. All the member variables are required by default. If you want to make a member variable optional, you can add the `required:"false"` tag to the member variable.

Here is an example of how to use this package.
```go
type Config struct {
	Port string // This is required by default
	OptionalSetting string `required:"false"`
}

var config Config

err := GetConfig(".env", &config)
if err != nil {
	t.Fatal(err)
}
```

An example of the logger impl you can find in the logging package app-log.


For local dev use `.env file. Example of file:
```bash
PORT=8080
```
