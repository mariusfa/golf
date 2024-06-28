# Transaction log package

This is a package for transaction logging.

## Usage

Here is an example of how to use this package.
Start with creating a package in `internal/logging` and add a file `logging.go` with the following content:
```go
package logging

import translog "github.com/mariusfa/gofl/v2/logging/trans-log"

var TransLogger *translog.TransLogger

func SetupTransLogger(appName string) {
	TransLogger = translog.NewTransLogger(appName)
}
```

In the main.go file, you can use the logger like this:
```go
package main

import (
	"<my-module>/internal/logging"

	"github.com/mariusfa/gofl/v2/config"
)

func main() {
	logging.SetupTransLogger("<my-module>")
	transLogger := logging.TransLogger

    // ... rest of the code
}
