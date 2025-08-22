# Transaction Log

This package provides transaction logging for tracking business operations and external API calls.

## Usage

Initialize and use the transaction logger:

```go
import "github.com/mariusfa/golf/logging/transactionlog"

// In main function
appName := "todo"
transactionlog.SetAppName(appName)

// Used for logging transactions
transactionlog.Info("User created successfully")
transactionlog.Error("Payment processing failed")
```

## Purpose

Used for logging:
- Business transactions and operations
- External API calls and responses  
- Audit trails for compliance
- Performance monitoring of critical operations