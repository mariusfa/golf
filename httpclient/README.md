# HTTP Client

This package provides a resilient HTTP client with built-in fault tolerance patterns for GET, POST, PUT, and DELETE requests.

## Features

- **Circuit Breaker**: Prevents cascading failures by monitoring request success/failure rates
- **Bulkhead**: Isolates resources to prevent one failing service from affecting others  
- **Timeout Handling**: 15-second default timeout for all requests
- **JSON Support**: Built-in JSON serialization/deserialization
- **Request Logging**: Interface for transaction logging (bring your own logger)

## Resilience Patterns

The client implements enterprise-grade resilience patterns automatically:
- Circuit breaker monitors request patterns and fails fast when services are down
- Bulkhead pattern isolates different service calls to prevent resource exhaustion
- Configurable timeouts prevent hanging requests

## Dependencies
- github.com/jarcoal/httpmock (for testing)

## Usage

### Setup
```go
import "github.com/mariusfa/golf/httpclient"

client := httpclient.NewHttpClient()
```

### GET Requests
```go
var dto MyDto
requestId := "req-123" // For transaction logging
url := "http://api.example.com/users"
headers := map[string]string{"Accept": "application/json"}

getRequest := httpclient.NewGetRequest(requestId, headers, url)
err := client.GetJson(getRequest, &dto)
if err != nil {
    // Handle error (circuit breaker, timeout, etc.)
}
```

### POST Requests
```go
requestDto := MyDto{Name: "John Doe"}
requestId := "req-124"
url := "http://api.example.com/users"
headers := map[string]string{"Content-Type": "application/json"}

postRequest := httpclient.NewPostRequest(requestId, headers, url, requestDto)
err := client.PostJson(postRequest, nil)
if err != nil {
    // Handle error
}
```

### PUT Requests
```go
updateDto := MyDto{Name: "Jane Doe"}
requestId := "req-125"
url := "http://api.example.com/users/1"
headers := map[string]string{"Content-Type": "application/json"}

putRequest := httpclient.NewPutRequest(requestId, headers, url, updateDto)
err := client.PutJson(putRequest, nil)
if err != nil {
    // Handle error
}
```

### DELETE Requests
```go
requestId := "req-126"
url := "http://api.example.com/users/1"
headers := map[string]string{}

deleteRequest := httpclient.NewDeleteRequest(requestId, headers, url)
err := client.Delete(deleteRequest)
if err != nil {
    // Handle error
}
```

## Error Handling

The client returns errors for various failure scenarios:
- Network timeouts (after 15 seconds)
- Circuit breaker open (service unavailable)
- Bulkhead capacity exceeded
- HTTP error status codes
- JSON serialization/deserialization errors